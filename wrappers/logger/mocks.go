package logger

type Mock struct {
	DebugFn func(format string, v ...interface{})
	InfoFn  func(format string, v ...interface{})
	ErrorFn func(format string, v ...interface{})
}

func (m Mock) Debug(format string, v ...interface{}) {
	if m.DebugFn != nil {
		m.DebugFn(format, v...)
	}
}

func (m Mock) Info(format string, v ...interface{}) {
	if m.InfoFn != nil {
		m.InfoFn(format, v...)
	}
}

func (m Mock) Error(format string, v ...interface{}) {
	if m.ErrorFn != nil {
		m.ErrorFn(format, v...)
	}
}
