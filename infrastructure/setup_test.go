package infrastructure

import (
	"github.com/acazau/cloud-manager/domain"
)

type _FakeLogger struct{}

func (logManager *_FakeLogger) Log(logLevelType domain.LogLevelType, message interface{}) error {
	return nil
}
