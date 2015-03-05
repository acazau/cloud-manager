package infrastructure

import (
	"github.com/acazau/cloud-manager/domain"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestDigitalOceanListInstances(t *testing.T) {
	Convey("Validate ListInstances inputs", t, func() {
		do := &DigitalOceanRepository{}
		logger := new(domain.LogManager)

		Convey("Validate when logger is null, returns not null error", func() {
			_, err := do.ListInstances()
			So(err, ShouldNotBeNil)
		})

		Convey("Validate when logger is valid, returns null error", func() {
			logger.InjectedLogManager = &_FakeLogger{}
			do.LogManager = *logger
			_, err := do.ListInstances()

			So(err, ShouldBeNil)
		})
	})
}
