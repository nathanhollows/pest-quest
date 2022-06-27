package blog

import (
	"github.com/nathanhollows/pest-quest/internal/domain"
	"gorm.io/gorm"
)

func GetLatest(db gorm.DB, count int) []domain.Blog {
	latest := []domain.Blog{}
	db.Limit(count).Order("id DESC").Find(&latest)
	return latest
}
