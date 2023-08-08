package handles

import (
	tele "gopkg.in/telebot.v3"
)

func Init(c tele.Context) error {
	// c.Chat().Permissions.CanManageChat

	// wantedType := []tele.ChatType{tele.ChatGroup, tele.ChatSuperGroup}

	return c.Send("initialisation...")
}
