package usecase

import "github.com/joaolima7/cloud_run-goexpert/internal/domain/weather/repository"

type WeatherInputDTO struct {
	City string `json:"city"`
}

type WeatherOutputDTO struct {
	CelsiusWeather    float32 `json:"temp_C"`
	FahrenheitWeather float32 `json:"temp_F"`
	KelvinWeather     float32 `json:"temp_K"`
}

type GetWeatherByCityUseCase struct {
	GetWeatherByCityRepository repository.GetWeatherByCityRepository
}

func NewGetWeatherByCityUseCase(getWeatherByCityRepository repository.GetWeatherByCityRepository) *GetWeatherByCityUseCase {
	return &GetWeatherByCityUseCase{GetWeatherByCityRepository: getWeatherByCityRepository}
}

func (uc *GetWeatherByCityUseCase) Execute(input WeatherInputDTO) (*WeatherOutputDTO, error) {
	weather, err := uc.GetWeatherByCityRepository.GetWeatherByCity(input.City)
	if err != nil {
		return nil, err
	}

	return &WeatherOutputDTO{
		CelsiusWeather:    weather.CelsiusWeather,
		FahrenheitWeather: weather.FahrenheitWeather,
		KelvinWeather:     weather.KelvinWeather,
	}, nil
}
