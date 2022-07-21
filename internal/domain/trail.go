package domain

import (
	"fmt"
	"log"
	"math/rand"
	"regexp"
	"strings"

	"gorm.io/gorm"
)

// Trail defines a collection of markers and challenges
type Trail struct {
	gorm.Model
	URL         string `gorm:"unique"`
	Title       string
	Description string
	Author      string
	UserID      int
	User        User `gorm:"foreignkey:UserID;references:ID"`
	Published   bool `sql:"DEFAULT:false"`
}

func (t *Trail) BeforeCreate(db *gorm.DB) (err error) {
	url := strings.ToLower(t.Title)
	url = strings.Replace(url, " ", "-", -1)
	reg, err := regexp.Compile("[^a-zA-Z0-9-]+")
	if err != nil {
		log.Fatal(err)
	}
	url = reg.ReplaceAllString(url, "")
	t.URL = url

	var c []Trail
	res := db.Where("url = ?", url).Find(&c)
	if res.RowsAffected > 0 {
		t.URL = fmt.Sprint(rand.Intn(10000), "-", t.URL)
	}

	return nil
}
