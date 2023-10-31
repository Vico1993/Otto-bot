package handles

import (
	"context"

	"github.com/Vico1993/Otto-bot/internal/utils"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func Ping(ctx context.Context, b *bot.Bot, update *models.Update) {
	utils.Reply(ctx, b, update, "Pong", false)
}
