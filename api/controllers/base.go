package controllers

import (
	"fmt"
	"github.com/MHNassar1/Auth/api/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"    //mysql database driver

)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize(DbUser, DbPassword, DbPort, DbHost, DbName string) {

	var err error

	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
	fmt.Println("**************** {DBURL} ***********************")
	server.DB, err = gorm.Open("mysql", DBURL)
	if err != nil {
		fmt.Printf("****** <<<< Cannot connect to mysql database >>> ****")
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the mysql database")
	}


	server.DB.Debug().AutoMigrate(&models.User{}) //database migration

	server.Router = mux.NewRouter()

	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
