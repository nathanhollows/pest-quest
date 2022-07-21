package domain

import (
	"gorm.io/gorm"
)

// Element keeps track of the different components in each mission
type Element struct {
	gorm.Model
	MissionID   int     `gorm:"index:idx_mission_builder"`
	Mission     Mission `gorm:"foreignkey:MissionID;references:ID"`
	Order       int     `gorm:"index:idx_mission_builder"`
	ChallengeID int
	NarrativeID int
	Narrative   Narrative `gorm:"foreignkey:NarrativeID;references:ID"`
	LocationID  int
	Location    Location `gorm:"foreignkey:LocationID;references:ID"`
	AdviceID    int
	Advice      Advice `gorm:"foreignkey:AdviceID;references:ID"`
}
