package configs

import (
	"errors"

	"github.com/spf13/viper"
)

var cfg *conf

type conf struct {
	WebServerViaCep  string `mapstructure:"WEB_SERVER_VIA_CEP"`
	WebserviceApiCep string `mapstructure:"WEB_SERVER_API_CEP"`
	WebServerPort    string `mapstructure:"SERVER_PORT"`
}

func LoadConfig(path string) (*conf, error) {
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	if cfg.WebServerViaCep == "" {
		errors.New("Endpoint Via CEP não encontrado na config")
		return nil, err
	}

	if cfg.WebserviceApiCep == "" {
		errors.New("Porta do server não encontrado na config")
		return nil, err
	}
	return cfg, err
}
