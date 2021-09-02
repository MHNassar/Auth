package api

import (
	"fmt"
	"log"
	"os"

	"github.com/MHNassar1/Auth/api/core"
	"github.com/MHNassar1/Auth/api/utils"
	"github.com/joho/godotenv"
)

var app = &core.App{}

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("sad .env file found")
	}
}

func Run() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	} else {
		fmt.Println("************** We are getting the env values ***************")
	}

	db, _ := utils.GetConnection(os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	router := Register()

	app.InitApp(db, router)

	fmt.Println("************** App Start <<>>>> ***************")

	app.Run(":8080")

}
