package handles

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/Vico1993/Otto-bot/internal/service"
	"github.com/Vico1993/Otto-bot/internal/utils"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func TagsList(ctx context.Context, b *bot.Bot, update *models.Update) {
	tags := ottoService.ListTags(strconv.FormatInt(update.Message.Chat.ID, 10), strconv.Itoa(update.Message.MessageThreadID))
	if len(tags) == 0 {
		utils.Reply(ctx, b, update, "Thank you for your input, but it appears this chat doesn't watch for any tags. Add some!!!", false)
		return
	}

	reply := buildTagListReply(tags)
	utils.Reply(ctx, b, update, reply, false)
}

func TagsDelete(ctx context.Context, b *bot.Bot, update *models.Update) {
	tags := ottoService.ListTags(strconv.FormatInt(update.Message.Chat.ID, 10), strconv.Itoa(update.Message.MessageThreadID))
	if len(tags) == 0 {
		utils.Reply(ctx, b, update, "Thank you for your input, but it appears this chat doesn't watch for any tags. Add some!!!", false)
		return
	}

	var keyboard [][]models.InlineKeyboardButton
	for _, tag := range tags {
		keyboard = append(keyboard, []models.InlineKeyboardButton{
			{
				Text:         tag,
				CallbackData: "deleteTags_" + tag + "_callback",
			},
		})
	}

	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:           update.Message.Chat.ID,
		Text:             "Please select the tag you want to delete",
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

func deleteTag(chatId string, threadId string, tag string) bool {
	return ottoService.DeleteTag(chatId, threadId, tag)
}

func TagsAdd(ctx context.Context, b *bot.Bot, update *models.Update) {
	var data []string
	text := strings.Trim(update.Message.Text, " ")
	if text != "" {
		data = strings.Split(text, " ")
	}

	if len(data) == 1 || len(data) == 0 {
		utils.Reply(ctx, b, update, "Sorry, something happens couldn't add your tags", false)
		return
	}

	payload := data[1]
	tagsToAdd := buildTagListFromPayload(payload)

	tags := ottoService.AddTags(strconv.FormatInt(update.Message.Chat.ID, 10), strconv.Itoa(update.Message.MessageThreadID), tagsToAdd)
	if tags == nil {
		utils.Reply(ctx, b, update, service.ReturnError(), false)
		return
	}

	utils.Reply(ctx, b, update, "Great news! We've successfully added "+strconv.Itoa(len(tagsToAdd))+" tags as requested", false)
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
	reply := "🔍 Here's what I'm on the lookout for in this chat \n"
	for _, tag := range list {
		reply += "\n #" + tag
	}

	return reply
}
