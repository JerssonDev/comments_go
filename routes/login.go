package routes

import (
	"github.com/go-workspace/Comments_Go/controllers"
	"github.com/gorilla/mux"
)

// SetLoginRouter ... ruta para el login
func SetLoginRouter(router *mux.Router) {

	prefix := "/api/login"

	router.HandleFunc(prefix, controllers.Login)

}
