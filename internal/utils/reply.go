package utils

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func Reply(ctx context.Context, b *bot.Bot, update *models.Update, text string, parseMode bool) {
	params := &bot.SendMessageParams{
		ChatID:           update.Message.Chat.ID,
		Text:             text,
		ReplyToMessageID: update.Message.ID,
	}

	if parseMode {
		params.ParseMode = models.ParseModeMarkdown
	}

	b.SendMessage(ctx, params)
}
