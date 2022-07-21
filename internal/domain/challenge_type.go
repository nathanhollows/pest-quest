package domain

import (
	"gorm.io/gorm"
)

// ChallenegeType stores the different types of challenges available
type ChallengeType struct {
	gorm.Model
	Name        string
	Description string
}
