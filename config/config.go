package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Database config
type Database struct {
	Host string
	Name string
}

type RabbitMq struct {
	Host string
	Type string
}

var (
	DB      Database
	Rabbit  RabbitMq
	GinMode = "debug"
)

func init() {
	CurrentEnvironment, ok := os.LookupEnv("GO_ENV")

	var err error

	if !ok {
		CurrentEnvironment = "development"
		err = godotenv.Load()
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("Environment: " + CurrentEnvironment)

	DB = setDatabaseConfig()
	Rabbit = setRabbtiConfig()
}

func setDatabaseConfig() Database {
	db := Database{
		Host: os.Getenv("MONGODB_HOST"),
		Name: os.Getenv("MONGODB_DATABASE"),
	}

	return db
}

func setRabbtiConfig() RabbitMq {
	rabbit := RabbitMq{
		Host: os.Getenv("RABBITMQ_HOST"),
		Type: os.Getenv("RABBITMQ_TYPE"),
	}

	return rabbit
}
