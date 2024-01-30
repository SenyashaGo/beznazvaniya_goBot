package services

import (
	"github.com/SenyashaGo/beznazvaniya/internal/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

//type RegistrationProcess struct {
//	sync.Mutex
//	InProgress bool
//}

//var registrationProcess = &RegistrationProcess{}

var startKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Регистрация", "registration"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Оплата", "payment"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Наши контакты", "contacts"),
	),
)

func (bot *Bot) StartMenu(update tgbotapi.Update) {
	switch update.CallbackData() {
	case "registration":
		text := "нажата кнопка регистрации"
		callback := tgbotapi.NewCallback(update.CallbackQuery.ID, text)
		if _, err := bot.api.Request(callback); err != nil {
			log.Println(err)
		}
		//// Создаем канал для ожидания ответа от пользователя
		//userResponseChan := make(chan string)
		//bot.UserResponses[update.CallbackQuery.Message.Chat.ID] = userResponseChan
		//
		//// Запускаем RegUsers в горутине
		//go func() {
		chatID := update.CallbackQuery.Message.Chat.ID
		rez, exists := Users[chatID]

		if !exists {
			rez = &models.UserReg{State: StateAskFullName} // Initialize state to ask for full name
			Users[chatID] = rez
		}
		bot.RegUsers(update, rez, chatID)
		//// Закрываем канал после завершения RegUsers
		//close(userResponseChan)
		//}()
		//
		//// Ожидаем ответа от пользователя, проверяя канал
		//select {
		//case userInput := <-userResponseChan:
		//	log.Printf("Received user input: %s", userInput)
		//}

	case "payment":
		text := "Нажата кнопка оплаты"
		callback := tgbotapi.NewCallback(update.CallbackQuery.ID, text)
		if _, err := bot.api.Request(callback); err != nil {
			log.Println(err)
		}
		msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, text)
		bot.api.Send(msg)
	case "contacts":
		text := "Нажата кнопка контактов"
		callback := tgbotapi.NewCallback(update.CallbackQuery.ID, text)
		if _, err := bot.api.Request(callback); err != nil {
			log.Println(err)
		}
		msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, text)
		bot.api.Send(msg)
	}
}
