package util

import (
	"io"
	"text/template"

	"github.com/labstack/echo/v4"
)

type TemplateRenderer interface {
	Render(w io.Writer, name string, data interface{}, c echo.Context) error
}

// TemplateRenderer is a custom html/template renderer for Echo framework
type templateRenderer struct {
	templates *template.Template
}

func NewTemplateRender() TemplateRenderer {
	return &templateRenderer{
		template.Must(template.ParseGlob("./resources/views/*.html")),
	}
}

// Render renders a template document
func (t *templateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}
	return t.templates.ExecuteTemplate(w, name, data)
}
