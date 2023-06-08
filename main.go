package main

import (
	"fmt"
	"os"

	"github.com/izaakdale/key/generate"

	"github.com/spf13/cobra"
)

func main() {
	root := &cobra.Command{}

	root.AddCommand(generate.Base64RSAPair)

	if err := root.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "execution error: %s", err.Error())
		os.Exit(1)
	}
}
