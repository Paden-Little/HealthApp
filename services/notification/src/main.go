package main

import (
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

var EmailQueue amqp.Queue

func main() {

	time.Sleep(5 * time.Second)
	Send(Email{})
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	declareQueues(conn)

	emailMsgs, err := ch.Consume(
		EmailQueue.Name, // queue
		"",              // consumer
		true,            // auto-ack
		false,           // exclusive
		false,           // no-local
		false,           // no-wait
		nil,             // args
	)
	failForError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range emailMsgs {
			log.Printf("Received a message: %s", d.Body)
			SendEmail()
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func declareQueues(conn *amqp.Connection) {
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	email, err := ch.QueueDeclare(
		"sendEmail", // name
		false,       // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	failForError(err, "Failed to declare a queue")
	EmailQueue = email
}

func failForError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
