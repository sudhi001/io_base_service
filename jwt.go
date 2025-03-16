package io_base_service

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
	"time"
)

// JWTManager handles token generation and validation
type JWTManager struct {
	secret string
}

// NewJWTManager initializes a new JWT manager
func NewJWTManager(secret string) *JWTManager {
	return &JWTManager{secret: secret}
}

// GenerateToken creates a JWT token
func (j *JWTManager) GenerateToken(deviceID string, expiry time.Duration) (string, error) {
	claims := map[string]interface{}{
		"device_id": deviceID,
		"exp":       time.Now().Add(expiry).Unix(),
	}

	// Convert claims to JSON
	claimsJSON, err := json.Marshal(claims)
	if err != nil {
		return "", err
	}

	// Encode header and claims
	header := base64.RawURLEncoding.EncodeToString([]byte(`{"alg":"HS256","typ":"JWT"}`))
	payload := base64.RawURLEncoding.EncodeToString(claimsJSON)

	// Create signature
	signature := j.sign(header, payload)
	token := header + "." + payload + "." + signature

	return token, nil
}

// ValidateToken checks if a JWT is valid and extracts the claims
func (j *JWTManager) ValidateToken(token string) (map[string]interface{}, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return nil, errors.New("invalid token format")
	}

	header, payload, signature := parts[0], parts[1], parts[2]
	expectedSignature := j.sign(header, payload)
	if expectedSignature != signature {
		return nil, errors.New("invalid token signature")
	}

	// Decode payload
	payloadBytes, err := base64.RawURLEncoding.DecodeString(payload)
	if err != nil {
		return nil, errors.New("invalid token payload")
	}

	// Parse claims
	var claims map[string]interface{}
	if err := json.Unmarshal(payloadBytes, &claims); err != nil {
		return nil, errors.New("invalid token claims")
	}

	// Check expiration
	exp, ok := claims["exp"].(float64)
	if !ok {
		return nil, errors.New("missing exp claim")
	}

	if int64(exp) < time.Now().Unix() {
		return nil, errors.New("token expired")
	}

	return claims, nil
}

// sign creates a simple HMAC-like signature (Not real HMAC, just a simple hashing method for demonstration)
func (j *JWTManager) sign(header, payload string) string {
	data := header + "." + payload + "." + j.secret
	hashed := base64.RawURLEncoding.EncodeToString([]byte(data))
	return hashed
}
