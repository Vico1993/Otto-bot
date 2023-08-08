package middleware

import (
	tele "gopkg.in/telebot.v3"
)

// Check if we are talking to the bot in a group
func TypeCheck(next tele.HandlerFunc) tele.HandlerFunc {
	return func(c tele.Context) error {

		// Check if in Group
		if c.Chat().Type != tele.ChatSuperGroup && c.Chat().Type != tele.ChatGroup {
			return c.Send("Sorry can't talk here! We need to be in a group!")
		}

		// TODO: Check if it's super group need Topic rights

		return next(c)
	}
}
