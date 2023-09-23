package main

import (
	"fmt"
	"github.com/nats-streaming-service-wildberries/internal/app/startup"
)

func main() {

	config, err := startup.NewConfig("./config/config.local.yaml")

	// Открытие соединения с postgres
	dbClient, err := startup.DatabaseConnect(config.Database)
	if err != nil {
		fmt.Errorf("failed to connect to database: %w", err)
	}

	defer func() {
		dbInstance, _ := dbClient.DB()
		if err = dbInstance.Close(); err != nil {
			fmt.Errorf("failed to close database: %w", err)
		}
	}()

	// Запуск http сервера
	//httpRouter := http.NewHttpRouter(config.Http, logger)
	//err := httpRouter.ListenAndServe(":3333", mux)
}
