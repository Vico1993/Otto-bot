package service

import (
	"math/rand"
)

var templates []string = []string{
	`Oh dear, it seems an error has popped up. Not to worry, I'm here to help. Please try your request again, and if the issue persists, feel free to let me know.`,
	`Oops! Something unexpected happened. Give it another shot, and I'll be here to assist. If the issue continues, just drop me a message.`,
	`Mistakes happen, even to us bots. Don't be discouraged. Please try your action one more time, and if you run into the same problem, I'm just a message away.`,
	`An error has been detected, but we're on top of it. Feel free to retry, and I'll be standing by to lend a hand if you need it.`,
	`I'm sorry, but something's not quite right at the moment. Give it another go, and if things don't improve, don't hesitate to reach out. I'm here to make things right.`,
}

// Return a friendly error message
func ReturnError() string {
	return templates[rand.Intn(len(templates))]
}
