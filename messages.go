package truequeslib

import "errors"

var (
	// messages error
	ErrCreatingApp  = errors.New("err-creating_app")
	ErrRunningApp   = errors.New("err-running_app")
	ErrItemNotFound = errors.New("err-item_not_found")
)
