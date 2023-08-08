package handles

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReturnRandomMessage(t *testing.T) {
	constGreeting := "Hello Test!"
	// override message
	greetings = []string{constGreeting}

	assert.Equal(t, getGreetingsText(), constGreeting, "There is only 1 element in greetings should return it")
}
