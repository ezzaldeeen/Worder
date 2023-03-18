package cmd

import (
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"os"
	"worder/custom"
	"worder/pkg/generating"
	"worder/utils"
)

type ByteUnit string

// sourceSamplePath the path of the sample template (Lorem Ipsum)
const sourceSamplePath string = "./resources/sample.txt"

// size of the generated files
// count of the generated files
// path of the destination for the generated files
var (
	size  string
	count int
	path  string

	logger = utils.GetLogger(zap.NewDevelopment)
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
		size, unit, err := utils.ParseSize(size)
		if err != nil {
			logger.Error(err)
			os.Exit(custom.InvalidSizeErrCode)
		}

		size, err = utils.ConvertToByte(size, unit)
		if err != nil {
			logger.Error(err)
			os.Exit(custom.InvalidUnitErrCode)
		}

		generator := generating.NewTXTFilesGenerator(
			logger, sourceSamplePath, path, count, size)

		generator.Generate()
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
