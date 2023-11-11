package handles

import (
	"context"
	"strconv"

	"github.com/Vico1993/Otto-bot/internal/service"
	"github.com/Vico1993/Otto-bot/internal/utils"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// Manage the `/start`
func Start(ctx context.Context, b *bot.Bot, update *models.Update) {
	threadId := strconv.Itoa(update.Message.MessageThreadID)
	if threadId == "0" {
		threadId = ""
	}

	chat := OttoClient.Chat.Create(
		strconv.FormatInt(update.Message.Chat.ID, 10),
		strconv.FormatInt(update.Message.From.ID, 10),
		threadId,
		[]string{},
	)
	if chat == nil {
		utils.Reply(ctx, b, update, service.ReturnError(), false)
		return
	}

	utils.Reply(ctx, b, update, "Hello "+update.Message.Chat.Username, false)
}
