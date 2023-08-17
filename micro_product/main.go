package main

import (
	"fmt"
	"log"
	"micro_product/server"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	if os.Getenv("MODE") == "GRPC" {
		fmt.Println("GRPC mode")
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
