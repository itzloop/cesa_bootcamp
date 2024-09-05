package boot

import (
	"fmt"
	_ "go-university/internal/app"
	"go-university/internal/config"
	"go-university/internal/db"
)

func BootServer() error {
	dbConfig := config.LoadDBConfig()

	err := db.ConnectDB(dbConfig)
	if err != nil {
		return err
	}

	fmt.Println("Database connected!")

	return nil
}
