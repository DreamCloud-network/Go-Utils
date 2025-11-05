package jwtauth

import (
	"crypto/rsa"
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

// ErrInvalidToken is returned when token validation fails.
var ErrInvalidToken = errors.New("invalid or expired token")

// Validator is responsible for parsing and validating JWTs.
// This component will be used by all your microservices *except* the iam-service.
type Validator struct {
	publicKey *rsa.PublicKey
	issuer    string
}

// NewValidator creates a new token validator.
// The caller (e.g., notification-service) must provide the public key
// and the expected issuer string (e.g., "iam-service").
func NewValidator(publicKey *rsa.PublicKey, issuer string) *Validator {
	return &Validator{
		publicKey: publicKey,
		issuer:    issuer,
	}
}

// ValidateToken parses and validates a token string.
// On success, it returns the CustomClaims.
// On failure, it returns a specific, debuggable error (e.g., ErrTokenExpired).
func (v *Validator) ValidateToken(tokenString string) (*CustomClaims, error) {
	claims := &CustomClaims{}

	// keyFunc is the callback function required by the jwt library.
	// It supplies the public key for verification.
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		// Security check: ensure the token's signing method is what we expect.
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("jwtauth.validator.ValidateToken - unexpected signing method: %v", token.Header["alg"])
		}
		// Return the public key
		return v.publicKey, nil
	}

	// Create a new parser with validation options
	parser := jwt.NewParser(
		// This option ensures the 'exp' (ExpiresAt) claim is required and validated.
		jwt.WithExpirationRequired(),
		// This option ensures the 'iss' (Issuer) claim is required and matches
		// the issuer string we provided to the Validator.
		jwt.WithIssuer(v.issuer),
	)

	// Parse the token
	token, err := parser.ParseWithClaims(tokenString, claims, keyFunc)

	// --- THIS IS THE KEY IMPROVEMENT ---
	// The 'err' returned by the parser is NOT generic. It will be a
	// specific error like jwt.ErrTokenExpired, jwt.ErrTokenInvalidIssuer,
	// jwt.ErrTokenSignatureInvalid, etc.
	if err != nil {
		// We return the error directly. The caller can now use errors.Is()
		// to check exactly *why* the token failed.
		return nil, fmt.Errorf("jwtauth.validator.ValidateToken - token validation failed: %w", err)
	}

	// We can still do this check for defense-in-depth
	if !token.Valid {
		return nil, errors.New("jwtauth.validator.ValidateToken - token is not valid")
	}

	// Token is valid and all claims have been parsed into the 'claims' struct
	return claims, nil
}
