package config

import (
	"github.com/dehwyy/Makoto/backend/distributor/logger"
	"github.com/spf13/viper"
)

var (
	l = logger.New()
)

func init() {
	viper.SetConfigName("cfg")
	viper.AddConfigPath("../../_config")
	if err := viper.ReadInConfig(); err != nil {
		l.Fatalf("Error reading config file: %s", err)
	}
}

func GetOptionByKey(key string) (value string, isFound bool) {
	val := viper.GetString(key)

	return val, len(val) > 1
}
