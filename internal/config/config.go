package config

import (
	"fmt"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Database    DatabaseConfig `env-prefix:"DB_"`
	Weather     WeatherConfig  `env-prefix:"WEATHER_"`
	Log         LogConfig      `env-prefix:"LOG_"`
	Telegram    TelegramConfig `env-prefix:"TELEGRAM_"`
	HTTP        HTTPConfig     `env-prefix:"HTTP_"`
	BotToken    string         `env:"BOT_TOKEN" env-required:"true"`
	Environment string         `env:"ENVIRONMENT" env-default:"development"`
}

type DatabaseConfig struct {
	Host     string `env:"HOST" env-required:"true"`
	Port     int    `env:"PORT" env-required:"true"`
	User     string `env:"USER" env-required:"true"`
	Password string `env:"PASSWORD" env-required:"true"`
	Name     string `env:"NAME" env-required:"true"`
}

type WeatherConfig struct {
	APIKey       string        `env:"API_KEY" env-required:"true"`
	BaseURL      string        `env:"BASE_URL" env-default:"http://api.weatherapi.com/v1"`
	RequestDelay time.Duration `env:"REQUEST_DELAY" env-default:"1s"`
}

type LogConfig struct {
	Level  string `env:"LEVEL" env-default:"info"`
	Format string `env:"FORMAT" env-default:"json"`
}

type TelegramConfig struct {
	Timeout int `env:"TIMEOUT" env-default:"60"`
}

type HTTPConfig struct {
	ClientTimeout time.Duration `env:"CLIENT_TIMEOUT" env-default:"10s"`
}

func LoadConfig() (*Config, error) {
	cfg := &Config{}
	err := cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to read config: %w", err)
	}
	return cfg, nil
}
