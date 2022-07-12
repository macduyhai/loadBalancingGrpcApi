package config

import (
	"log"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type Config struct {
	PortEngine string `env:"PORT_ENGINE"default:"9090"`
	MYSQLURL   string `env:"MYSQL_URL,required"`
	APIKEY     string `env:"API_KEY,required"`
	SecretKey  string `env:"SECRET_KEY,required"`
	PublicKey  string `env:"PUBLIC_KEY,required"`
	PrivateKey string `env:"PRIVATE_KEY,required"`
}

func NewConfig(files ...string) *Config {
	cfg := Config{}
	err := godotenv.Load(files...)
	if err != nil {
		log.Println("Env File could not found ")
		panic(err.Error())

	}
	err = env.Parse(&cfg)

	if err != nil {
		log.Println("Parse config file error:" + err.Error())
		return nil
	}
	return &cfg

}
