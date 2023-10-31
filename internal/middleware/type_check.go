package middleware

import (
	"context"
	"os"
	"strconv"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// Check if we are talking to the bot in a group
func TypeCheck(next bot.HandlerFunc) bot.HandlerFunc {
	return func(ctx context.Context, b *bot.Bot, update *models.Update) {
		chatId := update.Message.Chat.ID

		// Exception if Admin chat id
		if strconv.FormatInt(chatId, 10) == os.Getenv("TELEGRAM_ADMIN_CHAT_ID") {
			next(ctx, b, update)
			return
		}

		// Check if in Group
		if !isGroupChat(&update.Message.Chat) {
			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: chatId,
				Text:   "Sorry can't talk here! We need to be in a group!",
			})
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
