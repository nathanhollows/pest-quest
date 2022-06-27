package domain

import (
	"fmt"
	"html/template"

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
