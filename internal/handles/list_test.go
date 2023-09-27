package handles

import (
	"testing"

	"github.com/Vico1993/Otto-bot/internal/database"
	"github.com/Vico1993/Otto-bot/internal/repository"
	"github.com/stretchr/testify/assert"
)

// func TestBuildListReply() {}

func TestRetrieveFeedsUrlWithNoData(t *testing.T) {
	chatId := "1234"

	chatRepositoryMock := new(repository.MocksChatRep)
	repository.Chat = chatRepositoryMock

	chatRepositoryMock.On("FindByChatId", chatId).Return((*database.Chat)(nil), nil)

	res := retrieveFeedsUrl(chatId)

	assert.Len(t, res, 0, "Result of retrieveFeedsUrl should be 0 as repository return no informations")
}

func TestRetrieveFeedsUrlData(t *testing.T) {
	feed := database.NewFeed("https://google.com")
	chat := database.NewChat("1234", 123, []database.Feed{
		*feed,
	})

	chatRepositoryMock := new(repository.MocksChatRep)
	repository.Chat = chatRepositoryMock

	chatRepositoryMock.On("FindByChatId", chat.ChatId).Return(chat)

	res := retrieveFeedsUrl(chat.ChatId)

	assert.Len(t, res, 1, "Result of retrieveFeedsUrl should be 1 as repository return 1 chat")
	assert.Equal(t, res, []string{"https://google.com"})
}
