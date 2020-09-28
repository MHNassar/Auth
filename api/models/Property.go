package models

import "github.com/MHNassar1/Auth/api/core"

type Property struct {
	ID       uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Name     string `gorm:"size:255;not null;unique" json:"name"`
	ShowName string `gorm:"size:255;not null;unique" json:"show_name"`
	Type     string `gorm:"size:255;null;" json:"type"`
	ParentId int    `gorm:"size:255; null;" json:"parent_id"`

	SubProperty []Property `gorm:"foreignkey:ParentId"`
}

func (p *Property) FindAl() (*[]Property, error) {
	var err error
	var props []Property
	err = core.AppInstance.DB.Debug().Preload("SubProperty").Model(&Property{}).Where("parent_id = 0").Find(&props).Error
	if err != nil {
		return &[]Property{}, err
	}
	return &props, err
}
