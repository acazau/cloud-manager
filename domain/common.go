package domain

import (
	"errors"
)

type UUID [16]byte

type LogLevelType int

const (
	Debug LogLevelType = 1 << iota
	Error
	Info
	Fatal
	Trace
	Warning
)

type ILogManager interface {
	Log(logLevelType LogLevelType, message interface{}) error
}

type LogManager struct {
	InjectedLogManager ILogManager
}

func (logManager *LogManager) Log(logLevelType LogLevelType, message interface{}) error {
	if logManager.InjectedLogManager == nil {
		return errors.New("Injected LogManager cannot be null")
	}
	if message == nil {
		return errors.New("message param cannot be null")
	}
	logManager.InjectedLogManager.Log(logLevelType, message)
	return nil
}
