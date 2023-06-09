package main

import (
	"flag"
	"log"

	"github.com/m-kuzmin/golang-telegram-bot/internal/clients/telegram"
)

func main() {
	client := telegram.New("api.telegram.org", mustToken())
	offset := 0
	for {
		updates, err := client.Updates(offset, 10)
		if err != nil {
			log.Println(err)
		} else {
			for _, u := range updates {
				log.Println(u.Message.Text)
				if err := client.SendMessage(u.Message.Chat.ID, u.Message.Text); err != nil {
					log.Println("While sending respoce:", err)
				}
			}
			if len(updates) != 0 {
				// All updates have an ordered id.
				// The API requires the offset to be id+1
				offset = updates[len(updates)-1].ID + 1
			}
		}
	}
}

// Gets the token from the environment. The token is passed as CLI args
// to the bot exe and if it hasn't been found this funciton will panic
func mustToken() string {
	const tokenFlag = "tg-token"
	tgToken := flag.String(tokenFlag, "", "A telegram token from BotFather")
	flag.Parse()

	// token ptr could be nil (*nil == panic), but no token is fatal anyway
	if *tgToken == "" {
		log.Fatal("No telegram token specified, use ", "-"+tokenFlag, " TOKEN")
	}
	return *tgToken
}
