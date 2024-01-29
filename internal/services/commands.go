package services

import (
	"github.com/SenyashaGo/beznazvaniya/internal/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (bot *Bot) Commands(update tgbotapi.Update) {
	switch update.Message.Command() {
	case "start":
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Здравствуйте! Сначала Вам необходимо пройти "+
			"регистрацию, для этого нажмите команду /reg")
		//msg.ReplyMarkup = startKeyboard
		bot.api.Send(msg)
		return
	case "reg":
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Отправьте мне токен на свою Яндекс Музыку,"+
			" если вы не знаете как его получить, нажмите кнопку инструкция")
		bot.api.Send(msg)
		Users[update.Message.Chat.ID] = &models.User{}
		return
	case "instruction":

	case "install":
		//rez := Users[update.Message.Chat.ID]
		//accessToken := *rez.YaToken
		//client := yamusic.NewClient(yamusic.AccessToken(int(update.Message.Chat.ID), accessToken))
		//print(client.Playlists())
	case "help":
		bot.api.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Здесь будет реализована поддержка"))
		return
	}
}
