package repository

import "github.com/joaolima7/cloud_run-goexpert/internal/domain/weather"

type GetWeatherByCityRepository interface {
	GetWeatherByCity(city string) (*weather.Weather, error)
}
