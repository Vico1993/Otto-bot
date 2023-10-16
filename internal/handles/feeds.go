package handles

import (
	"fmt"
	"net/url"
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

func AddFeeds(c tele.Context) error {
	payload := c.Message().Payload
	if !isValidUrl(payload) {
		return c.Reply("Sorry, it seems your url is not well formated. Please give me a valid RSS url")
	}

	fmt.Println(payload)

	feed := ottoService.AddFeed(payload)
	if feed == nil {
		return c.Reply("Sorry, something happens couldn't add your feeds")
	}

	added := ottoService.LinkFeedToChat(strconv.FormatInt(c.Chat().ID, 10), feed.Id)
	if !added {
		return c.Reply("Sorry, something happens couldn't add your feeds")
	}

	return c.Reply("All good! I will watch this feed closely!")
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
