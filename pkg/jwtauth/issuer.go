package jwtauth

import (
	"crypto/rsa"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Issuer is responsible for creating and signing new JWTs.
// This component should ONLY be used by your iam-service.
type Issuer struct {
	privateKey *rsa.PrivateKey
	issuer     string
	tokenTTL   time.Duration
}

// NewIssuer creates a new instance of the Issuer.
// The caller (e.g., iam-service) must provide the private key and issuer details.
func NewIssuer(privateKey *rsa.PrivateKey, issuer string, tokenTTL time.Duration) *Issuer {
	return &Issuer{
		privateKey: privateKey,
		issuer:     issuer,
		tokenTTL:   tokenTTL,
	}
}

// GenerateToken creates a new signed JWT for a specific user and tenant.
func (i *Issuer) GenerateToken(userID, tenantID string) (string, error) {
	// Creates the claims (the payload)
	claims := CustomClaims{
		UserID:   userID,
		TenantID: tenantID,
		RegisteredClaims: jwt.RegisteredClaims{
			// Defines the expiration time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(i.tokenTTL)),
			// Defines the issuer
			Issuer: i.issuer,
			// Define when it was issued
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}

	// Creates a new token with RS256 signing method and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	// Signs the token with the private key
	respString, err := token.SignedString(i.privateKey)
	if err != nil {
		return "", fmt.Errorf("jwtauth.Issuer.GenerateToken - failed to sign token: %w", err)
	}
	return respString, nil
}
