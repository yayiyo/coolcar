package main

import (
	"context"
	"fmt"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

const exchange = "go-ex"

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	err = ch.ExchangeDeclare(exchange, "fanout", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	go subscribe(conn, "a--", exchange)
	go subscribe(conn, "b==", exchange)
	go subscribe(conn, "c>>", exchange)

	for i := 0; i < 1000; i++ {
		err = ch.PublishWithContext(context.Background(), exchange, "", false, false, amqp.Publishing{
			Body: []byte(fmt.Sprintf("%d", i)),
		})

		if err != nil {
			fmt.Printf("Error publishing %+v\n", err)
		}

		time.Sleep(time.Millisecond * 200)
	}
}

func subscribe(conn *amqp.Connection, cname, ex string) {
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare("", false, true, false, false, nil)
	if err != nil {
		panic(err)
	}
	defer ch.QueueDelete(q.Name, false, false, false)

	err = ch.QueueBind(q.Name, "", ex, false, nil)
	if err != nil {
		panic(err)
	}
	consume(cname, ch, q.Name)
}

func consume(name string, ch *amqp.Channel, q string) {
	msg, err := ch.Consume(q, name, true, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	for ms := range msg {
		fmt.Printf("%s : %s\n", name, ms.Body)
	}
}
