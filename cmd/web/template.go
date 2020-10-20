package main

import (
	"fmt"
	// New import
	"html/template" // New import
	"path/filepath" // New import

	"letsgo.net/snippetbox/pkg/models"
)

type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}

func newTemplateCache(dir string) (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob(filepath.Join(dir, "*.page.tmpl"))
	fmt.Println(pages, "test page")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		fmt.Println(name, "name")
		fmt.Println(page, "test page 29")
		ts, err := template.ParseFiles(page)
		if err != nil {
			return nil, err
		}
		fmt.Println(filepath.Join(dir, "*.layout.tmpl"), "test template 33")
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.tmpl"))

		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.tmpl"))
		if err != nil {
			return nil, err
		}
		cache[name] = ts
		fmt.Println(ts, "Cache name41")
	}
	fmt.Println(cache, "test cache")
	return cache, nil
}
