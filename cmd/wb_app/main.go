package main

import (
	"github.com/nats-streaming-service-wildberries/internal/app/startup"
	"github.com/nats-streaming-service-wildberries/internal/business_logic"
	"github.com/nats-streaming-service-wildberries/internal/http"
	"github.com/nats-streaming-service-wildberries/internal/service/nats_streaming"
	"log"
	"os"
)

func main() {

	configPath := os.Getenv("CONFIG_PATH")

	config, err := startup.NewConfig(configPath)
	if err != nil {
		log.Fatalf("failed to Config %v", err)
	}

	// Открытие соединения с postgres
	dbClient, err := startup.DatabaseConnect(config.Database)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	log.Printf("connecting to the database: success")

	defer func() {
		dbInstance, _ := dbClient.DB()
		if err = dbInstance.Close(); err != nil {
			log.Fatalf("failed to close database: %v", err)
		}
	}()

	// Клиент для реализации бизнес-логики
	client := business_logic.NewOrderClient(dbClient)
	err = client.Start()
	if err != nil {
		log.Fatalf("failed Start: %v", err)
	}

	natsClient, err := startup.NatsStreamingConnect(config.NatsStreaming)
	if err != nil {
		log.Fatalf("failed to connect to nats-streaming: %v", err)
	}

	log.Printf("connecting to the nats-streaming: success")

	defer func() {
		if err = natsClient.Close(); err != nil {
			log.Fatalf("failed to close nats-streaming: %v", err)
		}
	}()

	// Клиент для nats-streaming
	natsStreaming := nats_streaming.NewNatsSteamingSubscribe(config.NatsStreaming.Channel, natsClient, client)

	natsStreaming.NatsSteamingSubscribe()

	// Запуск http сервера
	httpRouter := http.NewHttpRouter(config.Http, client)
	log.Printf("start http-server")
	err = httpRouter.Start()
	if err != nil {
		log.Fatalf("failed to listen the port: %d %v", config.Http.Port, err)
	}
}
