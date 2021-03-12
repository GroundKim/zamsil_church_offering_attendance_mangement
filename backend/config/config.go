package config

import (
	"fmt"
	"zamsil_church_offering_attendance_mangement/backend/models"

	"github.com/spf13/viper"
)

var conf *models.Config

// Read the config file from the current directory and marshal
// into the conf config struct.
func GetConf(env string) *models.Config {
	filename := fmt.Sprintf("./%s_config", env)
	viper.AddConfigPath("./config")
	viper.SetConfigName(filename)
	err := viper.ReadInConfig()

	if err != nil {
		fmt.Printf("%v", err)
	}

	conf := &models.Config{}
	err = viper.Unmarshal(conf)
	if err != nil {
		fmt.Printf("unable to decode into config struct, %v", err)
	}
	return conf
}
