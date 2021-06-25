package bus

import (
	"log"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func GetQueue() (*amqp.Connection, *amqp.Channel, *amqp.Queue) {
	// TODO: Get from configuration
	conn, err := amqp.Dial("amqp://rabbitmq:rabbitmq@rabbit1:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")

	// We create a Queue to send the message to.
	q, err := ch.QueueDeclare(
		"go-reat-api-queue", // name
		false,               // durable
		false,               // delete when unused
		false,               // exclusive
		false,               // no-wait
		nil,                 // arguments
	)
	failOnError(err, "Failed to declare a queue")

	return conn, ch, &q
}

func SendTextMessage(ch *amqp.Channel, routingKey string, msg string) {
	sendMessage(ch, routingKey, "text/plain", []byte(msg))
}

func sendMessage(ch *amqp.Channel, routingKey string, contentType string, msg []byte) {
	err := ch.Publish(
		"",         // exchange
		routingKey, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: contentType,
			Body:        msg,
		})
	failOnError(err, "Failed to publish a message")
}
