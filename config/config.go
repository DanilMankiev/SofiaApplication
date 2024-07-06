package config

import (
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type (
	Config struct{
		Http `yaml:"http"`
		Postgres `yaml:"postgres"`
		StaticHost `yaml:"statichost"`
	}

	Http struct{
		PortHTTP string  `yaml:"port"`
	}
	Postgres struct{
		User string `yaml:"user"`
  		Host string `yaml:"host"`
  		Port string `yaml:"port"`
  		Dbname string `yaml:"dbname"`
  		Sslmode string `yaml:"sslmode"`
		Password string
	}
	StaticHost struct{
		StaticHost string `yaml:"statichost"`
	}
)

func InitConfig() (*Config,error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("../config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = godotenv.Load("../.env")
	if err != nil {
		return nil, err
	}
	cfg.Postgres.Password=os.Getenv("PG_PASSWORD")

	return cfg, nil
}