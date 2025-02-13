package pubsub

import (
	"context"
	"strings"
	"sync"

	"github.com/TimeleapLabs/timeleap/internal/consts"
	"github.com/TimeleapLabs/timeleap/internal/model"
	"github.com/TimeleapLabs/timeleap/internal/transport/server/packet"
	"github.com/TimeleapLabs/timeleap/internal/transport/server/websocket/queue"
)

type Subscriber struct {
	Writer            *queue.WebSocketWriter
	Channel           chan []byte
	Context           *context.Context
	ConnectionContext *context.Context
	Unsubscribe       context.CancelFunc
}

var topics = make(map[string][]Subscriber)
var mu sync.Mutex

func getTopicsByPrefix(topic string) map[string][]Subscriber {
	keys := make(map[string][]Subscriber)
	for key := range topics {
		if strings.HasPrefix(topic, key) {
			keys[key] = make([]Subscriber, len(topics[key]))
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
		for _, sub := range subscribers {
			go writeToChannel(sub.Channel, payload)
		}
	}
}

func Unsubscribe(topic string, writer *queue.WebSocketWriter) {
	mu.Lock()
	defer mu.Unlock()

	for key, sub := range topics[topic] {
		if sub.Writer == writer {
			topics[topic] = append(topics[topic][:key], topics[topic][key+1:]...)
			sub.Unsubscribe()
			close(sub.Channel)
			break
		}
	}
}

func IsSubscribed(topic string, writer *queue.WebSocketWriter) bool {
	for _, sub := range topics[topic] {
		if sub.Writer == writer {
			return true
		}
	}

	return false
}

func Subscribe(ctx context.Context, writer *queue.WebSocketWriter, topic string) (context.Context, Subscriber) {
	mu.Lock()
	defer mu.Unlock()

	subCtx, cancel := context.WithCancel(ctx)

	subscriber := Subscriber{
		Writer:      writer,
		Channel:     make(chan []byte),
		Unsubscribe: cancel,
	}

	topics[topic] = append(topics[topic], subscriber)
	return subCtx, subscriber
}
