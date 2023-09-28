package cache

import "fmt"

import (
	"github.com/nats-streaming-service-wildberries/internal/models"
)

type Cache struct {
	cache map[string]models.Order
}

func NewCache() *Cache {
	return &Cache{cache: make(map[string]models.Order)}
}

func (c *Cache) GetOrder(orderUid string) (*models.Order, error) {
	order, ok := c.cache[orderUid]
	if ok {
		return &order, nil
	}
	return nil, fmt.Errorf("order not found")
}

func (c *Cache) AddOrders(orders []models.Order) error {
	for _, order := range orders {
		c.cache[order.OrderUid] = order
	}
	return nil
}

func (c *Cache) AddOrder(order models.Order) {
	c.cache[order.OrderUid] = order
}
