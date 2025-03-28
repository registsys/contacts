package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	PostgresDSN string `yaml:"postgres_dsn"`
}

func New(file string) Config {
	cfg := Config{}

	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Printf("config file %s not found, using default config\n", file)
		return cfg
	}

	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		fmt.Printf("config file %s can't be parsed, using default config\n", file)
		return cfg
	}

	return cfg
}
