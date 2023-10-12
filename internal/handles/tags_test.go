package handles

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildTagListFromPayloadWithOnlyComma(t *testing.T) {
	res := buildTagListFromPayload("test1,test2,test3")

	assert.Len(t, res, 3, "Should have 3 element in the array")
	assert.Equal(t, []string{"test1", "test2", "test3"}, res, "Should be an slice with test1 test2 and test3")
}

func TestBuildTagListFromPayloadWithSpaces(t *testing.T) {
	res := buildTagListFromPayload("test1 test2 test3")

	assert.Len(t, res, 3, "Should have 3 element in the array")
	assert.Equal(t, []string{"test1", "test2", "test3"}, res, "Should be an slice with test1 test2 and test3")
}

func TestBuildTagListFromPayloadWithSpacesAndComma(t *testing.T) {
	res := buildTagListFromPayload("test1 test2,test3")

	assert.Len(t, res, 3, "Should have 3 element in the array")
	assert.Equal(t, []string{"test1", "test2", "test3"}, res, "Should be an slice with test1 test2 and test3")
}

func TestBuildTagListReply(t *testing.T) {
	res := buildTagListReply([]string{"test1", "test2", "test3"})

	assert.Equal(t, "\n #test1\n #test2\n #test3", res, "Should have a string that contain 3 tags")

}
