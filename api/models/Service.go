package models

import "github.com/MHNassar1/Auth/api/core"

type Service struct {
	ID       uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Name     string `gorm:"size:255;not null;unique" json:"name"`
	ShowName string `gorm:"size:255;not null;unique" json:"show_name"`
	ImageUrl string `gorm:"size:255;not null;unique" json:"image_url"`
	Status   int    `gorm:"size:255; null;" json:"status"`
}

func (p *Service) FindAl() (*[]Service, error) {
	var err error
	var services []Service
	err = core.AppInstance.DB.Debug().Model(&Service{}).Where("status = 1").Find(&services).Error
	if err != nil {
		return &[]Service{}, err
	}
	return &services, err
}
