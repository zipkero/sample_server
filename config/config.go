package config

import (
	"fmt"
	"github.com/go-yaml/yaml"
	"os"
)

type Config struct {
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
}

func NewConfig(filePath string) *Config {
	c := &Config{}
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	if err := yaml.NewDecoder(file).Decode(c); err != nil {
		panic(err)
	}

	fmt.Println(c)

	return c
}
