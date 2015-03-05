package infrastructure

import (
	"errors"
	"fmt"
	"github.com/acazau/cloud-manager/domain"
	"log"
)

type LoggerConsole struct{}

func (logManager *LoggerConsole) Log(logLevelType domain.LogLevelType, message interface{}) error {
	if message == nil {
		return errors.New("error: message param cannot be null")
	}
	switch logLevelType {
	case domain.Info:
		log.Println(fmt.Sprintf("Info: %s", message))
	case domain.Error:
		log.Println(fmt.Sprintf("Error: %s", message))
	}
	return nil
}
