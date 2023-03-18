package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"worder/pkg/counterv1"
)

var sourcePath string
var counterCmd = &cobra.Command{
	Use:     "count",
	Short:   "Word counter for the generated files",
	Long:    "",
	Example: "worder count --path=./data",
	Run: func(cmd *cobra.Command, args []string) {
		// todo: depend on interface
		counter := counterv1.Counter{Path: sourcePath}
		wordCount, err := counter.Count()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Total words count:", wordCount)
	},
}

func init() {
	counterCmd.Flags().StringVarP(&sourcePath, "path", "p",
		"./data", "Directory destination")
}
