package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/go-workspace/Comments_Go/configuration"

	"github.com/go-workspace/Comments_Go/commons"
	"github.com/go-workspace/Comments_Go/models"
)

// VoteRegister ... sontrolador para registrar un voto
func VoteRegister(w http.ResponseWriter, r *http.Request) {

	vote := models.Vote{}
	user := models.User{}
	currentVote := models.Vote{}
	m := &models.Message{}

	user, _ = r.Context().Value("user").(models.User)

	err := json.NewDecoder(r.Body).Decode(&vote)

	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = fmt.Sprintf("Error al leer el voto a registrar: %v", err.Error())
		commons.DisplayMessage(w, m)
		return
	}

	vote.UserID = user.ID

	db := configuration.GetConnection()
	defer db.Close()

	db.Where("comment_id = ? and user_id = ?", vote.CommentID, vote.UserID).First(&currentVote)

	// si no existe
	if currentVote.ID == 0 {
		db.Create(&vote)
		log.Println("voto: ", vote.Value)
		err := updateCommentVotes(vote.CommentID, vote.Value, false)

		if err != nil {
			m.Code = http.StatusBadRequest
			m.Message = err.Error()
			commons.DisplayMessage(w, m)
			return
		}
		m.Code = http.StatusCreated
		m.Message = "Voto registrado"
		commons.DisplayMessage(w, m)
		return

	} else if currentVote.Value != vote.Value {

		currentVote.Value = vote.Value
		db.Save(&currentVote)
		err := updateCommentVotes(vote.CommentID, vote.Value, true)

		if err != nil {
			m.Code = http.StatusBadRequest
			m.Message = err.Error()
			commons.DisplayMessage(w, m)
			return
		}

		m.Code = http.StatusOK
		m.Message = "Voto actualizado"
		commons.DisplayMessage(w, m)
		return
	}

	m.Code = http.StatusBadRequest
	m.Message = "Voto ya está registrado"
	commons.DisplayMessage(w, m)
}

func updateCommentVotes(commentID uint, vote bool, isUpdate bool) (err error) {

	comment := models.Comment{}

	db := configuration.GetConnection()
	defer db.Close()

	rows := db.First(&comment, commentID).RowsAffected

	if rows > 0 {

		if vote {
			comment.Votes++
			if isUpdate {
				comment.Votes++
			}
		} else {
			comment.Votes--
			if isUpdate {
				comment.Votes--
			}
		}

		db.Save(&comment)

	} else {
		err = errors.New("No se encontro un registro de comentario para asignar el voto")
	}
	return

}
