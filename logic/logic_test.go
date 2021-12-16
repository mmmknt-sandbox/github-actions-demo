package logic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEcho(t *testing.T) {
	assert.Equal(t, "test", Echo("test"))
}
