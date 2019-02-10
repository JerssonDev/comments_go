package commons

import (
	"crypto/rsa"
	"io/ioutil"
	"log"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/go-workspace/Comments_Go/models"
)

var (
	privateKey *rsa.PrivateKey

	// PublicKey ... se usa para validar el token
	PublicKey *rsa.PublicKey
)

func init() {

	privateBytes, err := ioutil.ReadFile("./keys/private.rsa")

	if err != nil {
		log.Printf("No se pudo leer el archivo privado %v", err.Error())
	}

	publicBytes, err := ioutil.ReadFile("./keys/public.rsa")

	if err != nil {
		log.Printf("No se pudo leer el archivo publico %v", err.Error())
	}

	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateBytes)

	if err != nil {
		log.Printf("No se pudo hacer el parse a llave privada %v", err.Error())
	}

	PublicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicBytes)

	if err != nil {
		log.Printf("No se pudo hacer el parse a llave publica %v", err.Error())
	}

}

// GenerateJWT ... genera el token para el cliente
func GenerateJWT(user *models.User) string {

	claims := models.Claim{
		User: *user,
		StandardClaims: jwt.StandardClaims{
			//ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer: "JerssonDev",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	result, err := token.SignedString(privateKey)

	if err != nil {
		log.Printf("No se pudo firmar el token %v", err.Error())
	}

	return result
}
