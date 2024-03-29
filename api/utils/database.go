package utils

import (
	"fmt"
	"log"

	"github.com/MHNassar1/Auth/api/models"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql" //mysql database driver
)

func GetConnection(DbUser, DbPassword, DbPort, DbHost, DbName string) (*gorm.DB, error) {
	var err error
	var DB *gorm.DB
	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
	fmt.Println("**************** {DBURL} ***********************")
	DB, err = gorm.Open("mysql", DBURL)
	if err != nil {
		fmt.Printf("****** <<<< Cannot connect to mysql database >>> ****")
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the mysql database")
	}

	migrate(DB)

	return DB, err
}

func migrate(DB *gorm.DB) {

	DB.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").AutoMigrate(
		&models.User{},
		&models.Property{},
		&models.UserProperty{},
		&models.Service{},
		&models.City{})
}
