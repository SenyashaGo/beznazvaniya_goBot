package services

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

var startKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Регистрация", "registration"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Оплата", "pay"),
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
		msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, text)
		bot.api.Send(msg)
	case "pay":
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
