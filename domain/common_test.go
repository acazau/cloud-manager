package domain

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

type _FakeLogger struct{}

func (logManager *_FakeLogger) Log(logLevelType LogLevelType, message interface{}) error {
	return nil
}

func TestInfo(t *testing.T) {
	Convey("Validate Info inputs", t, func() {

		Convey("Validate when logger is null, returns null error", func() {
			logManager := new(LogManager)
			err := logManager.Log(Info, "something")
			So(err, ShouldNotBeNil)
		})

		Convey("Validate when logger and message are valid, returns not null error", func() {
			logManager := new(LogManager)
			logManager.InjectedLogManager = &_FakeLogger{}
			err := logManager.Log(Info, "something")
			So(err, ShouldBeNil)
		})

		Convey("Validate when message is null, returns null error", func() {
			logManager := new(LogManager)
			logManager.InjectedLogManager = &_FakeLogger{}
			err := logManager.Log(Info, nil)
			So(err, ShouldNotBeNil)
		})
	})
}
