package main

import (
	"github.com/streadway/amqp"
	"simples/msgqueue"
	"strings"
	"os"
)

func main() {
	conn, err := amqp.Dial("amqp://root:root@localhost:5673/")
	checkErr(err)
	defer conn.Close()

	ch, err := conn.Channel()
	checkErr(err)
	defer ch.Close()

	err = ch.ExchangeDeclare(
		msgqueue.ExchangeName, // name
		"topic",               // type
		true,                  // durable
		false,                 // auto-deleted
		false,                 // internal
		false,                 // no-wait
		nil,                   // arguments
	)
	checkErr(err)

	body := body(os.Args)
	rk := routingKey(os.Args)

	err = ch.Publish(
		msgqueue.ExchangeName,
		rk,
		false,
		false,
		amqp.Publishing{
			Type: "text/plain",
			Body: []byte(body),
		},
	)
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func body(args []string) string {
	var s string
	if len(args) < 3 || args[2] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[2:], " ")
	}
	return s
}

func routingKey(args []string) string {
	var s string
	if len(args) < 2 || os.Args[1] == "" {
		s = "h5pay.test"
	} else {
		s = args[1]
	}
	return s
}
