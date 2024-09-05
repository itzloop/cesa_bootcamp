package main

import (
	"go-university/api"
	"go-university/internal/boot"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	err = boot.BootServer()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	router := api.NewRouter()

	err = router.Serve()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
