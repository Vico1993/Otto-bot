package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Vico1993/Otto-bot/internal/handles"
	"github.com/Vico1993/Otto-bot/internal/middleware"
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

	// Middleware
	b.Use(middleware.TypeCheck)

	// Handles
	b.Handle("/hello", handles.Hello)
	b.Handle("/ping", handles.Ping)
	b.Handle("/start", handles.Start)

	// Feeds
	b.Handle("/feeds", handles.ListFeeds)
	b.Handle("/feedsadd", handles.AddFeeds)
	b.Handle("/feedsdisabled", handles.DisableFeeds)

	// Tags
	b.Handle("/tags", handles.TagsList)
	b.Handle("/tagsadd", handles.TagsAdd)
	b.Handle("/tagsdelete", handles.TagsDelete)

	// onCallback
	b.Handle(tele.OnCallback, handles.OnCallback)

	fmt.Println("Ready to Go!...")

	b.Start()
}
