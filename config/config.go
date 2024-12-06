package config

import (
	"gin-gorm/src/utils/logger"

	"github.com/caarlos0/env/v11"
)

type AppConfig struct {
	Port string `env:"APP_PORT" envDefault:":3000"`
	Mode string `env:"APP_MODE" envDefault:"prod"`
}

type DBConfig struct {
	Driver   string `env:"DB_DRIVER"`
	Host     string `env:"DB_HOST"`
	User     string `env:"DB_USER"`
	Port     string `env:"DB_PORT"`
	Name     string `env:"DB_NAME"`
	Password string `env:"DB_PASSWORD"`
}

type Config struct {
	App AppConfig
	DB DBConfig
}

func GetConfig() Config {
	
	cfg := Config{}
	opts := env.Options{RequiredIfNoDef: true}

	if err := env.ParseWithOptions(&cfg, opts); err != nil {
		logger.Fatal("Error during parse .env", "err", err.Error())
	}

	return cfg
}
