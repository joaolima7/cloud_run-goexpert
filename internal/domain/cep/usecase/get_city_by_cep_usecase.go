package usecase

import (
	"github.com/joaolima7/cloud_run-goexpert/internal/domain/cep"
	"github.com/joaolima7/cloud_run-goexpert/internal/domain/cep/repository"
)

type CepInputDTO struct {
	Cep string `json:"cep"`
}

type CepOutputDTO struct {
	City string `json:"city"`
}

type GetCityByCepUseCase struct {
	GetCityByCepRepository repository.GetCityByCepRepository
}

func NewGetCityByCepUseCase(getCityByCepRepository repository.GetCityByCepRepository) *GetCityByCepUseCase {
	return &GetCityByCepUseCase{GetCityByCepRepository: getCityByCepRepository}
}

func (uc *GetCityByCepUseCase) Execute(input CepInputDTO) (*CepOutputDTO, error) {
	cep, err := cep.NewCep(input.Cep)
	if err != nil {
		return nil, err
	}

	city, err := uc.GetCityByCepRepository.GetCityByCep(cep)
	if err != nil {
		return nil, err
	}

	return &CepOutputDTO{City: city}, nil
}
