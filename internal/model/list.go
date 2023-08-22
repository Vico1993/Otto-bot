package model

// Find if the chat id is already in progress
func isChatIdInProgress(chatId int64) bool {
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
