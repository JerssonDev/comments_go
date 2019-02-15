package models

import jwt "github.com/dgrijalva/jwt-go"

// Claim ...Solicutud (Request), Token de usuario
type Claim struct {
	User User `json:"user"`
	jwt.StandardClaims
}
