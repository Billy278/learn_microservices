package db

import amqp "github.com/rabbitmq/amqp091-go"

func NewRabbitMQ() *amqp.Channel {
	// 1. Buat koneksi RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	//2.buat channel
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	//nnti coba tanpa excange
	// 3. Deklarasikan exchange
	err = ch.ExchangeDeclare("ex_learn_micro", "direct", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}
	//4.declatation queque
	q, err := ch.QueueDeclare("q_email", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}
	//coba tanpa quequebind
	//5.bind queque to exchange
	err = ch.QueueBind(q.Name, "PWSD", "ex_learn_micro", false, nil)
	if err != nil {
		panic(err)
	}
	return ch
}
