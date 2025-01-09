package rpc

import (
	"log"
	"net/http"

	"github.com/TimeleapLabs/unchained/internal/config"
	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/service/rpc/dto"
	"github.com/TimeleapLabs/unchained/internal/transport/client/conn"
	"github.com/TimeleapLabs/unchained/internal/utils"
	"github.com/gorilla/websocket"
)

// Runtime is a type that holds the runtime of a function.
type Runtime string

const (
	Mock      Runtime = "Mock"
	WebSocket Runtime = "WebSocket"
)

func WithMockTask(pluginName string, name string) func(s *Worker) {
	functions := map[string]config.Function{}
	functions[name] = config.Function{
		Name: name,
	}

	return func(s *Worker) {
		s.plugins[name] = plugin{
			name:      pluginName,
			runtime:   Mock,
			functions: functions,
		}
	}
}

func WithWebSocket(pluginName string, functions []config.Function, url string) func(s *Worker) {
	functionsMap := map[string]config.Function{}
	for _, f := range functions {
		functionsMap[f.Name] = f
	}

	return func(s *Worker) {
		p := plugin{
			name:      pluginName,
			runtime:   WebSocket,
			functions: functionsMap,
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

				packet := new(dto.RPCResponse).FromSiaBytes(message)
				utils.Logger.
					With("ID", packet.ID).
					Info("RPC Response")

				// Release the resources
				if task, ok := s.currentTasks[packet.ID]; ok {
					s.cpuUsage -= task.CPU
					s.gpuUsage -= task.GPU
					delete(s.currentTasks, packet.ID)
				}

				if s.overloaded {
					s.overloaded = false
					// TODO: Notify the broker that we're not overloaded anymore
				}

				conn.Send(consts.OpCodeRPCResponse, message)
			}
		}()

		p.conn = wsConn
		s.plugins[pluginName] = p
	}
}
