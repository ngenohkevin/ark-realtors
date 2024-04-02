package token

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/ngenohkevin/ark-realtors/internal/token"
	"github.com/ngenohkevin/ark-realtors/pkg/utils"
	"github.com/stretchr/testify/require"
	"regexp"
	"testing"
	"time"
)

func TestJWTMaker(t *testing.T) {
	maker, err := token.NewJWTMaker(utils.RandomString(32))
	require.NoError(t, err)

	username := utils.RandomUsername()
	duration := time.Minute

	// Generate token
	tokenString, err := maker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, tokenString)

	// Log the generated token
	fmt.Println("Generated Token:", tokenString)

	// Verify token
	payload, err := maker.VerifyToken(tokenString)
	if err != nil {
		fmt.Println("Error during token verification:", err)
	}

	require.NoError(t, err)
	require.NotNil(t, payload)

	// Log the payload
	fmt.Printf("Payload: %+v\n", payload)
}

func TestExpiredJWTToken(t *testing.T) {
	maker, err := token.NewJWTMaker(utils.RandomString(32))
	require.NoError(t, err)

	tokens, err := maker.CreateToken(utils.RandomUsername(), -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, tokens)

	payload, err := maker.VerifyToken(tokens)
	require.Error(t, err)

	fmt.Println("Actual error message:", err.Error())

	// regular expression to match the error message pattern
	expectedErrMsgPattern := regexp.MustCompile(`^.*token is expired.*$`)
	require.True(t, expectedErrMsgPattern.MatchString(err.Error()))

	require.Nil(t, payload)
}

// Compare this snippet from tests/token/jwt_maker_test.go:
func TestInvalidJWTToken(t *testing.T) {
	maker, err := token.NewJWTMaker(utils.RandomString(32))
	require.NoError(t, err)

	// Create a token with a valid expiration time
	tokenString, err := maker.CreateToken(utils.RandomUsername(), time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, tokenString)

	// Tamper with the token (for example, by modifying the payload)
	tamperedTokenString := utils.TamperToken(tokenString)
	fmt.Println("Tampered token:", tamperedTokenString)

	// Verify the tampered token
	payload, err := maker.VerifyToken(tamperedTokenString)
	fmt.Println("Error message:", err)

	// Assert that an error is returned and it matches the expected pattern
	require.Error(t, err)
	expectedErrMsgPattern := regexp.MustCompile(`^.*signature is invalid.*$`)
	require.True(t, expectedErrMsgPattern.MatchString(err.Error()))

	// Assert that payload is nil
	require.Nil(t, payload)
}

func TestInvalidJWTTokenAlgNone(t *testing.T) {
	// Create a payload
	payload, err := token.NewPayload(utils.RandomUsername(), time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	// Create custom claims as jwt.MapClaims
	customClaims := jwt.MapClaims{
		"id":       payload.ID,
		"username": payload.Username,
		"exp":      payload.ExpiredAt.Unix(),
	}

	// Create a JWT token with the "none" signing method
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone, customClaims)
	tokens, err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)
	require.NoError(t, err)

	// Attempt to verify the token (should fail because it's unsigned)
	maker, err := token.NewJWTMaker(utils.RandomString(32))
	require.NoError(t, err)

	_, err = maker.VerifyToken(tokens)
	require.Error(t, err)
	require.Contains(t, err.Error(), "unexpected signing method")
}
