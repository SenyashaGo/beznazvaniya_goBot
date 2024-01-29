package main

import "github.com/SenyashaGo/beznazvaniya/internal/services"

func main() {

	bot, err := services.NewBot()
	if err != nil {
		panic(err)
	}
	bot.Polling()
}
