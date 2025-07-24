package repository

import "github.com/joaolima7/cloud_run-goexpert/internal/domain/cep"

type GetCityByCepRepository interface {
	GetCityByCep(cepInput *cep.Cep) (string, error)
}
