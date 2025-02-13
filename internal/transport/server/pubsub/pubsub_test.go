package pubsub

import (
	"testing"

	"github.com/TimeleapLabs/timeleap/internal/config"
	"github.com/TimeleapLabs/timeleap/internal/consts"
	"github.com/TimeleapLabs/timeleap/internal/crypto"
	"github.com/TimeleapLabs/timeleap/internal/utils"
)

func TestMain(m *testing.M) {
	utils.SetupLogger("info")
	config.App.System.AllowGenerateSecrets = true

	crypto.InitMachineIdentity(
		crypto.WithEd25519Identity(),
	)

	m.Run()
}

func TestSubscribeDirectly(t *testing.T) {
	sub := Subscribe("a.b.c")

	Publish("a.b.c", []byte("Hello, world!"))

	msg := <-sub
	opcode := msg[0]
	if opcode != byte(consts.OpCodeBroadcast) {
		t.Fatalf("Received unexpected opcode: %d", opcode)
	}

	str := string(msg[1 : len(msg)-96])
	if str != "Hello, world!" {
		t.Fatalf("Received unexpected message: '%s'", str)
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

	Publish("a.b.c", []byte("Hello, world!"))

	msg := <-sub
	opcode := msg[0]
	if opcode != byte(consts.OpCodeBroadcast) {
		t.Fatalf("Received unexpected opcode: %d", opcode)
	}

	str := string(msg[1 : len(msg)-96])
	if str != "Hello, world!" {
		t.Fatalf("Received unexpected message: '%s'", str)
	}
}

func TestPublishWithoutSubscriber(_ *testing.T) {
	Publish("a.b.c", []byte("Hello, world!"))
}
