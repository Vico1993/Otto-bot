package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Vico1993/Otto-bot/internal/handles"
	"github.com/Vico1993/Otto-bot/internal/middleware"
	"github.com/Vico1993/Otto-bot/internal/service"
	"github.com/Vico1993/Otto-bot/internal/utils"
	"github.com/go-telegram/bot"
	"github.com/subosito/gotenv"
)

func main() {
	// load .env file if any otherwise use env set
	_ = gotenv.Load()

	opts := []bot.Option{
		// Middleware
		bot.WithMiddlewares(middleware.TypeCheck),
		bot.WithMiddlewares(middleware.UserAddedRemoved),
	}

	b, err := bot.New(os.Getenv("TELEGRAM_BOT_TOKEN"), opts...)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Initilisation of services
	service.Init(b)

	// Notify update if chat present
	service.AdminService.Log(`🤖 🤖 [BOT] Version: <b>` + utils.RetrieveVersion() + `</b> Succesfully deployed . 🤖 🤖`)

	// Initialise handles
	handles.Init()

	// Handles
	b.RegisterHandler(bot.HandlerTypeMessageText, "/hello", bot.MatchTypeExact, handles.Hello)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/ping", bot.MatchTypeExact, handles.Ping)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, handles.Start)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/help", bot.MatchTypeExact, handles.Help)

	// Feeds
	b.RegisterHandler(bot.HandlerTypeMessageText, "/feeds", bot.MatchTypeExact, handles.ListFeeds)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/feedsadd", bot.MatchTypeContains, handles.AddFeeds)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/feedsdisabled", bot.MatchTypeExact, handles.DisableFeeds)

	// Tags
	b.RegisterHandler(bot.HandlerTypeMessageText, "/tags", bot.MatchTypeExact, handles.TagsList)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/tagsadd", bot.MatchTypeContains, handles.TagsAdd)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/tagsdelete", bot.MatchTypeExact, handles.TagsDelete)

	// Callback
	b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "_callback", bot.MatchTypeContains, handles.OnCallback)

	fmt.Println("Ready to Go!...")
	b.Start(context.TODO())
}
