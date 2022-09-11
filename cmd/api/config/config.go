package config

import (
	"log"

	"github.com/spf13/viper"
)

func LoadConfig() Config {
	var config Config

	vi := viper.New()

	vi.SetConfigType("yml")
	vi.SetConfigName("config")
	vi.AddConfigPath(".")

	err := vi.ReadInConfig()
	if err != nil {
		log.Fatalln("ha ocurrido un error al leer la configuración: ", err)
	}

	err = vi.Unmarshal(&config)
	if err != nil {
		log.Fatalln("ha ocurrido un error al crear estructura de configuración: ", err)
	}

	return config
}

type Config struct {
	Server     Server     `mapstruct:"server"`
	Datasource Datasource `mapstruct:"datasource"`
}

type Server struct {
	Port string `mapstruct:"port"`
}

type Datasource struct {
	Name     string `mapstruct:"name"`
	Host     string `mapstruct:"host"`
	Port     string `mapstruct:"port"`
	User     string `mapstruct:"user"`
	Password string `mapstruct:"password"`
	Database string `mapstruct:"database"`
}
