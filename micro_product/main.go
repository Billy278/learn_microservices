package main

import (
	"log"
	"micro_product/server"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	server.NewServer()
}
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
