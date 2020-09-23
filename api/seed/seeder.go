package seed

import (
	"github.com/MHNassar1/Auth/api/models"
	"log"

	"github.com/jinzhu/gorm"
)

var users = []models.User{
	//models.User{
	//	UserName: "Mahmoud Nassar",
	//	Email:    "mnassar662@gmail.com",
	//	Password: "password",
	//},
	//models.User{
	//	UserName: "test test",
	//	Email:    "test@gmail.com",
	//	Password: "password",
	//},
}

var props = []models.Properties{
	//models.Properties{
	//	Name:     "personal_info",
	//	ShowName: "Personal Information",
	//	ParentId: 0,
	//},
	//models.Properties{
	//	Name:     "contact_info",
	//	ShowName: "Contacts",
	//	ParentId: 0,
	//},
}

func Load(db *gorm.DB) {

	var err error

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
			continue
		}
	}

	for i, _ := range props {
		err = db.Debug().Model(&models.Properties{}).Create(&props[i]).Error
		if err != nil {
			log.Fatalf("cannot seed Properties table: %v", err)
		}
	}
}
