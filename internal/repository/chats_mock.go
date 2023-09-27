package repository

import (
	"github.com/Vico1993/Otto-bot/internal/database"
	"github.com/stretchr/testify/mock"
)

type MocksChatRep struct {
	mock.Mock
}

func (m *MocksChatRep) GetAll() []*database.Chat {
	args := m.Called()
	return args.Get(0).([]*database.Chat)
}

func (m *MocksChatRep) FindByChatId(chatId string) *database.Chat {
	args := m.Called(chatId)
	return args.Get(0).(*database.Chat)
}

func (m *MocksChatRep) PushNewFeed(url string, chatId string) bool {
	args := m.Called(url, chatId)
	return args.Bool(0)
}

func (m *MocksChatRep) Create(chatid string, userid int64, tags []string, feeds []string) *database.Chat {
	args := m.Called(chatid, userid, tags, feeds)
	return args.Get(0).(*database.Chat)
}
