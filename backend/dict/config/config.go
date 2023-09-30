package config

import (
	"github.com/dehwyy/makoto/backend/dict/logger"
	"github.com/spf13/viper"
)

var (
	l = logger.New()
)

func init() {
	// read global config
	viper.SetConfigName("config")
	viper.AddConfigPath("../../_config")
	if err := viper.ReadInConfig(); err != nil {
		l.Fatalf("Error reading global config file: %s", err)
	}

	viper.SetConfigName("local.config")
	viper.AddConfigPath("./dict/config")
	viper.MergeInConfig()
}

func GetOptionByKeyWithFlag(key string) (value string, isFound bool) {
	val := viper.GetString(key)

	return val, len(val) > 1
}

func GetOptionByKey(key string) string {
	val := viper.GetString(key)

	return val
}
