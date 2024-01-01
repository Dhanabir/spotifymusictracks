package models

import (
	"gorm.io/gorm"
)

type Track struct {
	ISRC  string `gorm:"unique"`
	Title string
}

type TrackModel struct {
	DB *gorm.DB
}

// insert data into track table
func (t *TrackModel) Insert(isrc, title string) error {
	newTrack := &Track{ISRC: isrc, Title: title}
	result := t.DB.Create(newTrack)
	return result.Error
}

// get track data using isrc
func (t *TrackModel) GetTrackTitle(isrc string) string {
	newTrack := &Track{}
	result := t.DB.Where("isrc=?", isrc).Find(newTrack)
	if result.RowsAffected == 0 {
		return ""
	}
	return newTrack.Title
}
