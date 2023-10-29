package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/danielzinhors/multithreading-go-expert/infra/viacep"
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
	cepPesquisa := r.URL.Query().Get("cep")
	if cepPesquisa == "" {
		w.WriteHeader(http.StatusInternalServerError)
		error := Error{Message: "O CEP é obrigatório"}
		json.NewEncoder(w).Encode(error)
		return
	}
	cep, err := viacep.ConsultaCep(cepPesquisa)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	json.NewEncoder(w).Encode(cep)
	//cepVia := make(chan entity.CepVia)
	// go func() {
	// 	var cep *entity.CepVia
	// 	cep, err := viacep.ConsultaCep(cepPesquisa)
	// 	if err != nil {
	// 		w.WriteHeader(http.StatusInternalServerError)
	// 		error := Error{Message: err.Error()}
	// 		json.NewEncoder(w).Encode(error)
	// 		return
	// 	}
	// 	cepVia <- cep
	// }()

	// select {
	// case cep := <-cepVia:
	// 	w.WriteHeader(http.StatusOK)
	// 	json.NewEncoder(w).Encode(cep)
	// case <-time.After(time.Second * 3):
	// 	json.NewEncoder(w).Encode([]byte("Timeout"))
	// }
}
