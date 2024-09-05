package main

import (
	"context"
	"fmt"
	"go-university/internal/config"
	"go-university/internal/db"
	"log"
	"os"
)

func main() {
	dbConfig := config.LoadDBConfig()

	ctx := context.Background()

	err := db.ConnectDB(dbConfig)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	fmt.Println("Database connected!")
}
