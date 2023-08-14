package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"micro_email/db"
	"micro_email/modules/proto"
	"micro_email/pkg/crypto"
	"micro_email/pkg/middleware"
	"micro_email/server"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	ch := db.NewRabbitMQ()
	go func() {
		getQueQue(ch)
	}()
	if os.Getenv("MODE") == "GRPC" {
		fmt.Println("MODE GRPC API")
		server.NewGRPCServer()
	} else {
		server.NewServer()
	}

}
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func serviceClient() proto.EmailSrvClient {
	port := ":6060"
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		fmt.Println("could not connect to", port, err)
	}
	return proto.NewEmailSrvClient(conn)

}
func getQueQue(ch *amqp091.Channel) {
	msg, err := ch.Consume("q_email", "", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}
	if os.Getenv("MODE") == "GRPC" {
		emailClient := serviceClient()

		mEmail := proto.Email{}
		for v := range msg {

			err := protojson.Unmarshal(v.Body, &mEmail)

			if err != nil {
				log.Fatal(err)
				return
			}

			mt := metadata.MD{
				"key": {crypto.SharedKey},
			}
			c := metadata.NewOutgoingContext(context.Background(), mt)
			res, err := emailClient.SendEmail(c, &mEmail)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(res)

		}
	} else {
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

}
