package tbuffer

import (
	"sync"
	"time"
)

type Buffer[T any] struct {
	mu          sync.Mutex
	maxDelay    time.Duration
	lastFlushTS time.Time
	buffer      chan T
	ticker      *time.Ticker
	flushFn     func([]T)
}

func New[T any](size int, maxDelay time.Duration, flushFn func([]T)) *Buffer[T] {
	buffer := make(chan T, size)
	ticker := time.NewTicker(maxDelay)
	tb := &Buffer[T]{buffer: buffer, ticker: ticker, flushFn: flushFn, lastFlushTS: time.Now(), maxDelay: maxDelay}
	go tb.loop()
	return tb
}

func (tb *Buffer[T]) loop() {
	for range tb.ticker.C {
		tb.mu.Lock()
		if time.Since(tb.lastFlushTS) > tb.maxDelay {
			tb.flush()
		}
		tb.mu.Unlock()
	}
}

func (tb *Buffer[T]) flush() {
	bufLen := len(tb.buffer)
	if bufLen > 0 {
		tmp := make([]T, bufLen)
		for i := 0; i < bufLen; i++ {
			tmp[i] = <-tb.buffer
		}
		tb.flushFn(tmp)
		tb.lastFlushTS = time.Now()
	}
}

func (tb *Buffer[T]) Put(items ...T) {
	tb.mu.Lock()
	defer tb.mu.Unlock()
	for _, i := range items {
		select {
		case tb.buffer <- i:
		default:
			tb.flush()
			tb.buffer <- i
		}
	}
}

func (tb *Buffer[T]) Close() error {
	tb.mu.Lock()
	defer tb.mu.Unlock()
	tb.flush()
	close(tb.buffer)
	tb.ticker.Stop()
	return nil
}
