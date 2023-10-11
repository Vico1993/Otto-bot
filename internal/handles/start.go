package handles

import (
	"strconv"

	"github.com/Vico1993/Otto-bot/internal/service"
	tele "gopkg.in/telebot.v3"
)

var ottoService = service.NewOttoService()

func Start(c tele.Context) error {
	chat := ottoService.InitChat(strconv.FormatInt(c.Chat().ID, 10), c.Chat().Username, []string{"created"})
	if chat == nil {
		_ = c.Reply("Something happns")

		return nil
	}

	chats[chat.Id] = chat
	_ = c.Reply("You are initiated, thank you")

	return nil
}
