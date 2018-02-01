package config

import "github.com/BurntSushi/toml"

type Server struct {
	Port int `toml:"port"`
}

type OpenWeatherMap struct {
	BaseUrl string `toml:"base_url"`
	ApiKey  string `toml:"api_key"`
}

type Config struct {
	Server         Server
	OpenWeatherMap OpenWeatherMap
}

func NewConfig() *Config {
	var config Config
	_, err := toml.DecodeFile("config.toml", &config)
	if err != nil {
		panic(err)
	}
	return &config
}
