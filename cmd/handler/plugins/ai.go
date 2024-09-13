package plugins

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/TimeleapLabs/unchained/internal/service/ai"
	"github.com/spf13/cobra"
)

// aiPlugin represents the ai command.
var aiPlugin = &cobra.Command{
	Use:   "ai",
	Short: "Start the Unchained ai server for local invocation",
	Long:  `Start the Unchained ai server for local invocation`,

	Run: func(cmd *cobra.Command, _ []string) {
		wg, cancel := ai.StartServer(cmd.Context())

		// Set up signal handling
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

		go func() {
			sig := <-sigChan
			log.Printf("Received signal: %v. Shutting down gracefully...", sig)
			cancel() // Cancel the context to stop all managed processes
		}()

		// Wait for all processes to finish
		wg.Wait()
		log.Println("All processes have been stopped.")
	},
}

// WithAIPluginCmd appends the run command to the root command.
func WithAIPluginCmd(cmd *cobra.Command) {
	cmd.AddCommand(aiPlugin)
}
