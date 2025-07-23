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
	apiKey string
}

func NewGetCityByCepRepositoryImpl(client httpclient.HTTPClient, apiKey string) *GetCityByCepRepositoryImpl {
	return &GetCityByCepRepositoryImpl{
		client: client,
		apiKey: apiKey,
	}
}

func (r *GetCityByCepRepositoryImpl) GetCityByCep(cep *cep.Cep) (string, error) {
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep.CEP)
	ctx := context.Background()
	data, err := r.client.Get(ctx, url)
	if err != nil {
		return "", err
	}

	var cityData struct {
		City string `json:"localidade"`
	}

	if err := json.Unmarshal(data, &cityData); err != nil {
		return "", fmt.Errorf("error unmarshalling response: %w", err)
	}

	if cityData.City == "" {
		return "", fmt.Errorf("city not found for CEP: %s", cep.CEP)
	}

	return cityData.City, nil
}
