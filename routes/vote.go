package routes

import (
	"net/http"

	"github.com/go-workspace/Comments_Go/controllers"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

// SetVoteRouter .. ruta de manejo de votos
func SetVoteRouter(router *mux.Router) {

	prefix := "/api/votes"

	subRouter := mux.NewRouter().PathPrefix(prefix).Subrouter().StrictSlash(true)
	subRouter.HandleFunc("/", controllers.VoteRegister).Methods(http.MethodPost)

	router.PathPrefix(prefix).Handler(
		negroni.New(
			negroni.HandlerFunc(controllers.ValidateToken),
			negroni.Wrap(subRouter),
		),
	)

}
