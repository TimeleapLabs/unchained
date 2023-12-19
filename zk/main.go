package main

import (
	"log"
	"os"
	"unchained-zk/lib"

	"github.com/urfave/cli/v2"
)

func main() {

	app := cli.App{
		Name:  "unchained-zk",
		Usage: "Unchained zk-SNARKs toolkit",
		Commands: []*cli.Command{
			{
				Name:    "compile",
				Aliases: []string{"c"},
				Usage:   "Compile the zk-SNARKs circuit",
				Action:  lib.CompileCommand,
			},
			{
				Name:    "test",
				Aliases: []string{"t"},
				Usage:   "Run a test on the zk-SNARKs circuit",
				Action:  lib.TestCommand,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
