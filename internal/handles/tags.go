package handles

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Vico1993/Otto-bot/internal/service"
	"github.com/Vico1993/Otto-bot/internal/utils"
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
	tagsToDelete := buildTagListFromPayload(c.Message().Payload)

	deleted := []string{}
	for _, tag := range tagsToDelete {
		tag := tag
		done := ottoService.DeleteTag(strconv.FormatInt(c.Chat().ID, 10), strings.TrimSpace(tag))

		if done {
			deleted = append(deleted, tag)
		}
	}

	if len(deleted) != len(tagsToDelete) {
		fmt.Println("Some tags where not deleted")
		return c.Reply(service.ReturnError())
	}

	return c.Reply("Great news! We've successfully deleted " + strconv.Itoa(len(deleted)) + " tags as requested")
}

func TagsAdd(c tele.Context) error {
	tagsToAdd := buildTagListFromPayload(c.Message().Payload)

	tags := ottoService.AddTags(strconv.FormatInt(c.Chat().ID, 10), tagsToAdd)
	if tags == nil {
		return c.Reply(service.ReturnError())
	}

	return c.Reply("Great news! We've successfully added " + strconv.Itoa(len(tagsToAdd)) + " tags as requested")
}

func buildTagListFromPayload(payload string) []string {
	tags := []string{}

	if strings.Contains(payload, " ") {
		tags = strings.Split(payload, " ")
	} else {
		tags = append(tags, payload)
	}

	for k, tag := range tags {
		if strings.Contains(tag, ",") {
			tags = utils.RemoveFromSlice(tags, k)

			tags = append(tags, strings.Split(tag, ",")...)
		}
	}

	return tags
}

func buildTagListReply(list []string) string {
	reply := ""
	for _, tag := range list {
		reply += "\n #" + tag
	}

	return reply
}
