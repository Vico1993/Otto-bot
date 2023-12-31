package handles

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/Vico1993/Otto-bot/internal/utils"
	"github.com/Vico1993/Otto-client/otto"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func ListFeeds(ctx context.Context, b *bot.Bot, update *models.Update) {
	feeds := OttoClient.Feed.List(strconv.FormatInt(update.Message.Chat.ID, 10), strconv.Itoa(update.Message.MessageThreadID))
	if len(feeds) == 0 {
		utils.Reply(ctx, b, update, "Thank you for your input, but it appears this chat doesn't watch for any feed. Add some!!!", false)
		return
	}

	reply := buildListReply(feeds)
	utils.Reply(ctx, b, update, reply, false)
}

func buildListReply(list []otto.Feed) string {
	reply := "📚 Here's the lineup of feeds this chat is subscribed to: \n"
	for _, feed := range list {
		reply += "\n " + feed.Url
	}

	return reply
}

func DisableFeeds(ctx context.Context, b *bot.Bot, update *models.Update) {
	feeds := OttoClient.Feed.List(strconv.FormatInt(update.Message.Chat.ID, 10), strconv.Itoa(update.Message.MessageThreadID))

	var keyboard [][]models.InlineKeyboardButton
	for _, feed := range feeds {
		keyboard = append(keyboard, []models.InlineKeyboardButton{
			{
				Text:         feed.Url,
				CallbackData: "disableFeeds_" + feed.Id + "_callback",
			},
		})
	}

	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:           update.Message.Chat.ID,
		Text:             "Please select the feed you want to disable",
		ReplyToMessageID: update.Message.ID,
		MessageThreadID:  update.Message.MessageThreadID,
		ReplyMarkup: &models.InlineKeyboardMarkup{
			InlineKeyboard: keyboard,
		},
	})

	if err != nil {
		fmt.Println(err.Error())
	}
}

func disableFeedsCallBack(chatId string, threadId string, feedId string) {
	OttoClient.Feed.UnLink(chatId, threadId, feedId)
}

func AddFeeds(ctx context.Context, b *bot.Bot, update *models.Update) {
	var data []string
	text := strings.Trim(update.Message.Text, " ")
	if text != "" {
		data = strings.Split(text, " ")
	}

	if len(data) == 0 {
		utils.Reply(ctx, b, update, "Sorry, something happens couldn't add your feeds", false)
		return
	}

	payload := data[1]

	if !isValidUrl(payload) {
		utils.Reply(ctx, b, update, "Sorry, it seems your url is not well formated. Please give me a valid RSS url", false)
		return
	}

	fmt.Println(payload)

	feed := OttoClient.Feed.Create(payload)
	if feed == nil {
		utils.Reply(ctx, b, update, "Sorry, something happens couldn't add your feeds", false)
		return
	}

	added := OttoClient.Feed.Link(strconv.FormatInt(update.Message.Chat.ID, 10), strconv.Itoa(update.Message.MessageThreadID), feed.Id)
	if !added {
		utils.Reply(ctx, b, update, "Sorry, something happens couldn't add your feeds", false)
		return
	}

	utils.Reply(ctx, b, update, "All good! I will watch this feed closely!", false)
}

// isValidUrl tests a string to determine if it is a well-structured url or not.
func isValidUrl(str string) bool {
	_, err := url.ParseRequestURI(str)
	if err != nil {
		return false
	}

	u, err := url.Parse(str)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	return true
}
