package domain

import (
	"gorm.io/gorm"
)

// Challenge stores information for the different type of challenges... I think
type Challenge struct {
	gorm.Model
	Name   string
	Desc   string
	TypeID int
	Type   ChallengeType `gorm:"foreignkey:TypeID;references:ID"`
}
