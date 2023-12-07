package main

import (
	"echo-todoapp/internal/handler"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()

	// Serve Static files
	e.Static("/assets", "internal/assets")

	// Middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Templates configuration
	t := &Template{
		templates: template.Must(template.ParseGlob("internal/public/views/*.html")),
	}
	e.Renderer = t

	// Routes
	e.GET("/", handler.Index)

	e.Logger.Fatal(e.Start(":1234"))

}
