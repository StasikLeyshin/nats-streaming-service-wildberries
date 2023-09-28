package nats_streaming

import (
	"encoding/json"
	"fmt"
	"github.com/nats-io/stan.go"
	"github.com/nats-streaming-service-wildberries/internal/models"
	"log"
	"os"
	"time"
)

type NatsSteamingPublish struct {
	natsClient stan.Conn
	channel    string
}

func NewNatsSteamingPublish(channel string, natsClient stan.Conn) *NatsSteamingPublish {

	natsSteamingClient := NatsSteamingPublish{
		natsClient: natsClient,
		channel:    channel,
	}
	return &natsSteamingClient
}

func (n *NatsSteamingPublish) NatsSteamingSubscribe() {
	file, err := os.ReadFile("model.json")
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}

	var order models.Order

	err = json.Unmarshal(file, &order)
	if err != nil {
		log.Fatalf("failed Unmarshal: %v", err)
	}
	for {
		orderUid := fmt.Sprintf("%d", int(time.Now().Unix()))
		order.OrderUid = orderUid

		fmt.Println("Send Order Uid: ", order.OrderUid)

		bytes, err := json.Marshal(order)
		if err != nil {
			log.Println("ERROR: json.Marshal:", err)
			continue
		}

		err = n.natsClient.Publish(n.channel, bytes)
		if err != nil {
			log.Println("ERROR: conn.Publish:", err)
			continue
		}
		time.Sleep(time.Second * 4)
	}
}
