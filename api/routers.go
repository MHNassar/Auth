package api

import (
	"github.com/MHNassar1/Auth/api/controllers"
	"github.com/MHNassar1/Auth/api/middlewares"
	"github.com/gorilla/mux"
)

func Register() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", controllers.Home)
	router.HandleFunc("/login", middlewares.SetMiddlewareJSON(controllers.Login))

	router.HandleFunc("/users", middlewares.SetMiddlewareJSON(controllers.GetUsers)).Methods("GET")
	router.HandleFunc("/users", middlewares.SetMiddlewareJSON(controllers.CreateUser)).Methods("POST")
	router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(controllers.GetUser)).Methods("GET")
	router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(controllers.UpdateUser))).Methods("PUT")
	router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuthentication(controllers.DeleteUser)).Methods("DELETE")

	router.HandleFunc("/properties", middlewares.SetMiddlewareJSON(controllers.GetAllProperties)).Methods("GET")
	router.HandleFunc("/users/properties", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(controllers.AddProperties))).Methods("POST")
	router.HandleFunc("/users/properties/edit", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(controllers.UpdateProperties))).Methods("POST")

	router.HandleFunc("/init/settings", middlewares.SetMiddlewareJSON(controllers.GetInitSettings)).Methods("GET")

	return router
}
