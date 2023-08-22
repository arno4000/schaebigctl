package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func InitConfig() {
	viper.SetConfigName("schaebigctl")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME")
	viper.AddConfigPath("/etc")
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Fatalf("Failed to read in config %s", err)
	}
}
