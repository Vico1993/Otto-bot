package handles

import (
	"context"
	"math/rand"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

var greetings []string = []string{
	"Hello there! How can I assist you today?",
	"Greetings! I'm here and ready to lend a hand.",
	"Hey, it's me, Otto! What can I do for you?",
	"Good day! I'm at your service. How may I help?",
	"Salutations! I'm here to provide any assistance you need.",
	"Hi! Ready to tackle any questions or tasks you have.",
	"Welcome! I'm here to make your day a little easier.",
	"Hello! How may I be of help to you right now?",
	"Hi there! Let's work together to solve any challenges.",
	"Greetings! I'm your AI companion, here to support you.",
}

func Hello(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:           update.Message.Chat.ID,
		Text:             getGreetingsText(),
		ReplyToMessageID: update.Message.ID,
	})
}

// Return a random text
func getGreetingsText() string {
	return greetings[rand.Intn(len(greetings))]
}
