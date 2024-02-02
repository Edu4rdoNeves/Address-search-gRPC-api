package main

import (
	"fmt"
	"log"

	"github.com/Edu4rdoNeves/Address-search-gRPC-api/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Erro ao carregar o arquivo .env:", err)
		return
	}

	err = server.GrpcServer()
	if err != nil {
		log.Fatal("Failed to initialzate server. Error:", err)
	}
}
