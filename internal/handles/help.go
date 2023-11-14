package handles

import (
	"context"
	"fmt"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func Help(ctx context.Context, b *bot.Bot, update *models.Update) {
	sendMessage(ctx, b, update, `I'm designed to selectively send articles from RSS feeds to the channel based on specific user-defined tags.`, false)
	sendMessage(
		ctx,
		b,
		update,
		`It operates exclusively in group chats and supports the following commands:
			/hello: Initiates a greeting message from me.
			/ping: Responds with "pong".
			/start: Signals the start of interaction, enabling the bot to listen to commands.
			/feeds: Lists all feeds the bot is currently listening to in the chat.
			/feedsadd <URL>: Adds a new RSS feed to the bot's listening list.
			/feedsdisabled: Removes a specified feed that the bot is currently listening to in the chat.
			/tags: Lists all tags the bot is searching for in the chat.
			/tagsadd <TAG>: Adds a new tag for the bot to track in fetched articles.
			/tagsdelete: Removes a specified tag from the bot's tracking list.
		`,
		true,
	)
	sendMessage(ctx, b, update, `I operates only in group chats and does not support private conversations.`, false)
	sendMessage(ctx, b, update, `Report any bugs to t.me/Ocivp`, false)
}

// send message
// TODO: DO SOMETHING
func sendMessage(ctx context.Context, b *bot.Bot, update *models.Update, text string, isHtml bool) {
	params := &bot.SendMessageParams{
		ChatID:          update.Message.Chat.ID,
		Text:            text,
		MessageThreadID: update.Message.MessageThreadID,
	}

	if isHtml {
		params.ParseMode = models.ParseModeHTML
	}

	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:          update.Message.Chat.ID,
		Text:            text,
		MessageThreadID: update.Message.MessageThreadID,
	})

	if err != nil {
		fmt.Println("Couldn't post message for help")
		fmt.Println(err.Error())
		fmt.Println("------------------")
	}
}
