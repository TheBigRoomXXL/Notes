package server

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"

	_ "github.com/mattn/go-sqlite3"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

var T = &Template{
	templates: template.Must(template.ParseGlob("app/*/templates/*.html")),
}
