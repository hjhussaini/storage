package models

import "github.com/jinzhu/gorm"

// swagger:model
type Login struct {
	// The type of storage
	// required: false
	Storage string `json:"-" gorm:"column:storage"`

	// The application token
	// required: true
	Token string `json:"token" validate:"required"`
}

func (login *Login) Fetch(database *gorm.DB) {
	database.Where("storage = ?", login.Storage).Find(login)
}

func (login *Login) Save(database *gorm.DB) {
	login.Delete(database)
	database.Create(login)
}

func (login *Login) Delete(database *gorm.DB) {
	var row Login

	database.Where("storage = ?", login.Storage).Find(&row)
	database.Delete(&row)
}
