package redis

import (
	"go.thethings.network/lorawan-stack/v3/pkg/events"
	"go.thethings.network/lorawan-stack/v3/pkg/events/basic"
)

type subscription struct {
	basicSub *basic.Subscription
	patterns []string
}

func (s *subscription) matchPattern(evt events.Event) bool {
	if evt, ok := evt.(*patternEvent); ok {
		for _, pattern := range s.patterns {
			if pattern == evt.pattern {
				return true
			}
		}
	}
	return false
}

func (s *subscription) Match(evt events.Event) bool {
	if s == nil {
		return false
	}
	if !s.matchPattern(evt) {
		return false
	}
	return s.basicSub.Match(evt)
}

func (s *subscription) Notify(evt events.Event) {
	if s == nil {
		return
	}
	s.basicSub.Notify(evt)
}
