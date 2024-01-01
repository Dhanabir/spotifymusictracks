package models

import (
	"gorm.io/gorm"
)

type Image struct {
	URI  string
	ISRC string `gorm:"unique"`
}

type ImageModel struct {
	DB *gorm.DB
}

// insert data into image table
func (i *ImageModel) Insert(imageURI, isrc string) error {
	newImage := &Image{URI: imageURI, ISRC: isrc}
	result := i.DB.Create(newImage)
	return result.Error
}

// retrieve image uri using isrc value
func (i *ImageModel) GetImageURI(isrc string) string {
	newImage := &Image{}
	result := i.DB.Where("isrc=?", isrc).Find(newImage)
	if result.RowsAffected == 0 {
		return ""
	}
	return newImage.URI
}
