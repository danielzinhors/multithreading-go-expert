package entity

type CepAPi struct {
	Code       string `json:"code"`
	State      string `json:"state"`
	City       string `json:"city"`
	District   string `json:"district"`
	Address    string `json:"address"`
	Status     int    `json:"status"`
	Ok         bool   `json:"ok"`
	StatusText string `json:"statusText"`
}

func NewCepApi(code, state, complemento, city, district, address, statusText string, status int, ok bool) (*CepAPi, error) {
	return &CepAPi{
		Code:       code,
		State:      state,
		City:       city,
		District:   district,
		Address:    address,
		Status:     status,
		Ok:         ok,
		StatusText: statusText,
	}, nil
}
