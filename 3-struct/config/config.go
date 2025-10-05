package config

import "os"

type Config struct {
	MasterKey string
}

func NewConfig() *Config {
	masterKey := os.Getenv("MASTER_KEY")
	if masterKey == "" {
		panic("env MASTER_KEY not provided")
	}
	return &Config{
		MasterKey: masterKey,
	}
}
