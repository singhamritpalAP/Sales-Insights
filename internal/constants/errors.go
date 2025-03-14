package constants

import "errors"

var (
	ErrInvalidLimit     = errors.New("invalid 'n' parameter for total records")
	ErrInvalidStartDate = errors.New("invalid start_date")
	ErrInvalidEndDate   = errors.New("invalid end_date")
)
