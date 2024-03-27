package utils

import (
	"encoding/base64"
	"encoding/json"
	"strings"
)

// TamperToken tampers with the token
func TamperToken(tokenString string) string {
	// Split the token into header, payload, and signature parts
	parts := strings.Split(tokenString, ".")
	if len(parts) != 3 {
		// Token should have three parts: header, payload, and signature
		return tokenString
	}

	// Decode the payload part
	decodedPayload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		// Failed to decode the payload, return the original token
		return tokenString
	}

	// Modify the payload (for example, change a field value)
	var payload map[string]interface{}
	if err := json.Unmarshal(decodedPayload, &payload); err != nil {
		// Failed to unmarshal the payload, return the original token
		return tokenString
	}

	// Modify the payload data (for example, change a field value)
	payload["username"] = "tampered_username"

	// Encode the modified payload back to base64
	modifiedPayload, err := json.Marshal(payload)
	if err != nil {
		// Failed to marshal the payload, return the original token
		return tokenString
	}

	// Encode the modified payload to base64 URL encoding
	modifiedPayloadEncoded := base64.RawURLEncoding.EncodeToString(modifiedPayload)

	// Replace the original payload with the modified payload
	parts[1] = modifiedPayloadEncoded

	// Join the parts back together to form the tampered token
	tamperedToken := strings.Join(parts, ".")

	return tamperedToken
}
