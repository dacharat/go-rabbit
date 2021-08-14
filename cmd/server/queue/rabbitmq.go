package queue

import (
	"log"

	"github.com/dacharat/go-rabbit/config"
	"github.com/streadway/amqp"
)

var (
	ch   *amqp.Channel
	conn *amqp.Connection
)

func Init() {
	conn, err := amqp.Dial(config.Rabbit.Host)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to rabbitmq!! 🐇")

	ch, err = conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	err = Publish("hello", []byte("start server"))
	if err != nil {
		log.Fatal(err)
	}
}

func Close() {
	log.Println("Closed to rabbitmq!! 🐇")
	ch.Close()
	conn.Close()
}
