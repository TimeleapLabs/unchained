package plugins

import (
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	sia "github.com/TimeleapLabs/go-sia/v2/pkg"
	"github.com/spf13/cobra"
)

// worker represents the worker command.
var translatePlugin = &cobra.Command{
	Use:   "translate",
	Short: "Run the translate plugin locally",
	Long:  `Run the translate plugin locally`,

	Run: func(cmd *cobra.Command, _ []string) {

		input := cmd.Flags().Lookup("input").Value.String()
		from := cmd.Flags().Lookup("from").Value.String()
		to := cmd.Flags().Lookup("to").Value.String()

		var err error
		conn, _, err = websocket.DefaultDialer.Dial(
			"ws://localhost:8765", nil,
		)

		if err != nil {
			panic(err)
		}

		incoming := Read()

		requestUUID, err := uuid.NewV7()
		if err != nil {
			panic(err)
		}

		uuidBytes, err := requestUUID.MarshalBinary()
		if err != nil {
			panic(err)
		}

		payload := sia.New().
			AddUInt16(1).
			AddByteArrayN(uuidBytes).
			AddStringN(from).
			AddStringN(to).
			AddString16(input).
			Bytes()

		err = conn.WriteMessage(websocket.BinaryMessage, payload)
		if err != nil {
			panic(err)
		}

		// Start a goroutine to print dots every 3 seconds
		stopDots := make(chan struct{})
		go func() {
			for {
				select {
				case <-stopDots:
					return
				case <-time.After(1 * time.Second):
					fmt.Print(".")
				}
			}
		}()

		data := <-incoming

		// Stop the dot-printing goroutine
		close(stopDots)
		fmt.Println()

		// process data
		s := sia.NewFromBytes(data)
		uuidBytesFromResponse := s.ReadByteArrayN(16)
		responseUUID, err := uuid.FromBytes(uuidBytesFromResponse)
		if err != nil {
			panic(err)
		}

		if requestUUID != responseUUID {
			panic("UUID mismatch")
		}

		translated := s.ReadString16()
		fmt.Println(translated)

		CloseSocket()
		os.Exit(0)
	},
}

// WithRunCmd appends the run command to the root command.
func WithTranslatePluginCmd(cmd *cobra.Command) {
	cmd.AddCommand(translatePlugin)
}

func init() {
	translatePlugin.Flags().StringP(
		"input",
		"i",
		"",
		"Input text to translate",
	)
	translatePlugin.Flags().StringP(
		"from",
		"f",
		"en",
		"From language",
	)
	translatePlugin.Flags().StringP(
		"to",
		"t",
		"fr",
		"To language",
	)
}
