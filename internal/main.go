package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Vico1993/Otto-bot/internal/handles"
	"github.com/Vico1993/Otto-bot/internal/middleware"
	"github.com/Vico1993/Otto-bot/internal/utils"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/subosito/gotenv"
)

func main() {
	// load .env file if any otherwise use env set
	_ = gotenv.Load()

	opts := []bot.Option{
		// Middleware
		bot.WithMiddlewares(middleware.TypeCheck),
	}

	b, err := bot.New(os.Getenv("TELEGRAM_BOT_TOKEN"), opts...)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Notify update if chat present
	if os.Getenv("TELEGRAM_ADMIN_CHAT_ID") != "" {
		_, err := b.SendMessage(context.TODO(), &bot.SendMessageParams{
			ChatID:    os.Getenv("TELEGRAM_ADMIN_CHAT_ID"),
			Text:      `🤖 🤖 [BOT] Version: <b>` + utils.RetrieveVersion() + `</b> Succesfully deployed . 🤖 🤖`,
			ParseMode: models.ParseModeHTML,
		})
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	// Handles
	b.RegisterHandler(bot.HandlerTypeMessageText, "/hello", bot.MatchTypeExact, handles.Hello)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/ping", bot.MatchTypeExact, handles.Ping)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, handles.Start)

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
