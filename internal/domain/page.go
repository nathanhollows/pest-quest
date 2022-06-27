package domain

import (
	"gorm.io/gorm"
)

// Page stores a static page that can be accessed via QR code
type Page struct {
	gorm.Model
	URL       string `gorm:"unique"`
	Title     string
	Content   string
	Author    string
	Published bool `sql:"DEFAULT:false"`
	System    bool `sql:"DEFAULT:false"`
}
