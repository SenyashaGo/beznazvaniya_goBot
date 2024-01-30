package services

import (
	"github.com/SenyashaGo/beznazvaniya/internal/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

const (
	StateAskFullName = iota
	StateAskDateOfBirth
	StateAskPhoneNumber
	StateAskContactLink
	StateRegistrationComplete
	StateRegStatusFalse = false
	StateRegStatusTrue  = true
)

var Users = make(map[int64]*models.UserReg)

func (bot *Bot) RegUsers(update tgbotapi.Update, rez *models.UserReg, chatId int64) {
	//rez := Users[update.CallbackQuery.Message.Chat.ID]
	//if rez == nil {
	//	// Handle the case when rez is nil, e.g., by creating a new instance.
	//	rez = &models.UserReg{}
	//	Users[update.CallbackQuery.Message.Chat.ID] = rez
	//}
	//
	//textToUser := "Введите свое ФИО\n" +
	//	"_Формат: Иванов Иван Иванович_"
	//msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, textToUser)
	//msg.ParseMode = tgbotapi.ModeMarkdown
	//msg.ReplyMarkup = cancelKeyboard
	//_, err := bot.api.Send(msg)
	//if err != nil {
	//	log.Println("Can't send message to User (registration)")
	//}
	//if rez.FullName != nil {
	//
	//	textToUser := "Введите дату своего рождения\n" +
	//		"_Формат: 02.05.2004_"
	//	msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, textToUser)
	//	msg.ParseMode = tgbotapi.ModeMarkdown
	//	_, err := bot.api.Send(msg)
	//	if err != nil {
	//		log.Println("Can't send message to User (registration)")
	//	}
	//} else {
	//	rez.FullName = &update.CallbackQuery.Message.Text
	//}
	//if rez.DateOfBirth != nil {
	//
	//	textToUser := "Введите свой номер телефона\n" +
	//		"_Формат: 89856969271_"
	//	msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, textToUser)
	//	msg.ParseMode = tgbotapi.ModeMarkdown
	//	_, err := bot.api.Send(msg)
	//	if err != nil {
	//		log.Println("Can't send message to User (registration)")
	//	}
	//} else {
	//	rez.DateOfBirth = &update.CallbackQuery.Message.Text
	//}
	//if rez.PhoneNumber != nil {
	//
	//	textToUser := "Введите ссылку на свой TG/VK\n" +
	//		"_Формат: https://vk.com/g.kolb_"
	//	msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, textToUser)
	//	msg.ParseMode = tgbotapi.ModeMarkdown
	//	_, err := bot.api.Send(msg)
	//	if err != nil {
	//		log.Println("Can't send message to User (registration)")
	//	}
	//} else {
	//	rez.PhoneNumber = &update.CallbackQuery.Message.Text
	//}
	//if rez.ContactLink != nil {
	//	rez.ContactLink = &update.CallbackQuery.Message.Text
	//}
	//if rez.FullName != nil && rez.DateOfBirth != nil && rez.PhoneNumber != nil && rez.ContactLink != nil {
	//	textToUser := "Вы упешно авторизированы!"
	//	msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, textToUser)
	//	bot.api.Send(msg)
	//}

	switch rez.State {
	case StateAskFullName:
		textToUser := "Введите свое ФИО\n" +
			"_Формат: Иванов Иван Иванович_"
		msg := tgbotapi.NewMessage(chatId, textToUser)
		msg.ParseMode = tgbotapi.ModeMarkdown
		msg.ReplyMarkup = cancelKeyboard
		_, err := bot.api.Send(msg)
		if err != nil {
			log.Println("Can't send message to User (registration)")
		}
		rez.State = StateAskDateOfBirth
		rez.StateReg = StateRegStatusFalse
	case StateAskDateOfBirth:
		rez.FullName = &update.Message.Text
		textToUser := "Введите дату своего рождения\n" +
			"_Формат: 02.05.2004_"
		msg := tgbotapi.NewMessage(chatId, textToUser)
		msg.ParseMode = tgbotapi.ModeMarkdown
		_, err := bot.api.Send(msg)
		if err != nil {
			log.Println("Can't send message to User (registration)")
		}
		rez.State = StateAskPhoneNumber
		rez.StateReg = StateRegStatusFalse

	case StateAskPhoneNumber:
		rez.DateOfBirth = &update.Message.Text
		textToUser := "Введите свой номер телефона\n" +
			"_Формат: 89856969271_"
		msg := tgbotapi.NewMessage(chatId, textToUser)
		msg.ParseMode = tgbotapi.ModeMarkdown
		_, err := bot.api.Send(msg)
		if err != nil {
			log.Println("Can't send message to User (registration)")
		}
		rez.State = StateAskContactLink
		rez.StateReg = StateRegStatusFalse

	case StateAskContactLink:
		rez.PhoneNumber = &update.Message.Text
		textToUser := "Введите ссылку на свой TG/VK\n" +
			"_Формат: https://vk.com/g.kolb_"
		msg := tgbotapi.NewMessage(chatId, textToUser)
		msg.ParseMode = tgbotapi.ModeMarkdown
		_, err := bot.api.Send(msg)
		if err != nil {
			log.Println("Can't send message to User (registration)")
		}
		rez.State = StateRegistrationComplete
		rez.StateReg = StateRegStatusFalse

	case StateRegistrationComplete:
		rez.ContactLink = &update.Message.Text
		if rez.FullName != nil && rez.DateOfBirth != nil && rez.PhoneNumber != nil && rez.ContactLink != nil {
			textToUser := "Вы успешно авторизованы!"
			msg := tgbotapi.NewMessage(chatId, textToUser)
			bot.api.Send(msg)
			rez.StateReg = StateRegStatusTrue
			delete(Users, chatId)
		}

	default:
		log.Println("Unknown state in registration process")
	}
}
