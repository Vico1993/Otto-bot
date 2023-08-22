package handles

import (
	"fmt"

	"github.com/Vico1993/Otto-bot/internal/model"
	tele "gopkg.in/telebot.v3"
)

// Default behavior when receiving a text
func OnText(c tele.Context) error {
	if model.IsChatIdInProgress(c.Chat().ID) {
		chatProgress := model.ListOfChats[c.Chat().ID]

		fmt.Println("-----")
		// TODO: Easiest way to build this switch case ? :thinking:
		switch chatProgress.InProgress {
		case "init":
			fmt.Println("Found chat id working on init")
			_ = Init(c)
		default:
			fmt.Println("Found chat in progress for something weird... got to remove")
			fmt.Println("Got " + chatProgress.InProgress)

			// Clean remove key from map
			model.DeleteChatFromList(c.Chat().ID)
		}
		fmt.Println("-----")
	}

	return nil
}
