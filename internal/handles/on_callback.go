package handles

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func OnCallback(ctx context.Context, b *bot.Bot, update *models.Update) {
	var callback []string
	if strings.Contains(update.CallbackQuery.Data, "_") {
		callback = strings.Split(update.CallbackQuery.Data, "_")
	} else {
		callback = []string{update.CallbackQuery.Data}
	}

	cmd := callback[0]
	chatId := strconv.FormatInt(update.CallbackQuery.Message.Chat.ID, 10)

	var text string
	if cmd == "disableFeeds" {
		disableFeedsCallBack(chatId, callback[1])
		text = "Feed has been disabled"
	} else if cmd == "deleteTags" {
		deleteTag(chatId, callback[1])
		text = "Tag has been deleted"
	}

	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: chatId,
		Text:   text,
	})
	if err != nil {
		fmt.Println("Couldn't OnCallback response message: " + text)
		fmt.Println(err.Error())
		fmt.Println("------------------")
	}
}
