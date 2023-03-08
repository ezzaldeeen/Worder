package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var version = "0.0.1"

var rootCmd = &cobra.Command{
	Use:     "worder",
	Version: version,
	Short:   "worder - simple CLI to generate and count words",
	Long: `Worder CLI is a word generator, and counter.
Uses a sample from Lorem Ipsum to generate files based on different count, and size
and provide the capability to count the number of words in these generated files.`,
	Run: func(cmd *cobra.Command, args []string) {},
}

func init() {
	rootCmd.AddCommand(generatorCmd)
	rootCmd.AddCommand(counterCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		// todo: check the below hint
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'\n", err)
		os.Exit(1)
	}
}
