package log

import "fmt"

func Trace(args ...interface{}) {
	msg := fmt.Sprint(args...)
	logger.Trace().Msg(msg)
}

func Tracef(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	logger.Trace().Msg(msg)
}
