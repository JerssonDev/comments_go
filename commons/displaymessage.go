package commons

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-workspace/Comments_Go/models"
)

// DisplayMessage ... devuelme un mensaje al cliente
func DisplayMessage(w http.ResponseWriter, m *models.Message) {

	j, err := json.Marshal(m)

	if err != nil {
		log.Fatalf("Error al convertir el mensaje: %v", err.Error())
	}

	w.WriteHeader(m.Code)
	w.Write(j)
}
