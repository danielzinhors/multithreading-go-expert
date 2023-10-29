package configs

import "github.com/spf13/viper"

var cfg *conf

type conf struct {
	webServerViaCep  string `mapstructure:"WEB_SERVER_VIA_CEP"`
	webserviceApiCep string `mapstructure:"WEB_SERVER_API_CEP"`
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
	return cfg, err
}
