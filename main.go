package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

type Configuration struct {
	Chatid     int64
	Token      string
	ListenPort int
}

type recv_message_str struct {
	Message string
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

func SendTgMessage(tgbot *tgbotapi.BotAPI) func(rw http.ResponseWriter, request *http.Request) {
	if tgbot == nil {
		panic("nil tgbot session!")
	}
	return func(rw http.ResponseWriter, request *http.Request) {
		if request.Method == "POST" {
			conf := getargs()
			decoder := json.NewDecoder(request.Body)

			var req recv_message_str

			err := decoder.Decode(&req)
			msg := tgbotapi.NewMessage(conf.Chatid, req.Message)
			tgbot.Send(msg)
			if err != nil {
				panic(err)
			}
			rw.WriteHeader(http.StatusOK)
		} else {
			rw.WriteHeader(http.StatusMethodNotAllowed)
			rw.Write([]byte("Method not allowed."))
		}
	}
}

func main() {
	conf := getargs()
	bot, err := tgbotapi.NewBotAPI(conf.Token)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	// Addr: ":" + fmt.Sprint(HttpPort)
	http.HandleFunc("/", SendTgMessage(bot))
	go http.ListenAndServe(":"+fmt.Sprint(conf.ListenPort), nil)

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
