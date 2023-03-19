package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"worder/pkg/counting"
)

var sourcePath string
var counterCmd = &cobra.Command{
	Use:     "count",
	Short:   "Word counter for the generated files",
	Long:    "",
	Example: "worder count --path=./data",
	Run: func(cmd *cobra.Command, args []string) {
		counter := counting.NewWordCounter(
			logger, sourcePath)
		totalWords := counter.Count()
		fmt.Println("Total Number of Words:", totalWords)
	},
}

func init() {
	counterCmd.Flags().StringVarP(&sourcePath, "path", "p",
		"./data", "Directory destination")
}
