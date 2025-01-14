package renderer

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

var funcMap = template.FuncMap{}

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func NewTemplateRenderer() echo.Renderer {
	// t := template.New("")
	// t = t.Funcs(funcMap)
	// t = t.Option("missingkey=error")
	// t = template.Must(t.ParseGlob("templates/*.html"))

	return &TemplateRenderer{
		templates: template.Must(template.New("").Funcs(funcMap).Option("missingkey=error").ParseGlob("templates/*.html")),
	}
}
