package handles

import (
	"strconv"

	tele "gopkg.in/telebot.v3"
)

func List(c tele.Context) error {
	feeds := ottoService.ListFeeds(strconv.FormatInt(c.Chat().ID, 10))
	if len(feeds) == 0 {
		return c.Reply("Thank you for your input, but it appears this chat doesn't watch for any feed. Add some!!!")
	}

	reply := buildListReply(feeds)
	return c.Reply(reply, &tele.SendOptions{
		ParseMode: "Markdown",
	})
}

func buildListReply(feedsUrl []string) string {
	reply := ""
	for k, url := range feedsUrl {
		reply += "\n" + strconv.Itoa(k+1) + ". " + url
	}

	return reply
}
