package main

import (
	"os"

	"github.com/andrewdef/sasbeat/cmd"

	_ "github.com/andrewdef/sasbeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
