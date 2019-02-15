package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-workspace/Comments_Go/configuration"

	"github.com/go-workspace/Comments_Go/commons"

	"github.com/go-workspace/Comments_Go/models"
)

// CommentCreate ...crea comentarios
func CommentCreate(w http.ResponseWriter, r *http.Request) {

	comment := &models.Comment{}
	user := models.User{}
	m := &models.Message{}

	user, _ = r.Context().Value("user").(models.User)

	err := json.NewDecoder(r.Body).Decode(comment)

	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = fmt.Sprintf("Estructura no valida, error al leer el comentario: %v", err.Error())
		commons.DisplayMessage(w, m)
		return
	}
	comment.UserID = user.ID

	db := configuration.GetConnection()
	defer db.Close()

	err = db.Create(comment).Error

	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = fmt.Sprintf("Error al registrar el comentario : %v", err.Error())
		commons.DisplayMessage(w, m)
		return
	}

	m.Code = http.StatusCreated
	m.Message = "Comentado Registrado correctamente"
	commons.DisplayMessage(w, m)

}

// CommentGetAll ... obtiene todos los comentarios
func CommentGetAll(w http.ResponseWriter, r *http.Request) {

	comments := []models.Comment{}
	m := &models.Message{}
	user := models.User{}
	votes := &models.Vote{}

	user, _ = r.Context().Value("user").(models.User)

	vars := r.URL.Query() // devuelve el paso de variables despues del prefix

	db := configuration.GetConnection()
	defer db.Close()

	queryComment := db.Where("parent_id = ? ", 0)

	if order, ok := vars["order"]; ok {
		if order[0] == "votes" {
			queryComment = queryComment.Order("votes desc, created_at desc")
		}
	} else {
		if idLimit, ok := vars["idlimit"]; ok {

			registerByPage := 30
			offset, err := strconv.Atoi(idLimit[0])

			if err != nil {
				log.Println("Error : ", err.Error())
			}

			queryComment = queryComment.Where("id BETWEEN ? AND ?", offset-registerByPage, offset)
		}
		queryComment = queryComment.Order("id desc")
	}

	err := queryComment.Find(&comments).Error

	if err != nil {
		m.Code = http.StatusInternalServerError
		m.Message = "Error al guardar los comentarios en formato json"
		commons.DisplayMessage(w, m)
		return
	}

	for i := range comments {
		db.Model(&comments[i]).Related(&comments[i].User)
		comments[i].User[0].Password = ""
		comments[i].Childen = commentGetChildren(comments[i].ID)
		// se busca el voto del usuario en sesion
		votes.CommentID = comments[i].ID
		votes.UserID = user.ID

		count := db.Where(votes).Find(votes).RowsAffected

		if count > 0 {
			if votes.Value {
				comments[i].HasVote = 1
			} else {
				comments[i].HasVote = -1
			}
		}
	}

	j, err := json.Marshal(comments)

	if err != nil {
		m.Code = http.StatusInternalServerError
		m.Message = "Error al convertir los comentarios en json"
		commons.DisplayMessage(w, m)
		return
	}

	if len(comments) > 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	} else {
		m.Code = http.StatusNoContent
		m.Message = "No se encontraron comentarios"
		commons.DisplayMessage(w, m)
	}

}

func commentGetChildren(id uint) (children []models.Comment) {

	db := configuration.GetConnection()
	defer db.Close()

	db.Where("parent_id = ?", id).Find(&children)

	for i := range children {
		db.Model(&children[i]).Related(&children[i].User)
		children[i].User[0].Password = ""
	}

	return children
}
