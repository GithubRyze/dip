package conf

import (
	"github.com/spf13/viper"
)

type AppConfig struct {
	WebPort         string `yaml:"webPort"`
	EndpointTlsPort string `yaml:"endpointTlsPort"`
	EndpointPort    string `yaml:"endpointPort"`
	LoggerPath      string `yaml:"loggerPath"`
	ApplicationName string `yaml:"applicationName"`
	LogLevel        string `yaml:"logLevel"`
}

var (
	Config *AppConfig
)

func LoadAppConfig(configName string) error {
	v := viper.New()
	v.SetConfigName(configName)
	v.SetConfigType("yaml")
	v.AddConfigPath("./config")
	err := v.ReadInConfig()
	if err != nil {
		return err
	}
	Config = &AppConfig{}
	if err = v.Unmarshal(Config); err != nil {
		return err
	}
	return nil
}
