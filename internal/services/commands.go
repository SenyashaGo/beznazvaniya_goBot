package services

import (
	"github.com/SenyashaGo/beznazvaniya/internal/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
)

func (bot *Bot) Commands(update tgbotapi.Update) {
	switch update.Message.Command() {
	case "start":
		rez := &models.User{}
		// получаем полное имя
		rez.UserFirstName = &update.Message.Chat.FirstName
		rez.UserLastName = &update.Message.Chat.LastName
		// получаем юзернейм
		rez.Username = &update.Message.Chat.UserName
		// получаем чат-айди
		userChatId := int(update.Message.Chat.ID)
		rez.UserChatID = &userChatId
		parseAdmin, err := strconv.ParseInt(tokensJSON.AdminChatId, 10, 0)
		if err != nil {
			log.Println("Can't get Admin Chat ID")
		}
		textToAdmin := "❗НОВЫЙ ПОЛЬЗОВАТЕЛЬ❗\nChatID: " + strconv.Itoa(*rez.UserChatID) + "\n" +
			"Name: " + *rez.UserFirstName + " " + *rez.UserLastName + "\nUsername: @" + *rez.Username + "\n"
		_, err = bot.api.Send(tgbotapi.NewMessage(parseAdmin, textToAdmin))
		if err != nil {
			log.Println("Can't send message to Admin Chat")
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Здравствуйте! Сначала Вам необходимо пройти "+
			"регистрацию, для этого нажмите кнопку регистрации")
		msg.ReplyMarkup = startKeyboard
		bot.api.Send(msg)
		return
	case "registration":
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Отправьте мне токен на свою Яндекс Музыку,"+
			" если вы не знаете как его получить, нажмите кнопку инструкция")
		bot.api.Send(msg)
		//Users[update.Message.Chat.ID] = &models.User{}
		return
	case "help":
		bot.api.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Здесь будет реализована поддержка"))
		return
	}
}
