package main

import (
	"fmt"
	"net/url"
	"time"

	// New import
	"html/template" // New import
	"path/filepath" // New import

	"letsgo.net/snippetbox/pkg/forms"
	"letsgo.net/snippetbox/pkg/models"
)

type templateData struct {
	CurrentYear int
	Snippet     *models.Snippet
	Snippets    []*models.Snippet
	FormData    url.Values
	FormErrors  map[string]string
	IsAuthenticated bool
	Form        *forms.Form
	Flash       string
	CSRFToken 	string
}

func humanDate(t time.Time) string {
	if t.IsZero(){
		return ""
	}
	return t.UTC().Format("02 Jan 2006 at 15:04")
}

var functions = template.FuncMap{
	"humanDate": humanDate,
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

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
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
