package markers

import (
	"github.com/nathanhollows/pest-quest/internal/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetPublished(db gorm.DB) []domain.Location {
	var markers []domain.Location
	db.Where("published = ?", true).Preload(clause.Associations).Find(&markers)
	return markers
}
