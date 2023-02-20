package main

import (
	"context"
	"fmt"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	q, err := ch.QueueDeclare("go-q1", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	go consume("c1", conn, q.Name)
	go consume("c2", conn, q.Name)

	for i := 0; i < 1000; i++ {
		err = ch.PublishWithContext(context.Background(), "", q.Name, false, false, amqp.Publishing{
			Body: []byte(fmt.Sprintf("%d", i)),
		})

		if err != nil {
			fmt.Printf("Error publishing %+v\n", err)
		}

		time.Sleep(time.Millisecond * 200)
	}
}

func consume(name string, conn *amqp.Connection, q string) {
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	msg, err := ch.Consume(q, name, true, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	for ms := range msg {
		fmt.Printf("%s : %s\n", name, ms.Body)
	}
}
