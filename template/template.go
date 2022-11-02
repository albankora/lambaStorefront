package template

import (
	"bytes"
	"html/template"
	"log"
)

type Head struct {
	Lang            string
	PageTitle       string
	MetaDescription string
}

func Load(view string, data interface{}) (string, error) {
	theme := "base"
	file := "themes/" + theme + "/pages/" + view + ".html"
	layout := "themes/" + theme + "/layouts/base.html"
	header := "themes/" + theme + "/partials/header.html"
	footer := "themes/" + theme + "//partials/footer.html"
	head := "themes/" + theme + "/partials/head.html"

	t, _ := template.ParseFiles(file, layout, header, footer, head)

	var html bytes.Buffer
	if err := t.ExecuteTemplate(&html, "base", data); err != nil {
		log.Panic(err)
		return "", err
	}

	return html.String(), nil
}