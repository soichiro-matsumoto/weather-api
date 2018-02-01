package main

import (
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

func main() {

	// Echoのインスタンス作る
	e := echo.New()

	// 全てのリクエストで差し込みたいミドルウェア（ログとか）はここ
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルーティング
	e.GET("/weather", handlers.GetWeather())

	// サーバー起動
	e.Start(":" + strconv.Itoa(Config.Server.Port))
}
