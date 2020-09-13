package seed

import (
	"github.com/MHNassar1/Auth/api/models"
	"log"

	"github.com/jinzhu/gorm"
)

var users = []models.User{
	models.User{
		Nickname: "Mahmoud Nassar",
		Email:    "mnassar662@gmail.com",
		Password: "password",
	},
	models.User{
		Nickname: "test test",
		Email:    "test@gmail.com",
		Password: "password",
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}
}
