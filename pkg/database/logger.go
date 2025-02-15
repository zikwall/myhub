package database

type Logger struct {
	callback func(format string, v ...interface{})
}

func (l *Logger) SetCallback(callback func(format string, v ...interface{})) {
	l.callback = callback
}

// nolint:asasalint // it's OK
func (l *Logger) Printf(format string, v ...interface{}) {
	l.callback(format, v)
}
