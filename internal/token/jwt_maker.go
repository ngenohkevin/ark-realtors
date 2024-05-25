package token

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

const minSecretKeySize = 32

type JWTMaker struct {
	secretKey []byte
}

func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid key size: must be at least %d characters", minSecretKeySize)
	}

	return &JWTMaker{[]byte(secretKey)}, nil
}

func (maker *JWTMaker) CreateToken(username string, role string, duration time.Duration) (string, *Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return "", nil, err
	}

	payload, err := NewPayload(username, role, duration)
	if err != nil {
		return "", payload, err
	}

	claims := jwt.MapClaims{
		"id":        tokenID.String(),
		"username":  username,
		"role":      role,
		"exp":       payload.ExpiredAt.Unix(),
		"issued_at": payload.IssuedAt.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(maker.secretKey)
	if err != nil {
		return "", payload, err
	}

	return signedToken, payload, nil
}

func (maker *JWTMaker) VerifyToken(tokenString string) (*Payload, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return maker.secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	// Check if the token is valid
	if !token.Valid {
		return nil, ErrInvalidToken
	}

	// Extract claims from the token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, ErrInvalidToken
	}

	// Parse claims and create payload
	username, ok := claims["username"].(string)
	if !ok {
		return nil, ErrInvalidToken
	}

	role, ok := claims["role"].(string)
	if !ok {
		return nil, ErrInvalidToken
	}

	idString, ok := claims["id"].(string)
	if !ok {
		return nil, ErrInvalidToken
	}

	id, err := uuid.Parse(idString)
	if err != nil {
		return nil, ErrInvalidToken
	}

	issuedAtUnix, ok := claims["issued_at"].(float64)
	if !ok {
		return nil, ErrInvalidToken
	}

	expiredAtUnix, ok := claims["exp"].(float64)
	if !ok {
		return nil, ErrInvalidToken
	}

	issuedAt := time.Unix(int64(issuedAtUnix), 0)
	expiredAt := time.Unix(int64(expiredAtUnix), 0)

	return &Payload{
		ID:        id,
		Username:  username,
		Role:      role,
		IssuedAt:  issuedAt,
		ExpiredAt: expiredAt,
	}, nil
}
