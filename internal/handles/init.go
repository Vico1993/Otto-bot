package handles

import (
	tele "gopkg.in/telebot.v3"
)

func Init(c tele.Context) error {
	// c.Chat().ID

	return c.Send("initialisation...")
}
