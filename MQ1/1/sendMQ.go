package main

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	fmt.Println("RabbitMQ producer")

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println("amqp.Dial()", err)
		return
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		fmt.Println("Channel()", err)
		return
	}
	defer ch.Close()
	q, err := ch.QueueDeclare("GO", false, false, false, false, nil)
	if err != nil {
		fmt.Println("QueueDeclare()", err)
		return
	}
	fmt.Println("Queue:", q)
	massage := "Writing to RabbitMQ!"
	err = ch.PublishWithContext(nil, "", "GO", false, false, amqp.Publishing{ContentType: "text/plain", Body: []byte(massage)})
	if err != nil {
		fmt.Println("Publish()", err)
		return
	}
	fmt.Println("Massage published to Queue")
}
