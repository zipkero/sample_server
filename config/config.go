package config

import (
	"github.com/go-yaml/yaml"
	"os"
)

type Config struct {
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"sample_server"`
}

func NewConfig(filePath string) *Config {
	c := &Config{}

	if file, err := os.Open(filePath); err != nil {
		panic(err)
	} else if err = yaml.NewDecoder(file).Decode(c); err != nil {
		panic(err)
	}

	return c
}
