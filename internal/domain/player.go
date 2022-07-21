package domain

import (
	"gorm.io/gorm"
)

// Player stores the player information
type Player struct {
	gorm.Model
	Name        string
	DisplayName string
	Email       string
	Consent     bool         `sql:"DEFAULT:false"`
	MissionLog  []MissionLog `gorm:"foreignkey:ID;references:PlayerID"`
	Published   bool         `sql:"DEFAULT:false"`
	System      bool         `sql:"DEFAULT:false"`
}
