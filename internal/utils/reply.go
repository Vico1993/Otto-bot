package utils

import (
	"context"
	"fmt"

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

	_, err := b.SendMessage(ctx, params)
	if err != nil {
		fmt.Println("Couldn't post message: " + text)
		fmt.Println(err.Error())
		fmt.Println("------------------")
	}
}
