package entrypoint

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWMTSHandler(t *testing.T) {
	h := NewWMTSHandler()
	assert.NotNil(t, h)
}
