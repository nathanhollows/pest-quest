package domain

import (
	"gorm.io/gorm"
)

// Narrative stores the different narrative pieces that players can collect
type Narrative struct {
	gorm.Model
	Name       string
	Md         string
	Media      string
	CoverImage string
	Author     string
	TypeID     int
	Type       NarrativeType `gorm:"foreignkey:TypeID;references:ID"`
}
