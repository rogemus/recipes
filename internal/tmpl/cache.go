package tmpl

import (
	"io/fs"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"recipes.krogowski.dev/ui"
)

type TmplCache = map[string]*template.Template

func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006")
}

func avatarName(name string) string {
	return strings.ToUpper(name[:2])
}

func makeSlice(count int) []int {
	return make([]int, count)
}

var functions = template.FuncMap{
	"avatarName": avatarName,
	"humanDate":  humanDate,
	"makeSlice":  makeSlice,
}

func (t *Tmpl) NewCache() error {
	cache := TmplCache{}

	pages, err := fs.Glob(ui.Files, "html/pages/*.tmpl")

	if err != nil {
		return err
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
			return err
		}

		cache[name] = ts
	}

	t.Cache = cache
	return nil
}
