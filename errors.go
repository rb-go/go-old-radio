package oldradio

import "errors"

// package errors
var (
	ErrMainPipeEnded       = errors.New("oldradio: main pipe ended")
	ErrBackgroundPipeEnded = errors.New("oldradio: background pipe ended")
)
