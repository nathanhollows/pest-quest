package markers

import (
	"github.com/nathanhollows/pest-quest/internal/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetPublished(db gorm.DB) []domain.Marker {
	var markers []domain.Marker
	db.Where("published = ?", true).Preload(clause.Associations).Find(&markers)
	return markers
}
