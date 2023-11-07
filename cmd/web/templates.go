package main

import (
	"html/template"
	"path/filepath"
	"snippetbox.xuche.net/internal/models"
)

type templateData struct {
	Snippet *models.Snippet
	Snippets []*models.Snippet
}

func newTemplateCache() (map[string] *template.Template, error) {
	cache := map[string] *template.Template{}

	pages, err := filepath.Glob("./ui/html/pages/*.tmpl.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		files := []string{
			"./ui/html/base.tmpl.html",
			"./ui/html/partials/nav.tmpl.html",
			page,
		}

		ts, err := template.ParseFiles(files...)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}