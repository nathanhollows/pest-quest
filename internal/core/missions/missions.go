package missions

import (
	"github.com/nathanhollows/pest-quest/internal/domain"
	"gorm.io/gorm"
)

func Get(db *gorm.DB) []domain.Mission {
	missions := []domain.Mission{}
	db.Model(domain.Mission{}).Find(&missions)
	return missions
}
