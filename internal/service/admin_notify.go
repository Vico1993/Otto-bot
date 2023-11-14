package service

import (
	"context"
	"fmt"
	"os"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

var AdminService *admin

type admin struct {
	ChanelId string
	Bot      *bot.Bot
}

func newAdminService(bot *bot.Bot) *admin {
	return &admin{
		Bot:      bot,
		ChanelId: os.Getenv("TELEGRAM_ADMIN_CHAT_ID"),
	}
}

// Log in Admin Channel
func (a *admin) Log(text string) {
	if a.ChanelId == "" {
		fmt.Println("Admin log not enabled")
		return
	}

	fmt.Println("------")
	fmt.Println(text)
	fmt.Println("------")

	_, err := a.Bot.SendMessage(context.Background(), &bot.SendMessageParams{
		ChatID:    os.Getenv("TELEGRAM_ADMIN_CHAT_ID"),
		Text:      text,
		ParseMode: models.ParseModeHTML,
	})

	if err != nil {
		fmt.Println("Error sending Message in Admin Channel")
		fmt.Println(err.Error())
	}
}

// Initilisation of the Bot Service
func Init(bot *bot.Bot) {
	AdminService = newAdminService(bot)
}
