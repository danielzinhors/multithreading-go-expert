package apicep

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/danielzinhors/multithreading-go-expert/configs"
	"github.com/danielzinhors/multithreading-go-expert/internal/entity"
)

func ConsultaCep(cep string) *entity.CepAPi {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	endPointVia := configs.WebserviceApiCep
	println(cep)
	println(endPointVia)
	req, error := http.Get(endPointVia + cep + ".json")
	if error != nil {
		return nil
	}
	defer req.Body.Close()
	res, error := io.ReadAll(req.Body)
	println(string(res))
	if error != nil {
		return nil
	}

	var apiCep = *entity.NewCepApi()
	error = json.Unmarshal(res, &apiCep)
	if error != nil {
		return nil
	}

	return &apiCep
}
