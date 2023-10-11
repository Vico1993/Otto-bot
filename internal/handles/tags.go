package handles

import (
	"strconv"

	tele "gopkg.in/telebot.v3"
)

func TagsList(c tele.Context) error {
	tags := ottoService.ListTags(strconv.FormatInt(c.Chat().ID, 10))
	if len(tags) == 0 {
		return c.Reply("Thank you for your input, but it appears this chat doesn't watch for any tags. Add some!!!")
	}

	reply := buildTagListReply(tags)
	return c.Reply(reply, &tele.SendOptions{
		ParseMode: "Markdown",
	})
}

func TagsDelete(c tele.Context) error {
	// ottoService.DeleteTag()

	return nil
}

func buildTagListReply(list []string) string {
	reply := ""
	for _, tag := range list {
		reply += "\n #" + tag
	}

	return reply
}
