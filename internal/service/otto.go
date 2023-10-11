package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/stretchr/testify/mock"
)

type Chat struct {
	Id             string     `json:"id"`
	TelegramChatId string     `json:"TelegramChatId"`
	TelegramUserId string     `json:"TelegramUserId"`
	Tags           string     `json:"Tags"`
	LastTimeParsed *time.Time `json:"LastTimeParsed"`
}

type ChatResponse struct {
	Chat Chat `json:"chat"`
}

type IOttoService interface {
	InitChat(chatId string, userId string, tags []string) *Chat
}

type OttoService struct {
	baseUrl string
}

// Initiate an Otto service
func NewOttoService() IOttoService {
	return &OttoService{
		baseUrl: os.Getenv("OTTO_API_URL"),
	}
}

// Create the chat in Otto DB
func (s *OttoService) InitChat(chatId string, userId string, tags []string) *Chat {
	data := []byte(`{
		"chat_id": "` + chatId + `",
		"user_id": "` + userId + `",
		"tags": []
	}`)

	response, err := http.Post(
		os.Getenv("OTTO_API_URL")+"/chats",
		"application/json",
		bytes.NewBuffer(data),
	)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	body, _ := io.ReadAll(response.Body)
	var res ChatResponse
	_ = json.Unmarshal(body, &res)

	return &res.Chat
}

type MocksOttoService struct {
	mock.Mock
}
