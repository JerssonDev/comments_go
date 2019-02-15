package routes

import (
	"net/http"

	"github.com/urfave/negroni"

	"github.com/go-workspace/Comments_Go/controllers"
	"github.com/gorilla/mux"
)

// SetCommentRouter ...
func SetCommentRouter(router *mux.Router) {

	prefix := "/api/comments"

	subRouter := mux.NewRouter().PathPrefix(prefix).Subrouter().StrictSlash(true)
	subRouter.HandleFunc("/", controllers.CommentCreate).Methods(http.MethodPost)
	subRouter.HandleFunc("/", controllers.CommentGetAll).Methods(http.MethodGet)

	router.PathPrefix(prefix).Handler(
		negroni.New(
			negroni.HandlerFunc(controllers.ValidateToken),
			negroni.Wrap(subRouter),
		),
	)
}
