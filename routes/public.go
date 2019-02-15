package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// SetPublicRouter ...mostrara los archivos staticos
func SetPublicRouter(router *mux.Router) {

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./public/index.html")
	})

	router.HandleFunc("/example", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./public/comments.html")
	})

}
