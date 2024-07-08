package plugins

import (
	"os"

	"github.com/TimeleapLabs/unchained/internal/service/ai"
	"github.com/spf13/cobra"
)

// worker represents the worker command.
var textToImagePlugin = &cobra.Command{
	Use:   "text-to-image",
	Short: "Run the text-to-image plugin locally",
	Long:  `Run the text-to-image plugin locally`,

	Run: func(cmd *cobra.Command, _ []string) {

		prompt := cmd.Flags().Lookup("prompt").Value.String()
		negativePrompt := cmd.Flags().Lookup("negative-prompt").Value.String()
		output := cmd.Flags().Lookup("output").Value.String()
		model := cmd.Flags().Lookup("model").Value.String()
		loraWeights := cmd.Flags().Lookup("lora-weights").Value.String()
		steps, err := cmd.Flags().GetUint8("inference")

		if err != nil {
			panic(err)
		}

		outputBytes := ai.TextToImage(prompt, negativePrompt, model, loraWeights, steps)

		// write outputBytes as png to output file path of output flag
		os.WriteFile(output, outputBytes, 0644)

		CloseSocket()
		os.Exit(0)
	},
}

// WithRunCmd appends the run command to the root command.
func WithTextToImagePluginCmd(cmd *cobra.Command) {
	cmd.AddCommand(textToImagePlugin)
}

func init() {
	textToImagePlugin.Flags().StringP(
		"prompt",
		"p",
		"",
		"Prompt data to process",
	)
	textToImagePlugin.Flags().StringP(
		"negative-prompt",
		"n",
		"",
		"Negative prompt data to process",
	)
	textToImagePlugin.Flags().Uint8P(
		"inference",
		"i",
		16,
		"Number of inference steps",
	)
	textToImagePlugin.Flags().StringP(
		"model",
		"m",
		"OEvortex/PixelGen",
		"Model to use for inference",
	)
	textToImagePlugin.Flags().StringP(
		"lora-weights",
		"w",
		"",
		"Lora weights model name (if applicable)",
	)
	textToImagePlugin.Flags().StringP(
		"output",
		"o",
		"",
		"Output file path",
	)

	textToImagePlugin.MarkFlagRequired("prompt")
	textToImagePlugin.MarkFlagRequired("output")
}
