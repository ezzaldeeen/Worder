package pkg

import "errors"

// todo: SrcFileLoadingErrCode should be ErrCodeSrcFileLoading, to be faster to find
// todo: SrcFileLoadingErrCode should be ErrCodeLoadingSrcFile
// todo: these errors should be in the usage (closer to the usage)
const (
	SrcFileLoadingErrCode = 12
	SrcFileReadingErrCode = 13
	FileCreatingErrCode   = 14
	FileWritingErrCode    = 15
	InvalidSizeErrCode    = 16
	InvalidUnitErrCode    = 17
)

// todo: InvalidSizeErr should be ErrInvalidSize, to be faster to find
var (
	InvalidSizeErr    = errors.New("invalid size input, your input should look like: xxMB, xxKB")
	InvalidUnitErr    = errors.New("invalid unit, available units: KB, MB, and GB")
	SrcFileLoadingErr = errors.New("unable to load the source file, make sure that path is correct")
	SrcFileReadingErr = errors.New("unable to read the source file")
	FileCreatingErr   = errors.New("unable to create a file, make sure that the destination is correct")
	FileWritingErr    = errors.New("unable to writing to a file")
)
