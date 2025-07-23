//go:build wireinject
// +build wireinject

package app

import (
	"time"

	"github.com/google/wire"
	"github.com/joaolima7/cloud_run-goexpert/config"
	"github.com/joaolima7/cloud_run-goexpert/internal/domain/cep/repository"
	"github.com/joaolima7/cloud_run-goexpert/internal/domain/cep/usecase"
	cepRepo "github.com/joaolima7/cloud_run-goexpert/internal/infra/external/cep"
	httpclient "github.com/joaolima7/cloud_run-goexpert/internal/infra/http_client"
)

// --- Providers ---

func provideConfig() (*config.Config, error) {
	return config.LoadConfig()
}

func provideHTTPClient() httpclient.HTTPClient {
	return httpclient.NewHTTPClientImpl(10 * time.Second)
}

func provideCepRepository(client httpclient.HTTPClient) repository.GetCityByCepRepository {
	return cepRepo.NewGetCityByCepRepositoryImpl(client)
}

func provideGetCityUseCase(repo repository.GetCityByCepRepository) *usecase.GetCityByCepUseCase {
	return usecase.NewGetCityByCepUseCase(repo)
}

// --- Sets ---

var (
	configSet = wire.NewSet(provideConfig)
	clientSet = wire.NewSet(provideHTTPClient)

	repoSet = wire.NewSet(
		clientSet,
		provideCepRepository,
	)

	getCityUseCaseSet = wire.NewSet(
		repoSet,
		provideGetCityUseCase,
	)
)

// --- Injectors ---

func InitializeGetCityByCepUseCase() (*usecase.GetCityByCepUseCase, error) {
	wire.Build(getCityUseCaseSet)
	return &usecase.GetCityByCepUseCase{}, nil
}
