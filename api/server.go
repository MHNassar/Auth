package api

import (
	"fmt"
	"github.com/MHNassar1/Auth/api/core"
	"github.com/MHNassar1/Auth/api/seed"
	"github.com/MHNassar1/Auth/api/utils"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var app = &core.App{}

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("sad .env file found")
	}
}

func Run() {

	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	} else {
		fmt.Println("************** We are getting the env values ***************")
	}

	db, err := utils.GetConnection(os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	router := Register()

	app.InitApp(db, router)
	seed.Load(db)

	fmt.Println("************** App Start <<>>>> ***************")

	app.Run(":8080")

}
