package pubsub

import (
	"strings"
	"sync"

	"github.com/TimeleapLabs/unchained/internal/consts"
)

// topics is a map of topics to a slice of channels that are subscribed to that topic.
var topics = make(map[consts.Channels][]chan []byte)
var mu sync.Mutex

// getTopicsByPrefix returns a map of topics that have the given prefix.
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

// Publish sends a message to all subscribers of the given topic.
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

// Subscribe creates a new channel and appends it to the list of subscribers for the given topic.
func Subscribe(topic string) chan []byte {
	mu.Lock()
	defer mu.Unlock()

	ch := make(chan []byte)
	topics[consts.Channels(topic)] = append(topics[consts.Channels(topic)], ch)
	return ch
}
