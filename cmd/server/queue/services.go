package queue

import "github.com/streadway/amqp"

func Publish(routeName string, payload []byte) error {
	q, err := ch.QueueDeclare(
		routeName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return err
	}

	return ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        payload,
		})
}
