package handles

import (
	"fmt"

	"github.com/Vico1993/Otto-bot/internal/model"
	tele "gopkg.in/telebot.v3"
)

type Step struct {
	Run func(ctx tele.Context)
}

type Steps struct {
	Step []Step
}

// Return list of Steps
func NewSteps(step []Step) *Steps {
	return &Steps{
		Step: step,
	}
}

func Init(c tele.Context) error {
	// If not in progress yet
	if !model.IsChatIdInProgress(c.Chat().ID) {
		fmt.Println(c.Message().Sender.ID)

		model.UpsertChatToList(*model.NewChat(
			c.Chat().ID,
			c.Message().Sender.ID,
			1,
			"init",
		))
	}

	chatProgress := model.ListOfChats[c.Chat().ID]

	switch chatProgress.Step {
	case 1:
		initStep1(c)
	case 2:
		initStep2(c)
	case 3:
		initStep3(c)
	default:
		fmt.Printf("No idea why you are here")
	}

	// Update step
	chatProgress.Step += 1

	model.UpsertChatToList(chatProgress)

	return c.Send("initialisation...")
}

func initStep1(c tele.Context) {
	_ = c.Send("Welcome to first step")
}

func initStep2(c tele.Context) {
	_ = c.Send("Welcome to second step")
}

func initStep3(c tele.Context) {
	_ = c.Send("Welcome to last step")
}
