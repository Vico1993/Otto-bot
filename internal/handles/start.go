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
	chat := ottoService.InitChat(
		strconv.FormatInt(update.Message.Chat.ID, 10),
		strconv.FormatInt(update.Message.From.ID, 10),
		strconv.Itoa(update.Message.MessageThreadID),
		[]string{},
	)
	if chat == nil {
		utils.Reply(ctx, b, update, service.ReturnError(), false)
		return
	}

	chats[chat.Id] = chat
	utils.Reply(ctx, b, update, "Hello "+update.Message.Chat.Username, false)
}
