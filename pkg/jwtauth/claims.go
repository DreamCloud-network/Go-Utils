package jwtauth

import "github.com/golang-jwt/jwt/v5"

// CustomClaims defines the custom data payload we embed in our JWT.
// Any service importing this library will know how to read the token's
// data structure.
type CustomClaims struct {
	UserID   string `json:"uid"`
	TenantID string `json:"tid"`
	jwt.RegisteredClaims
}
