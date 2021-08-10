package config

import (
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		Srv Server   `yaml:"server"`
		Pg  Postgres `yaml:"postgres"`
		Log Logrus   `yaml:"logrus"`
		JWT JWT      `yaml:"jwt"`
	}

	Server struct {
		BindAddr string `yaml:"bind_addr"`
	}

	Postgres struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DBName   string `yaml:"dbname"`
		SSLMode  string `yaml:"sslmode"`
	}

	Logrus struct {
		LogLevel string `yaml:"log_level"`
	}

	JWT struct {
		Expire time.Duration `yaml:"expire"`
		Key    string        `yaml:"key"`
	}
)

func NewConfig(filePath string) (*Config, error) {
	var cfg Config
	if err := cleanenv.ReadConfig(filePath, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
