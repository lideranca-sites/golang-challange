package ports

type LoggerPort interface {
	Info(message string, args ...any)
	Debug(message string, args ...any)
	Warn(message string, args ...any)
	Error(message string, args ...any)
}
