package models

import (
	"github.com/MHNassar1/Auth/api/core"
)

type UserProperty struct {
	ID         uint32   `gorm:"primary_key;auto_increment" json:"id"`
	UserId     uint32   `gorm:"size:10;not null" json:"user_id"`
	PropertyID int      `gorm:"size:10;not null" json:"property_id"`
	Value      string   `gorm:"size:1200;not null" json:"value"`
	Property   Property `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (u *UserProperty) SaveProperty() error {
	var err error
	err = core.AppInstance.DB.Debug().Create(&u).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *UserProperty) UpdateProperty() error {

	db := core.AppInstance.DB.Debug().Model(&User{}).Where("user_id = ? and property_id = ?", u.UserId, u.PropertyID).Take(&UserProperty{}).UpdateColumns(
		map[string]interface{}{
			"value": u.Value,
		},
	)
	if db.Error != nil {
		return db.Error
	}
	return nil

}
