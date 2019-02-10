package controllers

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-workspace/Comments_Go/commons"
	"github.com/go-workspace/Comments_Go/configuration"
	"github.com/go-workspace/Comments_Go/models"
)

// Login .. controlador de login
func Login(w http.ResponseWriter, r *http.Request) {

	user := &models.User{}

	err := json.NewDecoder(r.Body).Decode(user)

	if err != nil {
		fmt.Fprintf(w, "Error : %v\n", err.Error())
		return
	}

	db := configuration.GetConnection()

	defer db.Close()

	coding := sha256.Sum256([]byte(user.Password))

	pwd := fmt.Sprintf("%x", coding)

	db.Where("email = ? and password = ?", user.Email, pwd).Find(user)

	log.Println(user.ID, pwd)

	if user.ID > 0 {
		user.Password = ""
		token := commons.GenerateJWT(user)
		j, err := json.Marshal(models.Token{Token: token})

		if err != nil {
			log.Fatalf("Error al convertir el token a json : %v\n", err.Error())
		}

		w.WriteHeader(http.StatusOK)
		w.Write(j)
	} else {
		m := &models.Message{
			Message: "Usuario o clave no valido",
			Code:    http.StatusUnauthorized,
		}

		commons.DisplayMessage(w, m)
	}
}

// UserCreate ... permite registrar un usuario
func UserCreate(w http.ResponseWriter, r *http.Request) {

	user := &models.User{}

	m := &models.Message{}

	err := json.NewDecoder(r.Body).Decode(user)

	if err != nil {
		m.Message = fmt.Sprintf("Error al leer el usuario a registrar: %v", err.Error())
		m.Code = http.StatusBadRequest
		commons.DisplayMessage(w, m)
		return
	}

	if user.Password != user.ConfirmPassword {
		m.Message = "Las contraseñas no coinciden"
		m.Code = http.StatusBadRequest
		commons.DisplayMessage(w, m)
		return
	}

	coding := sha256.Sum256([]byte(user.Password))
	pwd := fmt.Sprintf("%x", coding)

	user.Password = pwd

	picmd5 := md5.Sum([]byte(user.Email))
	picstr := fmt.Sprintf("%x", picmd5)

	user.Pickture = "https://gravatar.com/avatar/" + picstr + "?s=100"

	db := configuration.GetConnection()
	defer db.Close()

	err = db.Create(user).Error

	if err != nil {
		m.Message = fmt.Sprintf("Error al crear el registro: %v", err.Error())
		m.Code = http.StatusBadRequest
		commons.DisplayMessage(w, m)
		return
	}

	m.Message = "Usuario creado con exito"
	m.Code = http.StatusCreated
	commons.DisplayMessage(w, m)

}
