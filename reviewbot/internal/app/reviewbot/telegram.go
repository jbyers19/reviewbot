package reviewbot

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	APIKey = os.Getenv("TELEGRAM_API_KEY")
	bot    *tgbotapi.BotAPI

	// TemplatesDB and CustomerDB holds templates and customer data.
	TemplatesDB *Templates
	CustomerDB  *Customers
)

// StartTelegramBot starts the Telegram bot.
func StartTelegramBot(templatesDB *Templates, customerDB *Customers) {
	log.Printf("Starting Telegram bot service")
	var err error
	bot, err = tgbotapi.NewBotAPI(APIKey)
	if err != nil {
		log.Panic(err)
	}

	TemplatesDB = templatesDB
	CustomerDB = customerDB

	bot.Debug = false

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	ctx := context.Background()
	updates := bot.GetUpdatesChan(u)

	go receiveUpdates(ctx, updates)

	// Wait for SIGINT (Ctrl+C) or SIGTERM (docker stop)
	<-ctx.Done()
}

// receiveUpdates receives updates from the Telegram bot.
func receiveUpdates(ctx context.Context, updates tgbotapi.UpdatesChannel) {
	for {
		select {
		// Stop looping if ctx is cancelled
		case <-ctx.Done():
			return
		// Receive update from channel and then handle it
		case update := <-updates:
			if update.Message != nil {
				handleMessage(update.Message)
			}
		}
	}
}

// handleMessage handles a message received from the Telegram bot.
func handleMessage(message *tgbotapi.Message) {
	// Add the customer to the customers map if they don't already exist.
	err := CustomerDB.AddCustomer(message)
	if err != nil {
		log.Printf("error adding customer: %s", err.Error())
		return
	}

	// Different handlers can be attached here to control the bot's responses.
	msg, err := messageTemplateHandler(message)
	if err != nil {
		log.Printf("error handling message: %s", err.Error())
		return
	}

	if msg == "" {
		return
	}

	err = SendTelegramMessage(message.From.FirstName, message.From.LastName, msg)
	if err != nil {
		log.Printf("error sending message: %s", err.Error())
	}
}

// SendTelegramMessage sends a message to a customer.
func SendTelegramMessage(firstName, lastName, msgText string) error {
	customer := CustomerDB.GetCustomerByName(firstName, lastName)
	if customer.ChatID == 0 {
		return fmt.Errorf("customer %s %s has not started a chat with the Telegram bot", firstName, lastName)
	}

	if msgText == "" {
		return fmt.Errorf("message text is empty")
	}

	msg := tgbotapi.NewMessage(customer.ChatID, msgText)
	_, err := bot.Send(msg)
	if err != nil {
		log.Printf("error sending message: %s", err.Error())
		return err
	}

	return nil
}

// messageTemplateHandler handles a message received from the Telegram bot
// and returns a message generated from a template.
func messageTemplateHandler(message *tgbotapi.Message) (string, error) {
	for templName, templ := range TemplatesDB.TemplatesMap {
		for _, trigger := range templ.Triggers {
			if strings.Contains(strings.ToLower(message.Text), trigger) {
				str, err := TemplatesDB.GenerateMessage(templName, message.From.FirstName, message.From.LastName)
				if err != nil {
					return "", err
				}
				return str, nil
			}
		}
	}

	return "", nil
}
