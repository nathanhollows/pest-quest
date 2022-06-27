package domain

import (
	"net/url"
	"strconv"

	"gorm.io/gorm"
)

// Marker stores each map marker in the game
type Marker struct {
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

func (marker *Marker) Parse(values url.Values, db *gorm.DB) error {
	var err error
	marker.Name = values.Get("name")
	marker.Lng = values.Get("lng")
	marker.Lat = values.Get("lat")
	marker.URL = values.Get("url")
	marker.Text = values.Get("text")
	if values.Get("publish") == "on" {
		marker.Published = true
	}
	marker.TypeID, err = strconv.Atoi(values.Get("type"))
	if err != nil {
		return err
	}
	return nil
}
