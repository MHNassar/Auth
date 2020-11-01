package models

import "github.com/MHNassar1/Auth/api/core"

type City struct {
	ID       uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Name     string `gorm:"size:255;not null;unique" json:"name"`
	ShowName string `gorm:"size:255;not null;unique" json:"show_name"`
	ImageUrl string `gorm:"size:255;not null;unique" json:"image_url"`
	Status   int    `gorm:"size:255; null;" json:"status"`
}

func (p *City) FindAl() (*[]City, error) {
	var err error
	var cities []City
	err = core.AppInstance.DB.Debug().Model(&City{}).Where("status = 1").Find(&cities).Error
	if err != nil {
		return &[]City{}, err
	}
	return &cities, err
}
