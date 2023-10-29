package main

import (
	"github.com/danielzinhors/multithreading-go-expert/configs"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	EndPointViaCep := configs.WebServerViaCep
	if EndPointViaCep == "" {
		panic("Endpoint Via CEP não encontrado na config")
	}
	EndPointApiCep := configs.WebserviceApiCep
	if EndPointApiCep == "" {
		panic("Endpoint Api CEP não encontrado na config")
	}

}
