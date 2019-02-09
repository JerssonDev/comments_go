package models

import jwt "github.com/dgrijalva/jwt-go"

// Claim ...Solicutud (Request)
type Claim struct {
	User User `son:"user"`
	jwt.StandardClaims
}
