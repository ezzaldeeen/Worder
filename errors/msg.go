package errors

type ErrorMsg string

const (
	InvalidSizeValue ErrorMsg = `Wrong value for 'size' flag
Please make sure that the value following this pattern: xxMB, or xxKB - where xx is number.`
)
