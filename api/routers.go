package api

import (
	"github.com/gorilla/mux"
	"github.com/MHNassar1/Auth/api/controllers"
	"github.com/MHNassar1/Auth/api/middlewares"
)
func Register() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/",controllers.Home)
	router.HandleFunc("/login",middlewares.SetMiddlewareJSON(controllers.Login))

	return router
}