package middleware

import (
	"testing"

	"github.com/stretchr/testify/assert"
	tele "gopkg.in/telebot.v3"
)

func TestIsGroupChatWithPrivateChat(t *testing.T) {
	c := tele.Chat{
		Type: tele.ChatPrivate,
	}

	assert.False(t, isGroupChat(c), "This execption should return false if in Private chat")
}

func TestIsGroupChatWithGroupChat(t *testing.T) {
	c := tele.Chat{
		Type: tele.ChatGroup,
	}

	assert.True(t, isGroupChat(c), "This execption should return true if in Group chat")
}

func TestIsGroupChatWithSuperGroupChat(t *testing.T) {
	c := tele.Chat{
		Type: tele.ChatSuperGroup,
	}

	assert.True(t, isGroupChat(c), "This execption should return true if in Super Group chat")
}
