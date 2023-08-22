package model

type Chat struct {
	ID         int64
	UserId     int64
	Step       int8
	InProgress string
}

// Create a new Chat
func NewChat(chatId int64, userId int64, step int8, command string) *Chat {
	return &Chat{
		ID:         chatId,
		UserId:     userId,
		Step:       step,
		InProgress: command,
	}
}
