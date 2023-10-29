package viacep

import (
	"net/http"

	"github.com/danielzinhors/multithreading-go-expert/configs"
)

func consultaCep(cep string) (string, error) {
	req, err := http.Get(configs.WebServerViaCep + cep + "/json/")
	if err != nil {
		return "", err
	}
	defer req.Body.Close()
	// res, err := io.ReadAll(req.Body)
	// if err != nil {
	// 	return "", err
	// }
	// var endereco Endereco
	// err = json.Unmarshal(res, &endereco)
	// if err != nil {
	// 	return "", err
	// }
	return "", nil
}
