package weather

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/joaolima7/cloud_run-goexpert/internal/domain/weather"
	httpclient "github.com/joaolima7/cloud_run-goexpert/internal/infra/http_client"
	"github.com/joaolima7/cloud_run-goexpert/internal/utils"
)

type GetWeatherByCityRepositoryImpl struct {
	client httpclient.HTTPClient
	apiKey string
}

func NewGetWeatherByCityRepositoryImpl(client httpclient.HTTPClient, apiKey string) *GetWeatherByCityRepositoryImpl {
	return &GetWeatherByCityRepositoryImpl{
		client: client,
		apiKey: apiKey,
	}
}

func (r *GetWeatherByCityRepositoryImpl) GetWeatherByCity(city string) (*weather.Weather, error) {
	escapedCity := url.QueryEscape(city)
	url := "https://api.openweathermap.org/data/2.5/weather?q=" + escapedCity + "&appid=" + r.apiKey + "&units=metric"
	ctx := context.Background()
	data, err := r.client.Get(ctx, url)
	if err != nil {
		return nil, err
	}

	var celsiusWeather struct {
		Main struct {
			Temp float64 `json:"temp"`
		} `json:"main"`
	}

	if err := json.Unmarshal(data, &celsiusWeather); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	if celsiusWeather.Main.Temp == 0 {
		return nil, fmt.Errorf("weather data not found for city: %s", city)
	}

	return &weather.Weather{
		CelsiusWeather:    float32(celsiusWeather.Main.Temp),
		FahrenheitWeather: utils.ConvertCelsiusToFahrenheit(float32(celsiusWeather.Main.Temp)),
		KelvinWeather:     utils.ConvertCelsiusToKelvin(float32(celsiusWeather.Main.Temp)),
	}, nil
}
