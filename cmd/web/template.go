package main

import (
	"html/template"
	"io/fs"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/justinas/nosurf"
	"recipies.krogowski.dev/internal/models"
	"recipies.krogowski.dev/ui"
)

type templateData struct {
	CurrentYear     int
	Recipe          models.Recipe
	Recipies        []models.Recipe
	Form            any
	Flash           string
	IsAuthenticated bool
	CSRFToken       string
	UserName        string
}

func (app *application) newTemplateData(r *http.Request) templateData {
	return templateData{
		CurrentYear:     time.Now().Year(),
		Flash:           app.sessionManager.PopString(r.Context(), "flash"),
		IsAuthenticated: app.isAuthenticated(r),
		CSRFToken:       nosurf.Token(r),
		UserName:        app.sessionUserName(r),
	}
}

func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006")
}

func avatarName(name string) string {
	return strings.ToUpper(name[:2])
}

var functions = template.FuncMap{
	"avatarName": avatarName,
	"humanDate":  humanDate,
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := fs.Glob(ui.Files, "html/pages/*.tmpl")

	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		patterns := []string{
			"html/base.tmpl",
			"html/partials/*.tmpl",
			page,
		}

		ts, err := template.New(name).Funcs(functions).ParseFS(ui.Files, patterns...)

		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}
	return cache, nil
}
