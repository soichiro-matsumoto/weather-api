package main

import (
	"html/template"
	"io"
	"net/http"
	"strconv"
	"weather-api/config"
	"weather-api/handlers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	// Config ...
	Config = config.NewConfig()
)

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {

	// Echoのインスタンス作る
	e := echo.New()

	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("resources/templates/*.html")),
	}
	e.Renderer = renderer

	// Named route "foobar"
	e.GET("", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", map[string]interface{}{
			"name": "Dolly!",
		})
	})

	// 全てのリクエストで差し込みたいミドルウェア（ログとか）はここ
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルーティング
	e.GET("/weather", handlers.GetWeather())

	// サーバー起動
	e.Start(":" + strconv.Itoa(Config.Server.Port))
}
