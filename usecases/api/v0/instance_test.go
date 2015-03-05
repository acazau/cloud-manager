package usecases

import (
	"errors"
	"github.com/acazau/cloud-manager/domain"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

type _FakeLogger struct{}
type _GoodInstanceInteractor struct{}
type _BadInstanceInteractor struct{}

func (repo *_GoodInstanceInteractor) ListInstances() ([]Instance, error) {
	var instances = []Instance{
		Instance{
			Id:   domain.UUID{190, 125, 210, 117, 63, 141, 66, 212, 89, 174, 52, 5, 210, 65, 131, 250},
			Name: "test",
		},
	}
	return instances, nil
}

func (repo *_BadInstanceInteractor) ListInstances() ([]Instance, error) {
	var instances []Instance
	return instances, errors.New("Some error")
}

func TestListInstances(t *testing.T) {
	Convey("Validate Instance Interactor ListInstances results", t, func() {

		Convey("Validate when interactor is null, returns not null error", func() {
			_, err := (&InstanceProviderManager{}).ListInstances()
			So(err.Error(), ShouldEqual, `Instance Provider Repository cannot be null`)
		})

		Convey("Validate when no instances found, returns should be empty and null error", func() {
			instanceProviderManager := &InstanceProviderManager{InjectedInstanceProvider: &_BadInstanceInteractor{}}
			results, err := instanceProviderManager.ListInstances()
			So(&results, ShouldBeEmpty)
			So(err, ShouldNotBeNil)
		})

		Convey("Validate when instances found, returns should not be empty and null error", func() {
			instanceProviderManager := &InstanceProviderManager{InjectedInstanceProvider: &_GoodInstanceInteractor{}}
			results, err := instanceProviderManager.ListInstances()
			So(&results, ShouldNotBeEmpty)
			So(err, ShouldBeNil)
		})
	})
}
