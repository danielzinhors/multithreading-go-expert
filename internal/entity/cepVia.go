package entity

type CepVia struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func NewCepVia(cep, logradouro, complemento, bairro, localidade, uf, ibge, gia, ddd, siafi string) (*CepVia, error) {
	return &CepVia{
		Cep:         cep,
		Logradouro:  logradouro,
		Complemento: complemento,
		Bairro:      bairro,
		Localidade:  localidade,
		Uf:          uf,
		Ibge:        ibge,
		Gia:         gia,
		Ddd:         ddd,
		Siafi:       siafi,
	}, nil
}
