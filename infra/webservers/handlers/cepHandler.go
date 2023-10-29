package handlers

import (
	"net/http"

	"github.com/danielzinhors/multithreading-go-expert/internal/entity"
)

type Error struct {
	Message string `json:"message"`
}

type CepHandler struct {
	CepApi entity.CepAPi
	CepVia entity.CepVia
}

func NewCepHanbler() *CepHandler {
	return &CepHandler{}
}

// Busca CEP godoc
// @Summary Busca um Cep
// @Description Busca um Cep em 2 api e devolve a mais rapida
// @Tags cep
// @Param  cep path string true "cep"
// @Success  200 {object} entity.Product
// @Failure  404  {object}  Error
// @Failure  500  {object}  Error
// @Router  /cep [get]
func (c *CepHandler) BuscaCep(w http.ResponseWriter, r *http.Request) {
	cep := r.URL.Query().Get("cep")
	println(cep)
}
