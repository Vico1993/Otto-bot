package handles

import (
	"strconv"

	"github.com/Vico1993/Otto-bot/internal/service"
	tele "gopkg.in/telebot.v3"
)

// Manage the `/start`
func Start(c tele.Context) error {
	chat := ottoService.InitChat(strconv.FormatInt(c.Chat().ID, 10), c.Chat().Username, []string{"created"})
	if chat == nil {
		_ = c.Reply(service.ReturnError())

		return nil
	}

	chats[chat.Id] = chat
	_ = c.Reply("Hello " + c.Chat().Username)

	return nil
}
