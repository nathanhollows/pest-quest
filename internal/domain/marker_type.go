package domain

import (
	"gorm.io/gorm"
)

// MarkerType stores the types of marker that a marker can be
type MarkerType struct {
	gorm.Model
	Name  string
	Icon  string
	Color string
}
