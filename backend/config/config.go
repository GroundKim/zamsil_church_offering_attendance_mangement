package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	DATABASE struct {
		DB_USERNAME string `yaml:"DB_USERNAME"`
		DB_PASSWORD string `yaml:"DB_PASSWORD"`
		DB_PORT     string `yaml:"DB_PORT"`
		DB_HOST     string `yaml:"DB_HOST"`
		DB_NAME     string `yaml:"DB_NAME"`
	} `yaml:"database"`

	AUTH struct {
		SECRETKEY string `yaml:"SECRETKEY"`
	} `yaml:"AUTH"`
}

// Read the config file from the current directory and marshal
// into the conf config struct.
func GetConf(env string) *Config {
	filename := fmt.Sprintf("./%s_config", env)
	viper.AddConfigPath("./config")
	viper.SetConfigName(filename)
	err := viper.ReadInConfig()

	if err != nil {
		fmt.Printf("%v", err)
	}

	conf := &Config{}
	err = viper.Unmarshal(conf)
	if err != nil {
		fmt.Printf("unable to decode into config struct, %v", err)
	}
	return conf
}
