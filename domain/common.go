package domain

type UUID [16]byte

type ILogger interface {
	Info(message interface{})
}

type Logger struct {
	ILogger ILogger
}

func (logger *Logger) Info(message interface{}) {
	if logger.ILogger != nil {
		logger.ILogger.Info(message)
	}
	return
}
