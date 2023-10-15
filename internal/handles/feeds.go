package handles

import (
	"strconv"

	"github.com/Vico1993/Otto-bot/internal/service"
	tele "gopkg.in/telebot.v3"
)

func ListFeeds(c tele.Context) error {
	feeds := ottoService.ListFeeds(strconv.FormatInt(c.Chat().ID, 10))
	if len(feeds) == 0 {
		return c.Reply("Thank you for your input, but it appears this chat doesn't watch for any feed. Add some!!!")
	}

	reply := buildListReply(feeds)
	return c.Reply(reply, &tele.SendOptions{
		ParseMode: "Markdown",
	})
}

func buildListReply(list []service.Feeds) string {
	reply := ""
	for k, feed := range list {
		reply += "\n" + strconv.Itoa(k+1) + ". " + feed.Url
	}

	return reply
}

func DisableFeeds(c tele.Context) error {
	feeds := ottoService.ListFeeds(strconv.FormatInt(c.Chat().ID, 10))
	keyboard := make([][]tele.InlineButton, len(feeds))

	for _, feed := range feeds {
		keyboard = append(keyboard, []tele.InlineButton{
			{
				Text: feed.Url,
				Data: "disableFeeds_" + feed.Id,
			},
		})
	}

	return c.Send("Please select the feed you want to disable", &tele.ReplyMarkup{
		RemoveKeyboard: true,
		InlineKeyboard: keyboard,
	})
}

func disableFeedsCallBack(chatId string, feedId string) {
	service.NewOttoService().DisableFeeds(chatId, feedId)
}
