package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"strconv"
	"worder/errors"
)

var (
	size  string
	count int
	path  string
)

var generatorCmd = &cobra.Command{
	Use:     "generate",
	Aliases: []string{"gen"},
	Short:   "Files generator based on sample from Lorem Ipsum",
	Long:    "", // todo:
	Example: "worder generate --size=100MB --count=50 --path=./data",
	Run: func(cmd *cobra.Command, args []string) {
		value, unit := extractValueUnitFrom(size)
		fmt.Println(value, unit)
	},
}

func init() {
	generatorCmd.Flags().StringVarP(&size, "size", "s",
		"100MB", "The size of the generated file MB, and KB")
	generatorCmd.Flags().IntVarP(&count, "count", "c",
		50, "The number of generated files")
	generatorCmd.Flags().StringVarP(&path, "path", "p",
		"./data", "Directory destination")
}

// extractValueUnitFrom
func extractValueUnitFrom(input string) (int, string) {
	inputLen := len(input)
	if inputLen <= 2 {
		log.Fatalf("- %s\n", errors.InvalidSizeValue)
	}
	unit := input[inputLen-2:]
	value, err := strconv.Atoi(input[:inputLen-2])

	if err != nil {
		log.Fatalf("- %s\n", errors.InvalidSizeValue)
	}
	return value, unit
}
