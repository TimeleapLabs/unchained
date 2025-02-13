package pubsub

import (
	"strings"
	"sync"

	"github.com/TimeleapLabs/timeleap/internal/consts"
	"github.com/TimeleapLabs/timeleap/internal/model"
	"github.com/TimeleapLabs/timeleap/internal/transport/server/packet"
)

var topics = make(map[string][]chan []byte)
var mu sync.Mutex

func getTopicsByPrefix(topic string) map[string][]chan []byte {
	keys := make(map[string][]chan []byte)
	for key := range topics {
		if strings.HasPrefix(topic, key) {
			keys[key] = make([]chan []byte, len(topics[key]))
			copy(keys[key], topics[key])
		}
	}

	return keys
}

func writeToChannel(ch chan []byte, message []byte) {
	ch <- message
}

func PublishMessage(message []byte) {
	msg := new(model.Message).FromBytes(message[1:])
	Publish(msg.Topic, message)
}

func Publish(destinationTopic string, message []byte) {
	mu.Lock()
	defer mu.Unlock()

	allSubTopics := getTopicsByPrefix(destinationTopic)
	broadcast := append([]byte{byte(consts.OpCodeBroadcast)}, message...)
	payload := packet.New(broadcast).Sia().Bytes()

	for _, subscribers := range allSubTopics {
		for _, ch := range subscribers {
			go writeToChannel(ch, payload)
		}
	}
}

func Unsubscribe(topic string, ch chan []byte) {
	mu.Lock()
	defer mu.Unlock()

	for key, subscribers := range topics[topic] {
		if subscribers == ch {
			topics[topic] = append(topics[topic][:key], topics[topic][key+1:]...)
			break
		}
	}

	close(ch)
}

func Subscribe(topic string) chan []byte {
	mu.Lock()
	defer mu.Unlock()

	ch := make(chan []byte)
	topics[topic] = append(topics[topic], ch)
	return ch
}
