package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Vico1993/Otto-bot/internal/database"
	"github.com/Vico1993/Otto-bot/internal/handles"
	"github.com/Vico1993/Otto-bot/internal/middleware"
	"github.com/Vico1993/Otto-bot/internal/model"
	"github.com/Vico1993/Otto-bot/internal/repository"
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

	// Model initialisation
	model.Init()

	// Load the database
	database.Init()

	// Load repositories
	repository.Init()

	// Middleware
	b.Use(middleware.TypeCheck)

	// Handles
	b.Handle("/hello", handles.Hello)
	b.Handle("/init", handles.Init)
	b.Handle("/ping", handles.Ping)
	b.Handle("/list", handles.List)

	b.Handle(tele.OnText, handles.OnText)

	fmt.Println("Ready to Go!...")

	b.Start()
}
