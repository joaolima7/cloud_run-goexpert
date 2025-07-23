package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joaolima7/cloud_run-goexpert/internal/domain/cep/usecase"
)

type CepHandler struct {
	useCase *usecase.GetCityByCepUseCase
}

func NewCepHandler(uc *usecase.GetCityByCepUseCase) *CepHandler {
	return &CepHandler{useCase: uc}
}

func (h *CepHandler) RegisterRoutes(r chi.Router) {
	r.Route("/cep", func(r chi.Router) {
		r.Get("/{cep}", h.getCityByCep)
	})
}

func (h *CepHandler) getCityByCep(w http.ResponseWriter, r *http.Request) {
	cepParam := chi.URLParam(r, "cep")
	out, err := h.useCase.Execute(usecase.CepInputDTO{Cep: cepParam})
	if err != nil {
		h.respondError(w, http.StatusBadRequest, err)
		return
	}
	h.respondJSON(w, http.StatusOK, out)
}

func (h *CepHandler) respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func (h *CepHandler) respondError(w http.ResponseWriter, status int, err error) {
	h.respondJSON(w, status, map[string]string{"error": err.Error()})
}
