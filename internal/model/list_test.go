package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsChatIdInProgressWithChatId(t *testing.T) {
	// Reset list
	ListOfChats = make(map[int64]Chat)

	var chatIdToTest int64 = 12341
	ListOfChats[chatIdToTest] = *NewChat(chatIdToTest, 1, 2, "")

	assert.True(t, IsChatIdInProgress(chatIdToTest), "Chat id should be present and valid")
}

func TestIsChatIdInProgressWithoutChatId(t *testing.T) {
	// Reset list
	ListOfChats = make(map[int64]Chat)

	var chatIdToTest int64 = 12341
	ListOfChats[chatIdToTest] = *NewChat(chatIdToTest, 1, 2, "")

	assert.False(t, IsChatIdInProgress(12342), "Chat id requested should not be present")

}

func TestUpsertChatToListAdd(t *testing.T) {
	// Reset list
	ListOfChats = make(map[int64]Chat)

	assert.Len(t, ListOfChats, 0, "List of chats should be empty")

	var chatIdToTest int64 = 12341
	UpsertChatToList(*NewChat(chatIdToTest, 1, 2, ""))

	assert.Len(t, ListOfChats, 1, "List of chats should contain 1 chat now")
}

func TestUpsertChatToListEdit(t *testing.T) {
	// Reset list
	ListOfChats = make(map[int64]Chat)
	var chatIdToTest int64 = 12341
	ListOfChats[chatIdToTest] = *NewChat(chatIdToTest, 1, 2, "")

	assert.Len(t, ListOfChats, 1, "List of chats should contain 1 chat")

	UpsertChatToList(*NewChat(chatIdToTest, 1, 2, "command"))

	assert.Len(t, ListOfChats, 1, "List of chats should contain 1 chat now")
	assert.Equal(t, ListOfChats,
		map[int64]Chat{
			chatIdToTest: *NewChat(chatIdToTest, 1, 2, "command"),
		},
		"Chat should equal the result",
	)
}

func TestDeleteChatFromList(t *testing.T) {
	// Reset list
	ListOfChats = make(map[int64]Chat)
	var chatIdToTest int64 = 12341
	ListOfChats[chatIdToTest] = *NewChat(chatIdToTest, 1, 2, "")

	assert.Len(t, ListOfChats, 1, "List of chats should contain 1 chat")

	DeleteChatFromList(chatIdToTest)

	assert.Len(t, ListOfChats, 0, "List of chats should be empty now")
}
