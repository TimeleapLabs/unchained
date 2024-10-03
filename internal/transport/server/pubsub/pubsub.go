package pubsub

import (
	"strings"
	"sync"

	"github.com/TimeleapLabs/unchained/internal/consts"
)

var topics = make(map[consts.Channels][]chan []byte)
var mu sync.Mutex

func getTopicsByPrefix(topic consts.Channels) map[consts.Channels][]chan []byte {
	keys := make(map[consts.Channels][]chan []byte)
	for key := range topics {
		if strings.HasPrefix(string(topic), string(key)) {
			keys[key] = make([]chan []byte, len(topics[key]))
			copy(keys[key], topics[key])
		}
	}

	return keys
}

func Publish(destinationTopic consts.Channels, operation consts.OpCode, message []byte) {
	mu.Lock()
	defer mu.Unlock()

	allSubTopics := getTopicsByPrefix(destinationTopic)

	for _, subscribers := range allSubTopics {
		for _, ch := range subscribers {
			go func(ch chan []byte) {
				ch <- append([]byte{byte(operation)}, message...)
			}(ch)
		}
	}
}

func Unsubscribe(topic string, ch chan []byte) {
	mu.Lock()
	defer mu.Unlock()

	for key, subscribers := range topics[consts.Channels(topic)] {
		if subscribers == ch {
			topics[consts.Channels(topic)] = append(topics[consts.Channels(topic)][:key], topics[consts.Channels(topic)][key+1:]...)
			break
		}
	}

	close(ch)
}

func Subscribe(topic string) chan []byte {
	mu.Lock()
	defer mu.Unlock()

	ch := make(chan []byte)
	topics[consts.Channels(topic)] = append(topics[consts.Channels(topic)], ch)
	return ch
}
