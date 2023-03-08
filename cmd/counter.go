package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var counterCmd = &cobra.Command{
	Use:     "count",
	Short:   "Word counter for the generated files",
	Long:    "",
	Example: "worder count --path=./data",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("COUNTER :)")
	},
}
