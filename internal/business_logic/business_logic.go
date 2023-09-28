package business_logic

import (
	"encoding/json"
	"fmt"
	"github.com/nats-streaming-service-wildberries/internal/models"
	"github.com/nats-streaming-service-wildberries/internal/repository"
	"github.com/nats-streaming-service-wildberries/internal/service/cache"
	"gorm.io/gorm"
	"log"
)

type Client struct {
	db              *gorm.DB
	cache           *cache.Cache
	orderRepository *repository.OrderRepository
}

func NewOrderClient(db *gorm.DB) *Client {
	return &Client{
		db:              db,
		cache:           cache.NewCache(),
		orderRepository: repository.NewOrderRepository(db),
	}
}

func (c *Client) Start() error {
	err := c.orderRepository.OrderAutoMigrate()
	if err != nil {
		return err
	}
	orders, err := c.orderRepository.GetAllOrder()
	if err != nil {
		return err
	}
	c.cache.AddOrders(orders)
	return nil
}

func (c *Client) GetOrder(orderUid string) (*models.Order, error) {
	order, err := c.cache.GetOrder(orderUid)
	if err != nil {
		log.Printf("order not found %v", err)
		return nil, fmt.Errorf("order not found")
	}

	return order, nil
}

func (c *Client) AddOrder(data []byte) error {
	var order models.Order
	err := json.Unmarshal(data, &order)

	if err != nil {
		log.Printf("error happened in JSON unmarshal. Err: %v", err)
	}

	fmt.Println("Get Order Uid: ", order.OrderUid)
	_, err = c.cache.GetOrder(order.OrderUid)
	if err == nil {
		log.Printf("such an order already exists. Err: %v", err)
		return fmt.Errorf("such an order already exists")
	}

	err = c.orderRepository.AddOrder(&order)
	if err != nil {
		log.Printf("error create order: %v", err)
		return err
	}
	c.cache.AddOrder(order)
	return nil
}
