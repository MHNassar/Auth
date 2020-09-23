package core

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)

type App struct {
	DB     *gorm.DB
	Router *mux.Router
}

var AppInstance *App = nil

func (app *App) InitApp(DB *gorm.DB, Router *mux.Router) {
	app.DB = DB
	app.Router = Router

	AppInstance = app
}

func (App *App) Run(addr string) {
	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(addr, App.Router))
}

func (App *App) GetDB() *gorm.DB {
	return App.DB
}
