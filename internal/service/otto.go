package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Vico1993/Otto-bot/internal/utils"
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

type Feeds struct {
	Id  string `json:"id"`
	Url string `json:"url"`
}

type ChatFeedsResponse struct {
	Feeds []Feeds `json:"feeds"`
}

type ChatTagsResponse struct {
	Tags []string `json:"tags"`
}

type ChatDeleteFeedResponse struct {
	Deleted bool `json:"deleted"`
}

type IOttoService interface {
	InitChat(chatId string, userId string, tags []string) *Chat
	ListFeeds(chatId string) []Feeds
	ListTags(chatId string) []string
	DeleteTag(chatId string, tag string) bool
	AddTags(chatId string, tags []string) []string
	DisableFeeds(chatId string, feedId string) bool
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
		fmt.Println("Error requesting new chats: " + err.Error())
		return nil
	}

	if response.StatusCode != http.StatusOK {
		fmt.Println("Api respond with status code: " + strconv.Itoa(response.StatusCode))
		return nil
	}

	body, _ := io.ReadAll(response.Body)
	var res ChatResponse
	_ = json.Unmarshal(body, &res)

	return &res.Chat
}

// Retrieve all feeds attached to a chatId
func (s *OttoService) ListFeeds(chatId string) []Feeds {
	response, err := http.Get(
		os.Getenv("OTTO_API_URL") + "/chats/" + chatId + "/feeds",
	)
	if err != nil {
		fmt.Println("Error requesting chats feeds: " + err.Error())
		return nil
	}

	if response.StatusCode != http.StatusOK {
		fmt.Println("Api respond with status code: " + strconv.Itoa(response.StatusCode))
		return nil
	}

	body, _ := io.ReadAll(response.Body)
	var res ChatFeedsResponse
	_ = json.Unmarshal(body, &res)

	return res.Feeds
}

// Retrieve all tags attached to a chatId
func (s *OttoService) ListTags(chatId string) []string {
	response, err := http.Get(
		os.Getenv("OTTO_API_URL") + "/chats/" + chatId + "/tags",
	)

	if err != nil {
		fmt.Println("Error requesting chats feeds: " + err.Error())
		return nil
	}

	if response.StatusCode != http.StatusOK {
		fmt.Println("Api respond with status code: " + strconv.Itoa(response.StatusCode))
		return nil
	}

	body, _ := io.ReadAll(response.Body)
	var res ChatTagsResponse
	_ = json.Unmarshal(body, &res)

	return res.Tags
}

// Add tags to the chats
func (s *OttoService) AddTags(chatId string, tags []string) []string {
	data := []byte(`{
		"tags": ["` + strings.Join(tags, "\",\"") + `"]
	}`)

	response, err := http.Post(
		os.Getenv("OTTO_API_URL")+"/chats/"+chatId+"/tags",
		"application/json",
		bytes.NewBuffer(data),
	)
	if err != nil {
		fmt.Println("Error requesting new chats: " + err.Error())
		return nil
	}

	if response.StatusCode != http.StatusOK {
		fmt.Println("Api respond with status code: " + strconv.Itoa(response.StatusCode))
		return nil
	}

	body, _ := io.ReadAll(response.Body)
	var res ChatTagsResponse
	_ = json.Unmarshal(body, &res)

	return res.Tags
}

// Delete one tag from the chat
func (s *OttoService) DeleteTag(chatId string, tag string) bool {
	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("DELETE", os.Getenv("OTTO_API_URL")+"/chats/"+chatId+"/tags/"+tag, nil)
	if err != nil {
		fmt.Println("Error creating the request to delete the tag: " + err.Error())
		return false
	}

	// Fetch Request
	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making the request deleting the tag: " + err.Error())
		return false
	}
	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)
	var res ChatTagsResponse
	_ = json.Unmarshal(body, &res)

	return !utils.InSlice(tag, res.Tags)
}

func (s *OttoService) DisableFeeds(chatId string, feedId string) bool {
	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("DELETE", os.Getenv("OTTO_API_URL")+"/chats/"+chatId+"/feeds/"+feedId, nil)
	if err != nil {
		fmt.Println("Error creating the request to delete the feed: " + err.Error())
		return false
	}

	// Fetch Request
	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making the request deleting the feed: " + err.Error())
		return false
	}
	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)
	var res ChatDeleteFeedResponse
	_ = json.Unmarshal(body, &res)

	return res.Deleted
}

type MocksOttoService struct {
	mock.Mock
}
