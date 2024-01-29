package services

import (
	"encoding/json"
	"github.com/SenyashaGo/beznazvaniya/internal/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"io"
	"log"
	"os"
	"strconv"
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
		if update.CallbackQuery != nil {
			bot.StartMenu(update)
		} else if update.Message.IsCommand() {
			bot.Commands(update)
			//continue
		}
		if update.Message != nil {
			if _, ok := Users[update.Message.Chat.ID]; !ok {
				continue
			}
			rez := Users[update.Message.Chat.ID]
			if rez.UserFirstName == nil {
				// получаем полное имя
				rez.UserFirstName = &update.Message.Chat.FirstName
				rez.UserLastName = &update.Message.Chat.LastName
				//получаем юзернейм
				rez.Username = &update.Message.Chat.UserName
				//получаем чат-айди
				userChatId := int(update.Message.Chat.ID)
				rez.UserChatID = &userChatId
			}

			parseAdmin, err := strconv.ParseInt(os.Getenv("ADMIN_CHAT_ID"), 10, 0)

			if err != nil {
				panic(err)
			}

			textToAdmin := "❗НОВЫЙ ПОЛЬЗОВАТЕЛЬ❗\nChatID: " + strconv.Itoa(*rez.UserChatID) + "\n" +
				"Name: " + *rez.UserFirstName + " " + *rez.UserLastName + "\nUsername: @" + *rez.Username + "\n"
			_, err = bot.api.Send(tgbotapi.NewMessage(parseAdmin, textToAdmin))
			delete(Users, int64(*rez.UserChatID))
			if err != nil {
				log.Println(err) // logrus
			}

			_, err = bot.api.Send(tgbotapi.NewMessage(int64(*rez.UserChatID), "Вы успешно авторизированы!"))
			if err != nil {
				log.Println(err)
			}
		}
	}
}
