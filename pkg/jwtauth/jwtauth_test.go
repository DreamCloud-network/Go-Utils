package jwtauth_test

import (
	"crypto/rand"
	"crypto/rsa"
	"errors"
	"testing"
	"time"

	// Import your new package
	"github.com/DreamCloud-network/Go-Utils/pkg/jwtauth"

	// Import the jwt library to check for specific errors
	"github.com/golang-jwt/jwt/v5"
)

// generateTestKeys is a helper function to create a new RSA key pair for each test.
// This ensures test isolation.
func generateTestKeys(t *testing.T) *rsa.PrivateKey {
	t.Helper() // Marks this as a test helper function

	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatalf("Failed to generate RSA key pair: %v", err)
	}
	return privateKey
}

// TestJWTFlow_Success is the "happy path" test.
// It checks if a token can be generated and then successfully validated.
func TestJWTFlow_Success(t *testing.T) {
	// --- 1. SETUP ---
	privateKey := generateTestKeys(t)
	publicKey := &privateKey.PublicKey

	const (
		issuer   = "test-iam-service"
		ttl      = 15 * time.Minute
		userID   = "user-123-abc"
		tenantID = "tenant-789-xyz"
	)

	// --- 2. CREATE ISSUER & VALIDATOR ---
	tokenIssuer := jwtauth.NewIssuer(privateKey, issuer, ttl)
	tokenValidator := jwtauth.NewValidator(publicKey, issuer)

	// --- 3. GENERATE TOKEN ---
	tokenString, err := tokenIssuer.GenerateToken(userID, tenantID)
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}
	if tokenString == "" {
		t.Fatal("Generated token string is empty")
	}

	// --- 4. VALIDATE TOKEN ---
	claims, err := tokenValidator.ValidateToken(tokenString)
	if err != nil {
		t.Fatalf("Failed to validate token: %v", err)
	}

	// --- 5. ASSERT CLAIMS ---
	if claims.UserID != userID {
		t.Errorf("UserID mismatch: expected '%s', got '%s'", userID, claims.UserID)
	}
	if claims.TenantID != tenantID {
		t.Errorf("TenantID mismatch: expected '%s', got '%s'", tenantID, claims.TenantID)
	}
	if claims.Issuer != issuer {
		t.Errorf("Issuer mismatch: expected '%s', got '%s'", issuer, claims.Issuer)
	}
}

// TestValidator_Failures tests all the "sad paths" where validation
// is expected to fail.
func TestValidator_Failures(t *testing.T) {
	// --- SETUP (Shared for all sub-tests) ---
	keyA := generateTestKeys(t)
	keyB := generateTestKeys(t) // A different key

	const (
		issuerA = "issuer-A"
		issuerB = "issuer-B"
		userID  = "user-123"
		ttl     = 15 * time.Minute
	)

	// --- SUB-TEST: Expired Token ---
	t.Run("ExpiredToken", func(t *testing.T) {
		// 1. Create an issuer with a *negative* TTL to generate an expired token
		expiredIssuer := jwtauth.NewIssuer(keyA, issuerA, -5*time.Minute)
		tokenString, err := expiredIssuer.GenerateToken(userID, "tenant-1")
		if err != nil {
			t.Fatalf("Failed to generate expired token: %v", err)
		}

		// 2. Create a validator
		validator := jwtauth.NewValidator(&keyA.PublicKey, issuerA)

		// 3. Try to validate
		_, err = validator.ValidateToken(tokenString)

		// 4. Assert
		if err == nil {
			t.Fatal("Expected an error for expired token, but got nil")
		}
		if !errors.Is(err, jwt.ErrTokenExpired) {
			t.Errorf("Expected error '%v', got '%v'", jwt.ErrTokenExpired, err)
		}
	})

	// --- SUB-TEST: Invalid Issuer ---
	t.Run("InvalidIssuer", func(t *testing.T) {
		// 1. Create a token from "issuer-A"
		issuer := jwtauth.NewIssuer(keyA, issuerA, ttl)
		tokenString, _ := issuer.GenerateToken(userID, "tenant-1")

		// 2. Create a validator that expects "issuer-B"
		validator := jwtauth.NewValidator(&keyA.PublicKey, issuerB)

		// 3. Try to validate
		_, err := validator.ValidateToken(tokenString)

		// 4. Assert
		if err == nil {
			t.Fatal("Expected an error for invalid issuer, but got nil")
		}
		if !errors.Is(err, jwt.ErrTokenInvalidIssuer) {
			t.Errorf("Expected error '%v', got '%v'", jwt.ErrTokenInvalidIssuer, err)
		}
	})

	// --- SUB-TEST: Invalid Signature (Wrong Key) ---
	t.Run("InvalidSignature", func(t *testing.T) {
		// 1. Create a token with Key A
		issuer := jwtauth.NewIssuer(keyB, issuerA, ttl)
		tokenString, _ := issuer.GenerateToken(userID, "tenant-1")

		// 2. Create a validator that uses Key B
		validator := jwtauth.NewValidator(&keyA.PublicKey, issuerA)

		// 3. Try to validate
		_, err := validator.ValidateToken(tokenString)

		// 4. Assert
		if err == nil {
			t.Fatal("Expected an error for invalid signature, but got nil")
		}
		if !errors.Is(err, jwt.ErrTokenSignatureInvalid) {
			t.Errorf("Expected error '%v', got '%v'", jwt.ErrTokenSignatureInvalid, err)
		}
	})

	// --- SUB-TEST: Malformed Token ---
	t.Run("MalformedToken", func(t *testing.T) {
		validator := jwtauth.NewValidator(&keyA.PublicKey, issuerA)

		// Try to validate a garbage string
		_, err := validator.ValidateToken("not.a.real.token")

		if err == nil {
			t.Fatal("Expected an error for malformed token, but got nil")
		}
		if !errors.Is(err, jwt.ErrTokenMalformed) {
			t.Errorf("Expected error '%v', got '%v'", jwt.ErrTokenMalformed, err)
		}
	})
}
