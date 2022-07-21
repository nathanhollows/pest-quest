package domain

import (
	"gorm.io/gorm"
)

// Narrative stores the different narrative pieces that players can collect
type NarrativeType struct {
	gorm.Model
	Name        string
	Description string
}
