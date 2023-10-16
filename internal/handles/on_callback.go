package handles

import (
	"strconv"
	"strings"

	tele "gopkg.in/telebot.v3"
)

func OnCallback(c tele.Context) error {
	b := c.Callback()

	var callback []string
	if strings.Contains(b.Data, "_") {
		callback = strings.Split(b.Data, "_")
	} else {
		callback = []string{b.Data}
	}

	cmd := callback[0]

	if cmd == "disableFeeds" {
		disableFeedsCallBack(strconv.FormatInt(c.Chat().ID, 10), callback[1])
		return c.Send("Feed has been disabled")
	} else if cmd == "deleteTags" {
		deleteTag(strconv.FormatInt(c.Chat().ID, 10), callback[1])
		return c.Send("Tag has been deleted")
	}

	return nil
}
