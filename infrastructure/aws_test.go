package infrastructure

import (
	"github.com/acazau/cloud-manager/domain"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestAWSListInstances(t *testing.T) {
	Convey("Validate ListInstances inputs", t, func() {
		aws := &AWSRepository{}
		logger := new(domain.LogManager)

		Convey("Validate when logger is null, returns not null error", func() {
			_, err := aws.ListInstances()
			So(err, ShouldNotBeNil)
		})

		Convey("Validate when logger is valid, returns null error", func() {
			logger.InjectedLogManager = &_FakeLogger{}
			aws.LogManager = *logger
			_, err := aws.ListInstances()

			So(err, ShouldBeNil)
		})
	})
}
