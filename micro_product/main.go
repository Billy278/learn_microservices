package main

import (
	"log"
	"micro_product/server"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	if os.Getenv("MODE") == "GRPC" {
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
