package tbuffer

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFlushOnFull(t *testing.T) {
	var bufferContents []string
	tb := New(5, 30*time.Second, func(items []string) { bufferContents = append(bufferContents, items...) })
	for i := 0; i < 8; i++ {
		tb.Put(fmt.Sprintf("item%d", i))
		if i < 5 {
			assert.Equal(t, 0, len(bufferContents))
		}
	}
	assert.Equal(t, 5, len(bufferContents))
}

func TestFlushOnTimer(t *testing.T) {
	var bufferContents []string
	tb := New(5, 1*time.Second, func(items []string) { bufferContents = append(bufferContents, items...) })
	for i := 0; i < 3; i++ {
		tb.Put(fmt.Sprintf("item%d", i))
		assert.Equal(t, 0, len(bufferContents))
	}

	time.Sleep(1200 * time.Millisecond)
	assert.Equal(t, 3, len(bufferContents))
	tb.Close()
	assert.Equal(t, 3, len(bufferContents))
}

func TestFlushOnClose(t *testing.T) {
	var bufferContents []string
	tb := New(5, 30*time.Second, func(items []string) { bufferContents = append(bufferContents, items...) })
	for i := 0; i < 4; i++ {
		tb.Put(fmt.Sprintf("item%d", i))
	}

	tb.Close()
	assert.Equal(t, 4, len(bufferContents))
}
