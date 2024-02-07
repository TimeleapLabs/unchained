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

	context := newDebounceContext[KeyType, ArgType]()

	return func(key KeyType, arg ArgType) {
		context.Lock()
		defer context.Unlock()

		if timer, found := context.timers[key]; found {
			timer.Stop()
		}

		context.timers[key] = time.AfterFunc(wait, func() {
			context.Lock()
			defer context.Unlock()

			delete(context.timers, key)
			function(arg)
		})
	}
}
