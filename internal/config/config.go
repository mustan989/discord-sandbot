package config

import (
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

func ReadFile(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return Read(file)
}

func Read(reader io.Reader) (*Config, error) {
	cfg := new(Config)
	if err := yaml.NewDecoder(reader).Decode(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

type Config struct {
	HTTP HTTP `yaml:"http"`
}

type HTTP struct {
	Port int `yaml:"port"`
}
