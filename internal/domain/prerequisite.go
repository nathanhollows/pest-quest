package domain

import (
	"gorm.io/gorm"
)

// Prerequisite stores a list of blocking quests
// This list should be satisfied before any given quest is unlocked
type Prerequisite struct {
	gorm.Model
	LockedID  int
	Locked    Mission `gorm:"foreignkey:LockedID;references:ID"`
	DependsID int
	Depends   Mission `gorm:"foreignkey:LockedID;references:ID"`
}
