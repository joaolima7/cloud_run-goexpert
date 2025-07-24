package cep

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/joaolima7/cloud_run-goexpert/internal/domain/cep"
	httpclient "github.com/joaolima7/cloud_run-goexpert/internal/infra/http_client"
)

type GetCityByCepRepositoryImpl struct {
	client httpclient.HTTPClient
}

func NewGetCityByCepRepositoryImpl(client httpclient.HTTPClient) *GetCityByCepRepositoryImpl {
	return &GetCityByCepRepositoryImpl{
		client: client,
	}
}

func (r *GetCityByCepRepositoryImpl) GetCityByCep(cepInput *cep.Cep) (string, error) {
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cepInput.CEP)
	ctx := context.Background()
	data, err := r.client.Get(ctx, url)
	if err != nil {
		println("deu erro no cep")
		return "", err
	}

	var cityData struct {
		City string `json:"localidade"`
	}

	if err := json.Unmarshal(data, &cityData); err != nil {
		return "", fmt.Errorf("error unmarshalling response: %w", err)
	}

	if cityData.City == "" {
		return "", cep.ErrCepNotFound
	}

	return cityData.City, nil
}
