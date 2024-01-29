package services

import (
	"encoding/json"
	"github.com/SenyashaGo/beznazvaniya/internal/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"io"
	"log"
	"os"
)

var Users = make(map[int64]*models.User)

type Bot struct {
	token string
	api   *tgbotapi.BotAPI
}

type tokens struct {
	BotToken    string `json:"bot_token"`
	AdminChatId string `json:"admin_chat_id"`
}

func NewBot() (*Bot, error) {
	open, errOpen := os.Open("configs/config_bot.json")
	if errOpen != nil {
		log.Println("Can not open JSON, check the directory where the file is located")
	}

	bytes, errRead := io.ReadAll(open)
	if errRead != nil {
		log.Println("Can't read JSON file")
	}

	var tokensJSON tokens

	errUn := json.Unmarshal(bytes, &tokensJSON)
	if errUn != nil {
		log.Println("Can't read JSON file (Unmarshal)")
	}

	bot, err := tgbotapi.NewBotAPI(tokensJSON.BotToken)
	if err != nil {
		log.Println("Bot can't get token from JSON file")
	}

	bot.Debug = true
	return &Bot{
		token: tokensJSON.BotToken,
		api:   bot,
	}, nil
}

func (bot *Bot) Polling() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.api.GetUpdatesChan(u)

	for update := range updates {
		if update.Message.IsCommand() {
			bot.Commands(update)
			continue
		}
	}
}
