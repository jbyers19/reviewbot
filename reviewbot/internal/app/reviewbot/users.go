package reviewbot

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Customer struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	ChatID    int64  `json:"id"`
	UserID    int64  `json:"user_id"`
}

type Customers struct {
	CustomersMap map[int64]*Customer
}

// AddCustomer adds a customer to the Customers map.
func (c *Customers) AddCustomer(m *tgbotapi.Message) error {
	// Make sure we have all the required data.
	if m.From.FirstName == "" || m.From.LastName == "" || m.From.ID == 0 || m.Chat.ID == 0 {
		return fmt.Errorf("missing required customer data: %+v, %+v", m.From, m.Chat)
	}

	// If the customer is already in the map, do nothing.
	if _, ok := c.CustomersMap[m.From.ID]; ok {
		return nil
	}

	customer := &Customer{
		FirstName: m.From.FirstName,
		LastName:  m.From.LastName,
		// In a real scenario we would have to consider multiple chat IDs,
		// but for now we assume there is only a single chat per user.
		ChatID: m.Chat.ID,
	}
	c.CustomersMap[m.From.ID] = customer
	log.Printf("added customer %s %s", m.From.FirstName, m.From.LastName)
	return nil
}

// GetCustomerByName returns a customer from the Customers map.
func (c *Customers) GetCustomerByName(firstName, lastName string) Customer {
	for _, customer := range c.CustomersMap {
		if customer.FirstName == firstName && customer.LastName == lastName {
			return *customer
		}
	}

	return Customer{}
}
