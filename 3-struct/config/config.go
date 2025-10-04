package config

import "os"

type Config struct {
	Key string
}

func NewConfig() *Config {
	key := os.Getenv("KEY")
	if key == "" {
		panic("env KEY not provided")
	}
	return &Config{
		Key: key,
	}
}
