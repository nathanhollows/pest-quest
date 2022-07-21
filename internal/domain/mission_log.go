package domain

import (
	"time"

	"gorm.io/gorm"
)

// MissionLog tracks Quest progress for Teams and Players
type MissionLog struct {
	gorm.Model
	Session   string `gorm:"index"`
	PlayerID  int    `gorm:"index"`
	Player    Player `gorm:"foreignkey:PlayerID;references:ID"`
	StartTime time.Time
	EndTime   time.Time
	Active    bool `sql:"DEFAULT:true"`
	MissionID int
	Mission   Mission `gorm:"foreignkey:MissionID;references:ID"`
}
