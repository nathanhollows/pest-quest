package markdown

import (
	"html/template"
)

func MDtoHTML(md string) template.HTML {
	type markdown struct {
		MD string `json:"md"`
	}
	return ""
}
