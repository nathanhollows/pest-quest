package domain

import (
	"net/url"
	"strconv"

	"gorm.io/gorm"
)

// Advice stores tidbits from Sage for the users during their missions
type Advice struct {
	gorm.Model
	Author    string
	Cover     string
	Content   string
	MissionID int
	Mission   Mission `gorm:"foreignkey:MissionID;references:ID"`
}

func (advice *Advice) Parse(values url.Values, db *gorm.DB) error {
	var err error
	advice.Author = values.Get("author")
	advice.Cover = values.Get("cover")
	advice.Content = values.Get("content")
	advice.MissionID, err = strconv.Atoi(values.Get("missionID"))
	return err
}
