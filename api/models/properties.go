package models

type Properties struct {
	ID       uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Name     string `gorm:"size:255;not null;unique" json:"name"`
	ShowName string `gorm:"size:255;not null;unique" json:"show_name"`
	ParentId int    `gorm:"size:255;not null;" json:"parent id"`
}
