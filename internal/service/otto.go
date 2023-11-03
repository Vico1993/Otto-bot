package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Vico1993/Otto-bot/internal/utils"
)

type Chat struct {
	Id             string     `json:"id"`
	TelegramChatId string     `json:"TelegramChatId"`
	TelegramUserId string     `json:"TelegramUserId"`
	Tags           string     `json:"Tags"`
	LastTimeParsed *time.Time `json:"LastTimeParsed"`
}

type Feed struct {
	Id  string `json:"id"`
	Url string `json:"url"`
}

type FeedResponse struct {
	Feed Feed `json:"feed"`
}

type ChatFeedLinkResponse struct {
	Added bool `json:"added"`
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
	InitChat(chatId string, userId string, threadId string, tags []string) *Chat
	ListFeeds(chatId string) []Feeds
	ListTags(chatId string) []string
	AddFeed(feedUrl string) *Feed
	LinkFeedToChat(chatId string, feedId string) bool
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
func (s *OttoService) InitChat(chatId string, userId string, threadId string, tags []string) *Chat {
	data := []byte(`{
		"chat_id": "` + chatId + `",
		"user_id": "` + userId + `",
		"thread_id": "` + threadId + `",
		"tags": []
	}`)

	req, err := http.NewRequest(
		http.MethodPost,
		os.Getenv("OTTO_API_URL")+"/chats",
		strings.NewReader(
			string(data),
		),
	)
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		fmt.Println("Error creating the request to initiate chat: " + err.Error())
		return nil
	}

	body, err := s.executeRequest(req)
	if err != nil {
		fmt.Println("Error initiating chat: " + err.Error())
		return nil
	}

	var res ChatResponse
	_ = json.Unmarshal(body, &res)

	return &res.Chat
}

// Retrieve all feeds attached to a chatId
func (s *OttoService) ListFeeds(chatId string) []Feeds {
	req, err := http.NewRequest(
		http.MethodGet,
		os.Getenv("OTTO_API_URL")+"/chats/"+chatId+"/feeds",
		strings.NewReader(
			string([]byte{}),
		),
	)
	if err != nil {
		fmt.Println("Error creating the request to list feeds: " + err.Error())
		return nil
	}

	body, err := s.executeRequest(req)
	if err != nil {
		fmt.Println("Error listing feeds: " + err.Error())
		return nil
	}
	var res ChatFeedsResponse
	_ = json.Unmarshal(body, &res)

	return res.Feeds
}

// Retrieve all tags attached to a chatId
func (s *OttoService) ListTags(chatId string) []string {
	req, err := http.NewRequest(
		http.MethodGet,
		os.Getenv("OTTO_API_URL")+"/chats/"+chatId+"/tags",
		strings.NewReader(
			string([]byte{}),
		),
	)
	if err != nil {
		fmt.Println("Error creating the request to list tags: " + err.Error())
		return nil
	}

	body, err := s.executeRequest(req)
	if err != nil {
		fmt.Println("Error listing tags: " + err.Error())
		return nil
	}
	var res ChatTagsResponse
	_ = json.Unmarshal(body, &res)

	return res.Tags
}

// Add tags to the chats
func (s *OttoService) AddTags(chatId string, tags []string) []string {
	data := []byte(`{
		"tags": ["` + strings.Join(tags, "\",\"") + `"]
	}`)

	req, err := http.NewRequest(
		http.MethodPost,
		os.Getenv("OTTO_API_URL")+"/chats/"+chatId+"/tags",
		strings.NewReader(
			string(data),
		),
	)
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		fmt.Println("Error creating the request to add tag: " + err.Error())
		return nil
	}

	body, err := s.executeRequest(req)
	if err != nil {
		fmt.Println("Error adding new tags: " + err.Error())
		return nil
	}

	var res ChatTagsResponse
	_ = json.Unmarshal(body, &res)

	return res.Tags
}

// Delete one tag from the chat
func (s *OttoService) DeleteTag(chatId string, tag string) bool {
	// Create request
	req, err := http.NewRequest("DELETE", os.Getenv("OTTO_API_URL")+"/chats/"+chatId+"/tags/"+tag, nil)
	if err != nil {
		fmt.Println("Error creating the request to delete the tag: " + err.Error())
		return false
	}

	body, err := s.executeRequest(req)
	if err != nil {
		fmt.Println("Error deleting tag: " + err.Error())
		return false
	}

	var res ChatTagsResponse
	_ = json.Unmarshal(body, &res)

	return !utils.InSlice(tag, res.Tags)
}

func (s *OttoService) DisableFeeds(chatId string, feedId string) bool {
	// Create request
	req, err := http.NewRequest("DELETE", os.Getenv("OTTO_API_URL")+"/chats/"+chatId+"/feeds/"+feedId, nil)
	if err != nil {
		fmt.Println("Error creating the request to disabled the feed in chat: " + err.Error())
		return false
	}

	body, err := s.executeRequest(req)
	if err != nil {
		fmt.Println("Error deleting tag: " + err.Error())
		return false
	}

	var res ChatDeleteFeedResponse
	_ = json.Unmarshal(body, &res)

	return res.Deleted
}

// Add feed to the chat
func (s *OttoService) AddFeed(feedUrl string) *Feed {
	data := []byte(`{
		"url": "` + feedUrl + `"
	}`)

	req, err := http.NewRequest(
		http.MethodPost,
		os.Getenv("OTTO_API_URL")+"/feeds",
		strings.NewReader(
			string(data),
		),
	)
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		fmt.Println("Error creating the request to add feed: " + err.Error())
		return nil
	}

	body, err := s.executeRequest(req)
	if err != nil {
		fmt.Println("Error requesting new feed: " + err.Error())
		return nil
	}

	var res FeedResponse
	_ = json.Unmarshal(body, &res)

	return &res.Feed
}

// link feed to the current chat
func (s *OttoService) LinkFeedToChat(chatId string, feedId string) bool {
	req, err := http.NewRequest(
		http.MethodPost,
		os.Getenv("OTTO_API_URL")+"/chats/"+chatId+"/feeds/"+feedId,
		strings.NewReader(
			string([]byte{}),
		),
	)
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		fmt.Println("Error linking chat and feed: " + err.Error())
		return false
	}

	body, err := s.executeRequest(req)
	if err != nil {
		fmt.Println("Error linking chat and feed: " + err.Error())
		return false
	}

	var res ChatFeedLinkResponse
	_ = json.Unmarshal(body, &res)

	return res.Added
}

func (s *OttoService) executeRequest(req *http.Request) ([]byte, error) {
	// Create client
	client := &http.Client{}

	// Fetch Request
	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making the request: " + err.Error())
		return []byte{}, err
	}
	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)
	if response.StatusCode != http.StatusOK {
		fmt.Println("Api respond with status code: " + strconv.Itoa(response.StatusCode))
		fmt.Println(string(body))
		return []byte{}, err
	}

	return body, nil
}
