package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"worder/pkg/generatorv2"
)

type ByteUnit string

// MB represents MegaByte
// KB represents KiloByte
const (
	MB ByteUnit = "mb"
	KB ByteUnit = "kb"
)

// size of the generated files
// count of the generated files
// path of the destination for the generated files
var (
	size  string
	count int
	path  string
)

// generatorCmd is a command in the Worder ClI
// for generating files based on the sample file from Lorem Ipsum
// based on the given arguments which are: size, count, and path
var generatorCmd = &cobra.Command{
	Use:     "generate",
	Aliases: []string{"gen"},
	Short:   "Files generating based on sample from Lorem Ipsum",
	Example: "worder generate --size=100MB --count=50 --path=./data",
	Run: func(cmd *cobra.Command, args []string) {
		fileSize, unitSize := parse(size)
		// todo: depend on interface
		geor := generatorv2.TxtFileGenerator{
			Size:        fileSize,
			Unit:        unitSize,
			Count:       count,
			Destination: path,
		}
		wg := sync.WaitGroup{}
		wg.Add(count)

		geor.Generate(&wg)
		wg.Wait()
		fmt.Println(runtime.NumGoroutine())
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

// parse the size input expected to return
// the value, and the unit can be: MB, or KB
// for the wrong size the program will
// expose error message for the client
// todo: return an error
func parse(input string) (int, string) {
	inputLen := len(input)
	if inputLen <= 2 {
		//log.Fatalf("- %s\n", custom.InvalidUnitErr)
	}

	unit := strings.ToLower(input[inputLen-2:])
	value, err := strconv.Atoi(input[:inputLen-2])

	if err != nil {
		//log.Fatalf("- %s\n", custom.InvalidSizeValue)
	}

	switch ByteUnit(unit) {
	case MB, KB:
		return value, unit
	}
	//log.Fatalf("- %s\n", custom.InvalidSizeValue)
	return 0, ""
}
