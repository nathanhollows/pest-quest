package domain

import (
	"gorm.io/gorm"
)

// Advice stores tidbits from Sage for the users during their missions
type Advice struct {
	gorm.Model
	Cover     string
	Md        string
	MissionID int
	Mission   Mission `gorm:"foreignkey:QuestID;references:ID"`
}
