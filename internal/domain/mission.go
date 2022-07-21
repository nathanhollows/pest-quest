package domain

import (
	"database/sql"

	"gorm.io/gorm"
)

// Mission stores the list of quests available (or not) to Players
type Mission struct {
	gorm.Model
	Name        string
	Description string
	Fragment    bool `sql:"DEFAULT:false"`
	Difficulty  int
	Players     int  `sql:"DEFAULT:1"`
	Published   bool `sql:"DEFAULT:false"`
	Start       sql.NullTime
	End         sql.NullTime
	Deadline    sql.NullTime
}
