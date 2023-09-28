package main

import (
	"github.com/nats-streaming-service-wildberries/internal/app/startup"
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

	natsStreaming := nats_streaming.NewNatsSteamingPublish(config.NatsStreaming.Channel, natsClient)
	natsStreaming.NatsSteamingSubscribe()

}
