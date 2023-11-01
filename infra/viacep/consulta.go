package viacep

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/danielzinhors/multithreading-go-expert/configs"
	"github.com/danielzinhors/multithreading-go-expert/internal/entity"
)

func ConsultaCep(cep string) *entity.CepVia {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	endPointVia := configs.WebServerViaCep
	cep = strings.Replace(string(cep), `-`, "", -1)
	req, error := http.Get(endPointVia + cep + "/json/")
	if error != nil {
		return nil
	}
	defer req.Body.Close()
	res, error := io.ReadAll(req.Body)
	if error != nil {
		return nil
	}

	var cepVia = *entity.NewCepVia()
	error = json.Unmarshal(res, &cepVia)
	if error != nil {
		return nil
	}

	return &cepVia
}
