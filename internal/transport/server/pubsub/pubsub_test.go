package pubsub

import (
	"testing"

	"github.com/TimeleapLabs/unchained/internal/consts"
)

func TestSubscribeDirectly(t *testing.T) {
	sub := Subscribe("a.b.c")

	Publish("a.b.c", consts.OpCodeAttestation, []byte("Hello, world!"))

	msg := <-sub
	if string(msg[1:]) != "Hello, world!" {
		t.Fatalf("Received unexpected message: %s", msg)
	}
}

func TestGetTopicsByPrefix(t *testing.T) {
	const (
		first  consts.Channels = "a.b.c"
		second consts.Channels = "a"
		third  consts.Channels = "b.c.d"
	)

	topics = map[consts.Channels][]chan []byte{
		first:  make([]chan []byte, 0),
		second: make([]chan []byte, 0),
		third:  make([]chan []byte, 0),
	}

	trimmedTopics := getTopicsByPrefix(first)

	for topic := range trimmedTopics {
		if topic != first && topic != second {
			t.Fatalf("Unexpected topic: %s", topic)
		}
	}
}

func TestSubscribeWithPrefix(t *testing.T) {
	sub := Subscribe("a")

	Publish("a.b.c", consts.OpCodeAttestation, []byte("Hello, world!"))

	msg := <-sub

	if string(msg[1:]) != "Hello, world!" {
		t.Fatalf("Received unexpected message: %s", msg)
	}
}

func TestPublishWithoutSubscriber(_ *testing.T) {
	Publish("a.b.c", consts.OpCodeAttestation, []byte("Hello, world!"))
}
