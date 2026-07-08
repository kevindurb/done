package config

import "os"

type Config struct {
	ListenAddr string
	SecretKey  string
	DBPath     string
}

func FromEnv() Config {
	return Config{
		ListenAddr: os.Getenv("LISTEN_ADDR"),
		SecretKey:  os.Getenv("SECRET_KEY"),
		DBPath:     os.Getenv("DB_PATH"),
	}
}
