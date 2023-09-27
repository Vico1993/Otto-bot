package handles

import (
	"fmt"
	"strconv"

	"github.com/Vico1993/Otto-bot/internal/repository"
	tele "gopkg.in/telebot.v3"
)

func List(c tele.Context) error {
	feedsUrl := retrieveFeedsUrl(strconv.FormatInt(c.Chat().ID, 10))
	if len(feedsUrl) == 0 {
		fmt.Println(c.Chat().ID)
		c.Reply("No feed attached to this chat id")
	}

	reply := buildListReply(feedsUrl)

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

// From Chat id retrieve the list of feed url
func retrieveFeedsUrl(chatId string) []string {
	var feeds = []string{}
	chat := repository.Chat.FindByChatId(chatId)

	if chat == nil {
		return feeds
	}

	for _, feed := range chat.Feeds {
		feeds = append(feeds, feed.Url)
	}

	return feeds
}
