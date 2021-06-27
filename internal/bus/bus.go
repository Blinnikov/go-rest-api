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

func GetChannel(addr string) (*amqp.Connection, *amqp.Channel) {
	conn, err := amqp.Dial(addr)
	failOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")

	return conn, ch
}

func GetQueue(name string, ch *amqp.Channel) *amqp.Queue {
	// We create a Queue to send the message to.
	q, err := ch.QueueDeclare(
		name,  // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	return &q
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

func ReceiveMessage(ch *amqp.Channel, routingKey string) []string {
	msgs, err := ch.Consume(
		routingKey, // routing key
		"",         // consumer
		true,       // autoAck
		false,      // exclusive
		false,      // noLocal
		false,      // noWait
		nil,        // args
	)

	failOnError(err, "Failed to register a consumer")

	var result []string
	for msg := range msgs {
		result = append(result, string(msg.Body))
		log.Printf("Receiving message: %s", msg.Body)
	}
	return result
}
