package config

import (
	"github.com/joeshaw/envdecode"
	"log"
)

type Conf struct {
	AcessToken       string `env:"ACCESS_TOKEN,required"`
	AcessTokenSecret string `env:"ACCESS_TOKEN_SECRET,required"`
	ConsumerKey      string `env:"CONSUMER_KEY,required"`
	ConsumerSecret   string `env:"CONSUMER_SECRET,required"`
}

func AppConfig() *Conf {
	var c Conf
	if err := envdecode.StrictDecode(&c); err != nil {
		log.Println("Setup your environment variables ACCESS_TOKEN,ACCESS_TOKEN_SECRET,CONSUMER_KEY and CONSUMER_SECRET")
		log.Fatalf("Failed to decode: %s", err)
	}
	return &c
}
