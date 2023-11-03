package middleware

import (
	"context"
	"os"
	"strconv"

	"github.com/Vico1993/Otto-bot/internal/utils"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// Check if we are talking to the bot in a group
func TypeCheck(next bot.HandlerFunc) bot.HandlerFunc {
	return func(ctx context.Context, b *bot.Bot, update *models.Update) {
		if update.Message == nil {
			next(ctx, b, update)
			return
		}

		chatId := update.Message.Chat.ID

		// Exception if Admin chat id
		if strconv.FormatInt(chatId, 10) == os.Getenv("TELEGRAM_ADMIN_CHAT_ID") {
			next(ctx, b, update)
			return
		}

		// Check if in Group
		if !isGroupChat(&update.Message.Chat) {
			utils.Reply(ctx, b, update, "Sorry can't talk here! We need to be in a group!", false)
			return
		}

		// TODO: Check if it's super group need Topic rights
		next(ctx, b, update)
	}
}

// Check if the chat is in a correct group
func isGroupChat(c *models.Chat) bool {
	return c.Type == "supergroup" || c.Type == "group"
}
