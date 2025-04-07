package pkg

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type (
	Config struct {
		Name     string `env-required:"false" yaml:"name" env:"NAME"`
		BaseUrl  string `yaml:"base_url" env:"BASE_URL"`
		AppId    string `yaml:"app_id" env:"APP_ID"`
		BotToken string `yaml:"bot_token" env:"BOT_TOKEN"`
		ChatId   string `yaml:"chat_id" env:"CHAT_ID"`
	}
	// Redis -.
	// Redis struct {
	// 	RedisHost string `env-required:"false" yaml:"host" env:"REDIS_HOST"`
	// 	RedisPort int    `env-required:"false" yaml:"port" env:"REDIS_PORT"`
	// 	RedisUser string `env-required:"false" yaml:"user" env:"REDIS_USER"`
	// 	RedisPass string `env-required:"false" yaml:"pass" env:"REDIS_PASS"`
	// 	Enabled   bool   `env-required:"false" yaml:"enabled" env:"REDIS_ENABLED"`
	// }
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("func.yaml", cfg)
	if err != nil {
		return cfg, fmt.Errorf("config error: %w", err)
	}

	err = godotenv.Load(".env")
	if err != nil {
		return cfg, fmt.Errorf("error loading .env file: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}
