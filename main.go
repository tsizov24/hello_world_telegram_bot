package main

import (
	"fmt"
	"gopkg.in/telegram-bot-api.v4"
	"net/http"
)

// Input your data to constants
const (
	BotToken    = ""
	WebhookUrl  = ""
	SslCertFile = ""
	SslKeyFile  = ""
)

func main() {
	// Creating bot and authorizing
	bot, err := tgbotapi.NewBotAPI(BotToken)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Authorized as %s\n", bot.Self.UserName)

	if _, err = bot.SetWebhook(tgbotapi.NewWebhook(WebhookUrl)); err != nil {
		panic(err)
	}

	// Starting server
	go http.ListenAndServeTLS(":443", SslCertFile, SslKeyFile, nil)
	fmt.Println("start listen :443")

	// Start getting user's inputing data infinitely
	updates := bot.ListenForWebhook("/")
	for update := range updates {
		bot.Send(tgbotapi.NewMessage(
			update.Message.Chat.ID,
			"Hello, World!",
		))
	}
}
