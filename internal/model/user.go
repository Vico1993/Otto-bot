package model

var ChatsInProgress []Chat = []Chat{}

type Chat struct {
	ID         int64
	Step       int8
	InProgress string
}

// Create a new Chat
func newChat(chatId int64, step int8, command string) *Chat {
	return &Chat{
		ID:         chatId,
		Step:       step,
		InProgress: command,
	}
}
