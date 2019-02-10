package routes

import (
	"net/http"

	"github.com/urfave/negroni"

	"github.com/go-workspace/Comments_Go/controllers"
	"github.com/gorilla/mux"
)

// SetUserRouter ... ruta para crear el usuario
func SetUserRouter(router *mux.Router) {

	prefix := "/api/users"

	subRouter := mux.NewRouter().PathPrefix(prefix).Subrouter().StrictSlash(true)

	subRouter.HandleFunc("/", controllers.UserCreate).Methods(http.MethodPost)

	router.PathPrefix(prefix).Handler(
		negroni.New(
			negroni.Wrap(subRouter),
		),
	)

}
