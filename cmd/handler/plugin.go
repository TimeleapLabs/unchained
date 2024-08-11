package handler

import (
	"os"

	"github.com/TimeleapLabs/unchained/cmd/handler/plugins"
	"github.com/gorilla/websocket"
	"github.com/spf13/cobra"
)

var conn *websocket.Conn

func Read() <-chan []byte {
	out := make(chan []byte)

	go func() {
		for {
			_, payload, err := conn.ReadMessage()
			if err != nil {
				panic(err)
			}

			out <- payload
		}
	}()

	return out
}

// plugin represents the plugin command.
var plugin = &cobra.Command{
	Use:   "plugin",
	Short: "Run an Unchained plugin locally",
	Long:  `Run an Unchained plugin locally`,

	Run: func(cmd *cobra.Command, _ []string) {
		os.Exit(1)
	},
}

// WithPluginCmd appends the plugin command to the root command.
func WithPluginCmd(cmd *cobra.Command) {
	cmd.AddCommand(plugin)
}

func init() {
	plugins.WithAIPluginCmd(plugin)
	plugins.WithTextToImagePluginCmd(plugin)
	plugins.WithTranslatePluginCmd(plugin)
}
