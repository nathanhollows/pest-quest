package domain

import (
	"net/url"
	"strconv"

	"gorm.io/gorm"
)

// Location stores each map marker in the game
type Location struct {
	gorm.Model
	Lng       string
	Lat       string
	Text      string
	Name      string
	URL       string
	Published bool `sql:"DEFAULT:false"`
	TypeID    int
	Type      MarkerType `gorm:"foreignkey:TypeID;references:ID"`
}

func (location *Location) Parse(values url.Values, db *gorm.DB) error {
	var err error
	location.Name = values.Get("name")
	location.Lng = values.Get("lng")
	location.Lat = values.Get("lat")
	location.URL = values.Get("url")
	location.Text = values.Get("text")
	if values.Get("publish") == "on" {
		location.Published = true
	}
	location.TypeID, err = strconv.Atoi(values.Get("type"))
	if err != nil {
		return err
	}
	return nil
}
