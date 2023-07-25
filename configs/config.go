package configs

import (
	"log"

	"github.com/spf13/viper"
)

type ServerConfig struct {
	Env         string `mapstructure:"env"`
	ServiceName string `mapstructure:"service_name"`
	ServiceURL  string `mapstructure:"service_url"`
	Port        int    `mapstructure:"port"`
}

var Config *ServerConfig

func LoadConfigs(log *log.Logger) error {
	vp := viper.New()
	var config ServerConfig

	vp.SetConfigName("conf")
	vp.SetConfigType("yaml")
	vp.AddConfigPath(".")
	err := vp.ReadInConfig()
	if err != nil {
		log.Printf("error loading configs, err:%v", err)
		return err
	}

	err = vp.Unmarshal(&config)
	if err != nil {
		log.Printf("error unmarshalling configs, err:%v", err)
		return err
	}

	Config = &config

	return nil
}
