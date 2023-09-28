package nats_streaming

import (
	"github.com/nats-io/stan.go"
	"github.com/nats-streaming-service-wildberries/internal/business_logic"
	"log"
)

type NatsSteamingSubscribe struct {
	natsClient stan.Conn
	channel    string
	client     *business_logic.Client
}

func NewNatsSteamingSubscribe(channel string, natsClient stan.Conn, client *business_logic.Client) *NatsSteamingSubscribe {

	natsSteamingClient := NatsSteamingSubscribe{
		natsClient: natsClient,
		channel:    channel,
		client:     client,
	}
	return &natsSteamingClient
}

func (n *NatsSteamingSubscribe) NatsSteamingSubscribe() {
	_, err := n.natsClient.Subscribe(
		n.channel, func(m *stan.Msg) {
			err := n.client.AddOrder(m.Data)
			if err != nil {
				log.Printf("Error order adding %v", err)
			}
		})
	if err != nil {
		log.Fatal(err)
	}
}
