package main

import (
	"flag"
	"log"

	"github.com/m-kuzmin/golang-telegram-bot/internal/clients/telegram"
)

func main() {
	client := telegram.New("api.telegram.org", mustToken())
}

// Gets the token from the environment. The token is passed as CLI args
// to the bot exe and if it hasn't been found this funciton will panic
func mustToken() string {
	const tokenFlag = "tg-token"
	tg_token := flag.String(tokenFlag, "", "A telegram token from BotFather")
	flag.Parse()

	// token ptr could be nil (*nil == panic), but no token is fatal anyway
	if *tg_token == "" {
		log.Fatal("No telegram token specified, use", "-"+tokenFlag, "TOKEN")
	}
	return *tg_token
}
