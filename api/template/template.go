package template

import (
	"html/template"
	"isPalindrom/word"
)

type TemplateData struct {
	Template *template.Template
	Rows     []word.Word
}

// GetTemplate returns a form template
func GetTemplate(words []word.Word) (TemplateData, error) {
	data := TemplateData{}
	tmpl, err := template.ParseFiles("template/form.html")

	if err == nil {
		data.Template = tmpl

		data.Rows = words
	}

	return data, err
}
