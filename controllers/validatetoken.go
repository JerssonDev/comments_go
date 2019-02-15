package controllers

import (
	"context"
	"net/http"

	"github.com/go-workspace/Comments_Go/commons"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/go-workspace/Comments_Go/models"
)

// ValidateToken ... Permite validar el token del cliente
func ValidateToken(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	m := &models.Message{}

	token, err := request.ParseFromRequestWithClaims(
		r,
		request.OAuth2Extractor,
		&models.Claim{},
		func(t *jwt.Token) (interface{}, error) {
			return commons.PublicKey, nil
		},
	)

	if err != nil {
		m.Code = http.StatusUnauthorized

		switch err.(type) {

		case *jwt.ValidationError:

			vErr := err.(*jwt.ValidationError)

			switch vErr.Errors {

			case jwt.ValidationErrorExpired:
				m.Message = "Su token ha expirado"
				commons.DisplayMessage(w, m)
				return

			case jwt.ValidationErrorSignatureInvalid:
				m.Message = "La firma del token no coincide"
				commons.DisplayMessage(w, m)
				return

			default:
				m.Message = "su token no es valido"
				commons.DisplayMessage(w, m)
				return

			}

		}
	}

	if token.Valid {

		ctx := context.WithValue(r.Context(), "user", token.Claims.(*models.Claim).User)
		next(w, r.WithContext(ctx))
	} else {

		m.Code = http.StatusUnauthorized
		m.Message = "su token no es valido"
		commons.DisplayMessage(w, m)
	}

}
