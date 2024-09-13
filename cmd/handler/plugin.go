package handler

import (
	"os"

	"github.com/TimeleapLabs/unchained/cmd/handler/plugins"
	"github.com/spf13/cobra"
)

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
