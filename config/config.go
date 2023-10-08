package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

// OAuth2

type oauth2_provider struct {
	Id          string `yaml:"client_id" env-description:"ClientId" env-required`
	Secret      string `yaml:"secret" env-description:"ClientId" env-required`
	RedirectURL string `yaml:"redirect_url" env-description:"ClientId" env-required`
}

// Summary Config

type config struct {
	// Ports
	Ports struct {
		Gateway string `yaml:"gateway" env-default:"5000"`
		Auth    string `yaml:"auth" env-default:"5001"`
	} `yaml:"ports"`

	// Databases
	Databases struct {
		Auth string `yaml:"auth" env-required`
	} `yaml:"databases"`

	// OAuth2
	Oauth2 struct {
		Google oauth2_provider `yaml:"google"`
	} `yaml:"oauth2"`
}

// should end with "/"
func New(path_to_root string) config {
	var cfg config
	cleanenv.ReadConfig(createPath(path_to_root), &cfg)
	return cfg
}

func createPath(path_to_root string) string {
	return path_to_root + "config/config.yaml"
}
