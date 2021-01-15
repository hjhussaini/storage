package models

import "github.com/jinzhu/gorm"

// swagger:model
type Upload struct {
	// The type of storage
	// required: false
	Storage string `json:"-" gorm:"column:storage"`

	// The path of source file
	// required: true
	Source string `json:"src" validate:"required" gorm:"column:src"`

	// The path of destination
	// required: true
	Destination string `json:"dst" validate:"required,path" gorm:"column:dst"`

	// Whether remove source after uploading or not
	// required: false
	RemoveSource bool `json:"remove" gorm:"column:remove"`
}

func (upload *Upload) Save(database *gorm.DB) {
	database.Create(upload)
}

func (upload *Upload) Delete(database *gorm.DB) {
	database.Delete(upload)
}

type Uploads []Upload

func (uploads *Uploads) Fetch(database *gorm.DB, storage string) {
	database.Where("storage = ?", storage).Find(&uploads)
}
