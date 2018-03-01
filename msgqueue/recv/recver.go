package main

import (
	"github.com/streadway/amqp"
	"simples/msgqueue"
	"os"
	"fmt"
)

func main() {
	conn, err := amqp.Dial("amqp://root:root@localhost:5673/")
	checkErr(err)
	defer conn.Close()

	ch, err := conn.Channel()
	checkErr(err)
	defer ch.Close()

	ch.ExchangeDeclare(
		msgqueue.ExchangeName,
		"topic",
		true,  // auto-deleted
		false, // internal
		false, // no-wait
		false, // arguments
		nil,
	)

	q, err := ch.QueueDeclare("", false, false, true, false, nil)
	checkErr(err)

	if len(os.Args) < 2 {
		os.Exit(0)
	}

	for _, rk := range os.Args[1:] {

		ch.QueueBind(
			q.Name,
			rk,
			msgqueue.ExchangeName,
			false,
			nil)
	}

	msg, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	checkErr(err)
	for d := range msg {
		fmt.Println(string(d.Body))
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
