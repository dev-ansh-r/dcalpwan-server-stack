package udp

import (
	"sync"
	"time"
)

type timestamps struct {
	position int
	mu       sync.RWMutex
	items    []time.Time
}

func newTimestamps(count int) *timestamps {
	return &timestamps{
		items: make([]time.Time, count),
	}
}

// Append adds new timestamp and returns oldest.
func (i *timestamps) Append(val time.Time) time.Time {
	i.mu.Lock()
	defer i.mu.Unlock()
	result := i.items[i.position]
	i.items[i.position] = val
	i.position = (i.position + 1) % len(i.items)

	return result
}
