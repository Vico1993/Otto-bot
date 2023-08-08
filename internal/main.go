package main

import (
	"log"
	"os"
	"time"

	"github.com/Vico1993/Otto-bot/internal/handles"
	"github.com/subosito/gotenv"
	tele "gopkg.in/telebot.v3"
)

func main() {
	// load .env file if any otherwise use env set
	_ = gotenv.Load()

	pref := tele.Settings{
		Token:  os.Getenv("TELEGRAM_BOT_TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Handles
	b.Handle("/hello", handles.Hello)

	b.Start()
}
