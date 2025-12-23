package main

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	fmt.Println("RabbitMQ consumer")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println("Failed Initialization Broker Connection", err)
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		fmt.Println("Failed Initialization Channel", err)
		panic(err)
	}
	defer ch.Close()
	msg, err := ch.Consume("GO", "", false, false, false, false, nil)
	if err != nil {
		fmt.Println("Failed Initialization Consumer ", err)
	}
	forever := make(chan bool)
	go func() {
		for d := range msg {
			fmt.Printf("Received a message: %s\n", d.Body)
		}
	}()
	fmt.Println("Connected to the RabbitMQ server")
	<-forever
}
