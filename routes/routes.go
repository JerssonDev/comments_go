package routes

import (
	"github.com/gorilla/mux"
)

// InitRoutes ... unifica todas las rutas aqui
func InitRoutes() *mux.Router {

	router := mux.NewRouter().StrictSlash(false)

	SetUserRouter(router)
	SetLoginRouter(router)
	SetCommentRouter(router)
	SetVoteRouter(router)
	SetRealTimeRouter(router)
	SetPublicRouter(router)

	return router
}
