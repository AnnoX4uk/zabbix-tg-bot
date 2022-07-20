package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

type Configuration struct {
	Chatid int64
	Token  string
}

func getargs() Configuration {
	file, _ := os.Open("config.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	return (configuration)
}

func main() {
	conf := getargs()
	bot, err := tgbotapi.NewBotAPI(conf.Token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)
	msg := tgbotapi.NewMessage(conf.Chatid, "Start notify system")
	bot.Send(msg)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		if update.Message.IsCommand() {
			command := update.Message.Command()
			switch command {
			case "getchat":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprint(update.Message.Chat.ID))
				bot.Send(msg)
			}
		}
	}
}
