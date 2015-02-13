package infrastructure

import (
	"log"
)

type Logger struct{}

func (logger *Logger) Info(message interface{}) {
	log.Println(message)
	return
}
