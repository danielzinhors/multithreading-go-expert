package viacep

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/danielzinhors/multithreading-go-expert/configs"
	"github.com/danielzinhors/multithreading-go-expert/internal/entity"
)

func ConsultaCep(cep string) (*entity.CepVia, error) {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	endPointVia := configs.WebServerViaCep
	req, error := http.Get(endPointVia + cep + "/json/")
	if error != nil {
		return nil, error
	}
	defer req.Body.Close()
	res, error := io.ReadAll(req.Body)
	if error != nil {
		return nil, error
	}
	cepVia := entity.NewCepVia()
	error = json.Unmarshal(res, &cepVia)
	if error != nil {
		return nil, error
	}

	return cepVia, nil
}
