# cloud_run-goexpert

Projeto em Go para consulta de cidade por CEP e previsão do tempo, pronto para deploy no Google Cloud Run.

## Funcionalidades

- Consulta cidade a partir de um CEP (ViaCEP)
- Consulta previsão do tempo para a cidade (OpenWeatherMap)
- API HTTP simples

## Variáveis de Ambiente

Crie um arquivo `.env` na raiz com:

```env
WEB_SERVER_PORT=8080
WEATHER_API_KEY=sua_chave_openweathermap
```

## Como rodar localmente

```bash
go run cmd/main.go
```

## Endpoints

- `GET /cep/{cep}`: Retorna a previsão do tempo para a cidade do CEP informado.

Acesse: `http://localhost:8080/cep/{cep}`

Exemplo:
```
GET http://localhost:8080/cep/15041785
```

## Resposta

```json
{
  "temp_C": 25.0,
  "temp_F": 77.0,
  "temp_K": 298.15
}
```

## ENDPOINT DE PRODUÇÃO

```
https://cloud-run-goexpert-aae7wkoyrq-uc.a.run.app/cep/15041785
```