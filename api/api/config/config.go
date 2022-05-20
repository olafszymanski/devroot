package config

import (
	"os"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		Port string `yaml:"port" env:"SERVER_PORT"`
	} `yaml:"server"`
	Database struct {
		Host     string `yaml:"host" env:"DB_HOST"`
		Port     string `yaml:"port" env:"DB_PORT"`
		Driver   string `yaml:"driver" env:"DB_DRIVER"`
		Username string `yaml:"username" env:"DB_USERNAME"`
		Password string `yaml:"password" env:"DB_PASSWORD"`
		Name     string `yaml:"name" env:"DB_NAME"`
	} `yaml:"database"`
}

func NewConfig() *Config {
	var cfg Config

	f, err := os.Open("config.yml")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// YAML
	decoder := yaml.NewDecoder(f)
	if err := decoder.Decode(&cfg); err != nil {
		panic(err)
	}

	// ENV
	if err := envconfig.Process("", &cfg); err != nil {
		panic(err)
	}
	return &cfg
}
