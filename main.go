package main

import (
	"os"

	"github.com/chuck-horowitz/pvbeat/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
