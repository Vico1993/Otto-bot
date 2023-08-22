package model

// Find if the chat id is already in progress
// TODO: Add the command to be sure it's on progress for the correct command
func IsChatIdInProgress(chatId int64) bool {
	for k := range ListOfChats {
		if k == chatId {
			return true
		}
	}

	return false
}

// Add/update chat to list
func UpsertChatToList(chat Chat) {
	ListOfChats[chat.ID] = chat
}

// Remove a chat id from the map
func DeleteChatFromList(chatId int64) {
	delete(ListOfChats, chatId)
}
