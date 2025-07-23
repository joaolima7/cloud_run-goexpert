package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joaolima7/cloud_run-goexpert/internal/domain/cep"
	"github.com/joaolima7/cloud_run-goexpert/internal/domain/cep/usecase"
	weatherUC "github.com/joaolima7/cloud_run-goexpert/internal/domain/weather/usecase"
)

type AppHandler struct {
	useCaseCep     *usecase.GetCityByCepUseCase
	useCaseWeather *weatherUC.GetWeatherByCityUseCase
}

func NewAppHandler(uc *usecase.GetCityByCepUseCase, weatherUC *weatherUC.GetWeatherByCityUseCase) *AppHandler {
	return &AppHandler{
		useCaseCep:     uc,
		useCaseWeather: weatherUC,
	}
}

func (h *AppHandler) RegisterRoutes(r chi.Router) {
	r.Route("/cep", func(r chi.Router) {
		r.Get("/{cep}", h.getWeatherByCity)
	})
}

func (h *AppHandler) getWeatherByCity(w http.ResponseWriter, r *http.Request) {
	cepParam := chi.URLParam(r, "cep")
	out, err := h.useCaseCep.Execute(usecase.CepInputDTO{Cep: cepParam})
	if err != nil {
		if err == cep.ErrInvalidCep {
			h.respondError(w, http.StatusUnprocessableEntity, err)
			return
		} else if err == cep.ErrCepNotFound {
			h.respondError(w, http.StatusNotFound, err)
			return
		} else {
			h.respondError(w, http.StatusInternalServerError, err)
			return
		}
	}

	outWeather, err := h.useCaseWeather.Execute(weatherUC.WeatherInputDTO{City: out.City})
	if err != nil {
		h.respondError(w, http.StatusInternalServerError, err)
		return
	}

	h.respondJSON(w, http.StatusOK, outWeather)
}

func (h *AppHandler) respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func (h *AppHandler) respondError(w http.ResponseWriter, status int, err error) {
	h.respondJSON(w, status, map[string]string{"error": err.Error()})
}
