package config

import "os"

type DBConfig struct {
	Host           string
	Port           string
	User           string
	Password       string
	DBName         string
	MigrationsPath string
}

type HttpConfig struct {
	Url  string
	Port string
}

func LoadDBConfig() DBConfig {
	return DBConfig{
		Host:           os.Getenv("DB_HOST"),
		Port:           os.Getenv("DB_PORT"),
		User:           os.Getenv("DB_USER"),
		Password:       os.Getenv("DB_PASSWORD"),
		DBName:         os.Getenv("DB_NAME"),
		MigrationsPath: os.Getenv("DB_MIGRATIONS_PATH"),
	}
}

func LoadHttpConfig() HttpConfig {
	return HttpConfig{
		Url:  os.Getenv("HTTP_URL"),
		Port: os.Getenv("HTTP_PORT"),
	}
}
