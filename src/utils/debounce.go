package utils

import (
	"sync"
	"time"
)

type debounceContext[KeyType comparable, ArgType any] struct {
	sync.Mutex
	timers map[KeyType]*time.Timer
}

func newDebounceContext[KeyType comparable, ArgType any]() *debounceContext[KeyType, ArgType] {
	return &debounceContext[KeyType, ArgType]{
		timers: make(map[KeyType]*time.Timer),
	}
}

func Debounce[KeyType comparable, ArgType any](
	wait time.Duration, function func(ArgType)) func(key KeyType, arg ArgType) {
	debounce := newDebounceContext[KeyType, ArgType]()

	return func(key KeyType, arg ArgType) {
		debounce.Lock()
		defer debounce.Unlock()

		if timer, found := debounce.timers[key]; found {
			timer.Stop()
		}

		debounce.timers[key] = time.AfterFunc(wait, func() {
			debounce.Lock()

			go func() {
				defer debounce.Unlock()
				delete(debounce.timers, key)
				function(arg)
			}()
		})
	}
}
