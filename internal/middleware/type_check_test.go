package middleware

import (
	"testing"

	"github.com/go-telegram/bot/models"
	"github.com/stretchr/testify/assert"
)

func TestIsGroupChatWithPrivateChat(t *testing.T) {
	c := &models.Chat{
		Type: "private",
	}

	assert.False(t, isGroupChat(c), "This execption should return false if in Private chat")
}

func TestIsGroupChatWithGroupChat(t *testing.T) {
	c := &models.Chat{
		Type: "group",
	}

	assert.True(t, isGroupChat(c), "This execption should return true if in Group chat")
}

func TestIsGroupChatWithSuperGroupChat(t *testing.T) {
	c := &models.Chat{
		Type: "supergroup",
	}

	assert.True(t, isGroupChat(c), "This execption should return true if in Super Group chat")
}
