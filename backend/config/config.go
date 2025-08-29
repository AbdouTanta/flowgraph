package config

import (
	"log"
	"os"
)

type AppConfig struct {
	MongodbUri string
	SigningKey string
}

var Config AppConfig

func InitConfig() {
	if uri := os.Getenv("MONGODB_URI"); uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environment variable. See\n\t https://docs.mongodb.com/drivers/go/current/usage-examples/")
	} else {
		Config.MongodbUri = uri
	}

	if signingKey := os.Getenv("SIGNING_KEY"); signingKey == "" {
		log.Fatal("You must set your 'SIGNING_KEY' environment variable.")
	} else {
		Config.SigningKey = signingKey
	}
}
