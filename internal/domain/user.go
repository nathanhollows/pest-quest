package domain

import "gorm.io/gorm"

// User represents both players and admins
type User struct {
	gorm.Model
	DisplayName string `json:"displayname"`
	Email       string `gorm:"unique" json:"email"`
	Password    string `json:"password"`
	Admin       bool   `json:"admin" sql:"DEFAULT:false"`
}

func (user *User) FindByID(db *gorm.DB, id string) {
	db.Where("id = ?", id).Find(&user)
}

func (user *User) FindByEmail(db *gorm.DB, email string) {
	db.Where("email = ?", email).Find(&user)
}
