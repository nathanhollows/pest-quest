package domain

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/url"
	"regexp"
	"strings"

	"github.com/nathanhollows/pest-quest/internal/helpers"
	"gorm.io/gorm"
)

// Blog stores the blog post content
type Blog struct {
	gorm.Model
	URL       string `gorm:"unique"`
	Title     string
	Content   string
	Author    string
	Published bool `sql:"DEFAULT:false"`
	System    bool `sql:"DEFAULT:false"`
}

func (blog Blog) GetURL() template.HTML {
	url := helpers.URL(fmt.Sprint("blog/", blog.URL))
	return template.HTML(url)
}

func (blog *Blog) GenURL(db *gorm.DB) {
	url := strings.ToLower(blog.Title)
	url = strings.Replace(url, " ", "-", -1)
	reg, err := regexp.Compile("[^a-zA-Z0-9-]+")
	if err != nil {
		log.Fatal(err)
	}
	url = reg.ReplaceAllString(url, "")
	blog.URL = url

	var b []Blog
	res := db.Where("url = ?", url).Find(&b)
	if res.RowsAffected > 0 {
		blog.URL = fmt.Sprint(rand.Intn(10000), "-", blog.URL)
	}
}

func (blog *Blog) Parse(values url.Values, db *gorm.DB) error {
	blog.Title = values.Get("title")
	blog.Author = values.Get("author")
	blog.Content = values.Get("content")
	if blog.URL == "" || (values.Get("title") != values.Get("originalTitle")) {
		blog.GenURL(db)
	}
	if values.Get("publish") == "on" {
		blog.Published = true
	}
	return nil
}
