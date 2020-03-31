package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

func SetupViper() {
	viper.AutomaticEnv()

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	log.Print("CURRENT-ENV: ")
	log.Print(viper.GetString("env"))
}
