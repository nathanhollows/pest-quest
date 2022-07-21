package domain

import (
	"gorm.io/gorm"
)

// LocationTag stores each tag for a location
type LocationTag struct {
	gorm.Model
	Tag        string
	LocationID int
	Location   Location `gorm:"foreignkey:LocationID;references:ID"`
}
