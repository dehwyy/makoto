package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type ServerPort int

const (
	PortAuth ServerPort = iota + 5001
	PortHashmap
)

type Config struct {
	DatabaseDsn string `required:"true"    envconfig:"DATABASE_DSN"`
	JwtSecret   string `required:"true"    envconfig:"JWT_SECRET"`

	// OAuth2
	// google
	GoogleClientId     string `required:"true"    envconfig:"GOOGLE_CLIENT_ID"`
	GoogleClientSecret string `required:"true"    envconfig:"GOOGLE_CLIENT_SECRET"`
	GoogleRedirectURL  string `required:"true"    envconfig:"GOOGLE_REDIRECT_URL"`
	// github
	GithubClientId     string `required:"true"    envconfig:"GITHUB_CLIENT_ID"`
	GithubClientSecret string `required:"true"    envconfig:"GITHUB_CLIENT_SECRET"`
	GithubRedirectURL  string `required:"true"    envconfig:"GITHUB_REDIRECT_URL"`
	// discord
	DiscordClientId     string `required:"true"    envconfig:"DISCORD_CLIENT_ID"`
	DiscordClientSecret string `required:"true"    envconfig:"DISCORD_CLIENT_SECRET"`
	DiscordRedirectURL  string `required:"true"    envconfig:"DISCORD_REDIRECT_URL"`

	// Mode
	NodeEnv string `required:"false" envconfig:"NODE_ENV" default:"development" `

	// Urls
	TwirpGatewayUrl string `required:"true"    envconfig:"TWIRP_GATEWAY_URL"`
	HttpGatewayUrl  string `required:"true"    envconfig:"HTTP_GATEWAY_URL"`
	AuthUrl         string `required:"true"    envconfig:"AUTH_URL"`
	HashmapUrl      string `required:"true"    envconfig:"HASHMAP_URL"`

	// Gmail
	GmailSennderName    string `required:"true"    envconfig:"GMAIL_SENDER_NAME"`
	GmailSenderAddr     string `required:"true"    envconfig:"GMAIL_SENDER_ADDRESS"`
	GmailSenderPassword string `required:"true"    envconfig:"GMAIL_SENDER_PASSWORD"`
}

func New() *Config {
	var cfg Config

	wd, err := os.Getwd()
	if err != nil {
		fmt.Printf("failed to get current working directory: %v\n", err)
	}

	wd = filepath.Join(wd, "..", "..", "..")

	envPath := filepath.Join(wd, ".env")
	_ = godotenv.Load(envPath)

	envPath = filepath.Join(wd, "..", ".env")
	_ = godotenv.Load(envPath)

	if err = envconfig.Process("", &cfg); err != nil {
		fmt.Printf("failed to process env: %v\n", err)
	}

	return &cfg
}
