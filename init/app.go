package init

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
	"sync"
)

type AppConfig struct {
	WebPort         string `yaml:"webPort"`
	EndpointTlsPort string `yaml:"endpointTlsPort"`
	EndpointPort    string `yaml:"endpointPort"`
}

var (
	Config *AppConfig
	once   sync.Once // 确保只初始化一次
)

func InitAppConfig(configName string) {
	once.Do(func() {
		//logger.Printf("loading application configuration %s \n", configName)
		var err error
		var AppPath string
		if AppPath, err = filepath.Abs(filepath.Dir(os.Args[0])); err != nil {
			panic(err)
		}
		log.Printf("AppPath dir : %s \n", AppPath)
		WorkPath, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		appConfigPath := filepath.Join(WorkPath, "/configs")
		log.Printf("config path : %s \n", appConfigPath)
		v := viper.New()
		v.SetConfigName(configName)
		v.SetConfigType("yaml")
		v.AddConfigPath(appConfigPath)
		err = v.ReadInConfig()
		if err != nil {
			log.Printf("file not found : %s \n", err)
			panic(err)
		}
		Config = &AppConfig{}
		if err := v.Unmarshal(Config); err != nil {
			log.Printf("file read error : %s \n", err)
			panic(err)
		}
		log.Printf("bootstrap config loaded success %+v \n", Config)
	})
}
