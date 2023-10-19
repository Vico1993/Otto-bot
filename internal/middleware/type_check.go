package middleware

import (
	"os"
	"strconv"

	tele "gopkg.in/telebot.v3"
)

// Check if we are talking to the bot in a group
func TypeCheck(next tele.HandlerFunc) tele.HandlerFunc {
	return func(c tele.Context) error {

		// Exception if Admin chat id
		if strconv.FormatInt(c.Chat().ID, 10) == os.Getenv("TELEGRAM_ADMIN_CHAT_ID") {
			return next(c)
		}

		// Check if in Group
		if !isGroupChat(*c.Chat()) {
			return c.Send("Sorry can't talk here! We need to be in a group!")
		}

		// TODO: Check if it's super group need Topic rights

		return next(c)
	}
}

// Check if the chat is in a correct group
func isGroupChat(c tele.Chat) bool {
	return c.Type == tele.ChatSuperGroup || c.Type == tele.ChatGroup
}
