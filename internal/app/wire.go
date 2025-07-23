//go:build wireinject
// +build wireinject

package app

import (
	"time"

	"github.com/google/wire"
	"github.com/joaolima7/cloud_run-goexpert/config"
	"github.com/joaolima7/cloud_run-goexpert/internal/domain/cep/repository"
	"github.com/joaolima7/cloud_run-goexpert/internal/domain/cep/usecase"
	weatherRepository "github.com/joaolima7/cloud_run-goexpert/internal/domain/weather/repository"
	weatherUsecase "github.com/joaolima7/cloud_run-goexpert/internal/domain/weather/usecase"
	cepRepo "github.com/joaolima7/cloud_run-goexpert/internal/infra/external/cep"
	weatherRepo "github.com/joaolima7/cloud_run-goexpert/internal/infra/external/weather"
	httpclient "github.com/joaolima7/cloud_run-goexpert/internal/infra/http_client"
)

// --- Providers ---

func provideConfig() (*config.Config, error) {
	return config.LoadConfig()
}

func provideWeatherAPIKey(cfg *config.Config) string {
	return cfg.WeatherKey
}

func provideHTTPClient() httpclient.HTTPClient {
	return httpclient.NewHTTPClientImpl(10 * time.Second)
}

// --- Repositories ---
func provideCepRepository(client httpclient.HTTPClient) repository.GetCityByCepRepository {
	return cepRepo.NewGetCityByCepRepositoryImpl(client)
}

func provideWeatherRepository(client httpclient.HTTPClient, apiKey string) weatherRepository.GetWeatherByCityRepository {
	return weatherRepo.NewGetWeatherByCityRepositoryImpl(client, apiKey)
}

// --- Use Cases ---
func provideGetCityUseCase(repo repository.GetCityByCepRepository) *usecase.GetCityByCepUseCase {
	return usecase.NewGetCityByCepUseCase(repo)
}

func provideGetWeatherByCityUseCase(repo weatherRepository.GetWeatherByCityRepository) *weatherUsecase.GetWeatherByCityUseCase {
	return weatherUsecase.NewGetWeatherByCityUseCase(repo)
}

// --- Sets ---

var (
	configSet = wire.NewSet(provideConfig)
	clientSet = wire.NewSet(provideHTTPClient)

	repoSet = wire.NewSet(
		clientSet,
		provideCepRepository,
	)

	repoWeatherSet = wire.NewSet(
		configSet,
		clientSet,
		provideWeatherAPIKey,
		provideWeatherRepository,
	)

	getCityUseCaseSet = wire.NewSet(
		repoSet,
		provideGetCityUseCase,
	)

	getWeatherUseCaseSet = wire.NewSet(
		repoWeatherSet,
		provideGetWeatherByCityUseCase,
	)
)

// --- Injectors ---

func InitializeGetCityByCepUseCase() (*usecase.GetCityByCepUseCase, error) {
	wire.Build(getCityUseCaseSet)
	return &usecase.GetCityByCepUseCase{}, nil
}

func InitializeGetWeatherByCityUseCase() (*weatherUsecase.GetWeatherByCityUseCase, error) {
	wire.Build(getWeatherUseCaseSet)
	return &weatherUsecase.GetWeatherByCityUseCase{}, nil
}
