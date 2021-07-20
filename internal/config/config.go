package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		Srv Server   `yaml:"server"`
		Pg  Postgres `yaml:"postgres"`
		Log Logrus   `yaml:"logrus"`
	}

	Server struct {
		BindAddr string `yaml:"bind_addr" env:"BINDADDR"`
	}

	Postgres struct {
		Host     string `yaml:"host" env:"HOST"`
		Port     int    `yaml:"port" env:"PORT"`
		User     string `yaml:"user" env:"USER"`
		Password string `yaml:"password" env:"PASSWORD"`
		DBName   string `yaml:"dbname" env:"DBNAME"`
		SSLMode  string `yaml:"sslmode" env:"SSLMODE"`
	}

	Logrus struct {
		LogLevel string `yaml:"log_level" env:"LOGLEVEL"`
	}
)

func NewConfig(filePath string) (*Config, error) {
	var cfg Config
	if err := cleanenv.ReadConfig(filePath, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
