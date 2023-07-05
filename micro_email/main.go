package main

import (
	"bytes"
	"fmt"
	"log"
	"micro_email/db"
	"micro_email/pkg/middleware"
	"micro_email/server"
	"net/http"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rabbitmq/amqp091-go"
)

func main() {
	ch := db.NewRabbitMQ()
	go func() {
		getQueQue(ch)
	}()
	server.NewServer()

}
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func getQueQue(ch *amqp091.Channel) {
	msg, err := ch.Consume("q_email", "", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}
	client := http.Client{
		Transport: http.DefaultTransport,
		Timeout:   10 * time.Second,
	}
	for v := range msg {
		fmt.Println(v.Body)
		reader := bytes.NewReader(v.Body)
		req, err := http.NewRequest(http.MethodPost, "http://localhost:6060/email", reader)
		if err != nil {
			fmt.Println(err)
			return
		}
		req.Header.Add("KEY", middleware.SharedKey)
		client.Do(req)
	}
}
