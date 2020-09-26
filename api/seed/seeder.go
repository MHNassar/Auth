package seed

import (
	"github.com/MHNassar1/Auth/api/models"
	"log"

	"github.com/jinzhu/gorm"
)

var users = []models.User{
	models.User{
		UserName: "Mahmoud Nassar",
		Email:    "mnassar662@gmail.com",
		Password: "password",
	},
	models.User{
		UserName: "test test",
		Email:    "test@gmail.com",
		Password: "password",
	},
}

var props = []models.Property{
	models.Property{
		Name:     "personal_info",
		ShowName: "Personal Information",
		ParentId: 0,
	},
	models.Property{
		Name:     "contact_info",
		ShowName: "Contacts",
		ParentId: 0,
	},
	models.Property{
		Name:     "first_name",
		ShowName: "First Name",
		ParentId: 1,
	},
	models.Property{
		Name:     "last_name",
		ShowName: "Last Name",
		ParentId: 1,
	},
}

func Load(db *gorm.DB) {

	var err error

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}

	for i, _ := range props {
		err = db.Debug().Model(&models.Property{}).Create(&props[i]).Error
		if err != nil {
			log.Fatalf("cannot seed Properties table: %v", err)
		}
	}
}
