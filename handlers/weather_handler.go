package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

type WeatherContext struct {
	*BaseContext
}

func NewWeatherContext(c echo.Context) *WeatherContext {
	return &WeatherContext{
		NewBaseContext(c),
	}
}

func GetWeather() echo.HandlerFunc {
	return func(c echo.Context) error {
		ct := NewWeatherContext(c)
		return ct.JSON(http.StatusOK, &Weather{})
	}
}

type Weather struct {
}
