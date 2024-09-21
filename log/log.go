package log

var logger = func() *Logger {
	l := NewLogger()
	l.setCalldepth(4)
	return l
}()

func SetOutputFile(fileName string) {
	logger.SetOutputFile(fileName)
}

func SetLogLevel(level logLevel) {
	logger.SetLogLevel(level)
}

func Trace(format string, v ...interface{}) {
	logger.Trace(format, v...)
}

func Debug(format string, v ...interface{}) {
	logger.Debug(format, v...)
}

func Info(format string, v ...interface{}) {
	logger.Info(format, v...)
}

func Warn(format string, v ...interface{}) {
	logger.Warn(format, v...)
}

func Error(format string, v ...interface{}) {
	logger.Error(format, v...)
}

func Fatal(format string, v ...interface{}) {
	logger.Fatal(format, v...)
}

func Critical(format string, v ...interface{}) {
	logger.Critical(format, v...)
}
