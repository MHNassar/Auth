package models

type Property struct {
	ID       uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Name     string `gorm:"size:255;not null;unique" json:"name"`
	ShowName string `gorm:"size:255;not null;unique" json:"show_name"`
	Type     string `gorm:"size:255;null;" json:"type"`
	ParentId int    `gorm:"size:255; null;" json:"parent id"`
}
