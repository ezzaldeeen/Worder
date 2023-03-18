package custom

import "errors"

var (
	InvalidUnitErr    = errors.New("invalid unit, available units: KB, MB, and GB")
	SrcFileLoadingErr = errors.New("unable to load the source file, make sure that path is correct")
	SrcFileReadingErr = errors.New("unable to read the source file")
	FileCreatingErr   = errors.New("unable to create a file, make sure that the destination is correct")
	FileWritingErr    = errors.New("unable to writing to a file")
)
