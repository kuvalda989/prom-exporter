package config

import (
	"log"
	"os"
)

type Config struct {
	Port             string `env:"PROM_PORT"`
	Source_type      string `env:"PROM_SOURCE_TYPE"`
	Source           string `env:"PROM_SOURCE"`
	RenewTimeSeconds string `env:"PROM_RENEW"`
	Token            string `env:"PROM_TOKEN"`
}

func checkEnvs(env string) string {
	// проверка что переменная окружения существует
	if os.Getenv(env) == "" {
		log.Fatal("Env ", env, " is not defined")
	}
	// тут нужны проверки переменных на валидность и неплохо бы сразу в нужный тип перевести
	return os.Getenv(env)
}

func Get() Config {

	cfg := Config{
		Port:             checkEnvs("PROM_PORT"),
		Source_type:      checkEnvs("PROM_SOURCE_TYPE"),
		Source:           checkEnvs("PROM_SOURCE"),
		RenewTimeSeconds: checkEnvs("PROM_RENEW"),
		Token:            checkEnvs("PROM_TOKEN"),
	}
	return cfg
}
