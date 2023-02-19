package truequeslib

import "errors"

var (
	// messages error
	ErrLogFilenameEmpty = errors.New("err-log_filename_empty")
	ErrCreatingLogFile  = errors.New("err-creating_logFile")
	ErrLoadingConfig    = errors.New("err-load_config")
	ErrRunningApp       = errors.New("err-running_app")
	ErrItemNotFound     = errors.New("err-item_not_found")

	// messages info
)
