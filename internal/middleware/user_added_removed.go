package middleware

import (
	"context"
	"os"

	"github.com/Vico1993/Otto-bot/internal/service"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// Middleware to check if bot is added in a group
func UserAddedRemoved(next bot.HandlerFunc) bot.HandlerFunc {
	return func(ctx context.Context, b *bot.Bot, update *models.Update) {
		if update.Message != nil {
			for _, user := range update.Message.NewChatMembers {
				if user.Username == os.Getenv("TELEGRAM_BOT_USERNAME") {
					service.AdminService.Log(`🤖 [BOT] Added in a chat group: <b>` + update.Message.Chat.Title + `</b> by <b>` + update.Message.From.FirstName + `</b> t.me/` + update.Message.From.Username + ``)

					break
				}
			}

			userLeft := update.Message.LeftChatMember
			if userLeft != nil && userLeft.Username == os.Getenv("TELEGRAM_BOT_USERNAME") {
				service.AdminService.Log(`🤖 [BOT] Removed from a chat group: <b>` + update.Message.Chat.Title + `</b> by <b>` + update.Message.From.FirstName + `</b> t.me/` + update.Message.From.Username + ``)
			}
		}

		next(ctx, b, update)
	}
}
