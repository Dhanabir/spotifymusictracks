package models

import (
	"gorm.io/gorm"
)

type Artist struct {
	Name string
	ISRC string
}
type ArtistModel struct {
	DB *gorm.DB
}

// insert data into artist table
func (a *ArtistModel) Insert(artistname []string, isrc string) error {
	var artistsList []Artist
	for i := 0; i < len(artistname); i++ {
		artistsList = append(artistsList, Artist{Name: artistname[i], ISRC: isrc})
	}
	result := a.DB.Create(&artistsList)
	return result.Error
}

// get artist names by isrc value
func (a *ArtistModel) GetArtistNamesByISRC(isrc string) []string {
	var artistsList []Artist
	result := a.DB.Where("isrc=?", isrc).Find(&artistsList)
	if result.RowsAffected == 0 {
		return []string{}
	}

	var artistNames []string
	for _, v := range artistsList {
		artistNames = append(artistNames, v.Name)
	}

	return artistNames
}

// get isrc values by artist name
func (a *ArtistModel) GetISRCByName(name string) []string {
	var artistsList []Artist
	result := a.DB.Where("name like ?", "%"+name+"%").Find(&artistsList)
	if result.RowsAffected == 0 {
		return []string{}
	}

	var isrcList []string
	for _, v := range artistsList {
		isrcList = append(isrcList, v.ISRC)
	}
	return isrcList
}
