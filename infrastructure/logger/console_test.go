package infrastructure

import (
	"github.com/acazau/cloud-manager/domain"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestInfo(t *testing.T) {
	Convey("Validate Info inputs", t, func() {
		logger := new(LoggerConsole)

		Convey("Validate when info and message is valid, returns null error", func() {
			err := logger.Log(domain.Info, "something")
			So(err, ShouldBeNil)
		})

		Convey("Validate when error and message is valid, returns null error", func() {
			err := logger.Log(domain.Error, "something")
			So(err, ShouldBeNil)
		})

		Convey("Validate when message is null, returns not null error", func() {
			err := logger.Log(domain.Info, nil)
			So(err, ShouldNotBeNil)
		})

	})
}
