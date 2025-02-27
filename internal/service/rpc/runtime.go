package rpc

import (
	"log"
	"net/http"

	"crypto/ed25519"

	"github.com/TimeleapLabs/timeleap/internal/config"
	"github.com/TimeleapLabs/timeleap/internal/consts"
	"github.com/TimeleapLabs/timeleap/internal/service/rpc/dto"
	"github.com/TimeleapLabs/timeleap/internal/service/rpc/worker"
	"github.com/TimeleapLabs/timeleap/internal/transport/client/conn"
	"github.com/TimeleapLabs/timeleap/internal/transport/server/packet"
	"github.com/TimeleapLabs/timeleap/internal/transport/server/websocket/queue"
	"github.com/TimeleapLabs/timeleap/internal/utils"
	"github.com/btcsuite/btcutil/base58"
	"github.com/gorilla/websocket"
)

func WithMockTask(pluginName string, name string) func(s *worker.Worker) {
	functions := map[string]config.Function{}
	functions[name] = config.Function{
		Name: name,
	}

	return func(s *worker.Worker) {
		s.Plugins[name] = dto.Plugin{
			Name:      pluginName,
			Runtime:   worker.Mock,
			Functions: functions,
		}
	}
}

func WithWebSocket(
	pluginName string,
	functions []config.Function,
	url string,
	publicKey string,
) func(s *worker.Worker) {
	pkBytes := base58.Decode(publicKey)
	pk := ed25519.PublicKey(pkBytes)

	functionsMap := map[string]config.Function{}
	for _, f := range functions {
		functionsMap[f.Name] = f
	}

	return func(s *worker.Worker) {
		p := dto.Plugin{
			Name:      pluginName,
			Runtime:   worker.WebSocket,
			Functions: functionsMap,
		}

		wsConn, httpResp, err := websocket.DefaultDialer.Dial(url, nil)
		if err != nil {
			panic(err)
		}

		if httpResp.StatusCode != http.StatusSwitchingProtocols {
			panic("Failed to establish websocket connection")
		}

		err = httpResp.Body.Close()
		if err != nil {
			panic(err)
		}

		go func() {
			for {
				_, message, err := wsConn.ReadMessage()
				if err != nil {
					log.Println("Read error:", err)
					break
				}

				if len(message) == 0 {
					continue
				}

				p := new(packet.Packet).FromBytes(message)

				if !p.IsValid() {
					utils.Logger.Error("Invalid Packet")
					continue
				}

				if !pk.Equal(p.Signer) {
					utils.Logger.
						With("Signer", base58.Encode(p.Signer)).
						With("Expected", base58.Encode(pk)).
						Error("Invalid Public Key")
					continue
				}

				response := new(dto.RPCResponse).FromSiaBytes(message)

				utils.Logger.
					With("ID", response.ID).
					Info("RPC Response")

				// Release the resources
				if task, ok := s.CurrentTasks.Load(response.ID); ok {
					s.CPUUsage -= task.CPU
					s.GPUUsage -= task.GPU
					s.RAMUsage -= task.RAM
					s.CurrentTasks.Delete(response.ID)
				}

				if s.CPUUsage < s.MaxCPU && s.GPUUsage < s.MaxGPU && s.RAMUsage < s.MaxRAM {
					// TODO: Notify the broker that we're not overloaded anymore
					s.Overloaded = false
				}

				conn.SendSigned(consts.OpCodeRPCResponse, message)
			}
		}()

		p.Conn = wsConn
		p.Writer = queue.NewWebSocketWriter(wsConn, 100) // TODO: Make the buffer size configurable
		s.Plugins[pluginName] = p
	}
}
