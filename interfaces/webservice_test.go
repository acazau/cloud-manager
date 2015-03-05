package interfaces

import (
	"errors"
	"github.com/acazau/cloud-manager/domain"
	api "github.com/acazau/cloud-manager/usecases/api/v0"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"net/http/httptest"
	"testing"
)

type _FakeLogger struct{}
type _GoodInstanceRepository struct{}
type _BadInstanceRepository struct{}

func (logManager *_FakeLogger) Log(logLevelType domain.LogLevelType, message interface{}) error {
	return nil
}

func (repo *_GoodInstanceRepository) ListInstances() ([]api.Instance, error) {
	var instances = []api.Instance{
		api.Instance{
			Id:   domain.UUID{190, 125, 210, 117, 63, 141, 66, 212, 89, 174, 52, 5, 210, 65, 131, 250},
			Name: "test",
		},
	}
	return instances, nil
}

func (repo *_BadInstanceRepository) ListInstances() ([]api.Instance, error) {
	var instances []api.Instance
	return instances, errors.New("Some exception")
}

func TestListInstances(t *testing.T) {
	logger := new(domain.LogManager)
	logger.InjectedLogManager = &_FakeLogger{}

	Convey("Validate API ListInstances results", t, func() {
		Convey("should be able to list instances", func() {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/instances", nil)
			instanceHandlers := make(map[string]IInstance)
			instanceHandlers["good"] = &_GoodInstanceRepository{}
			handler := &WebServiceHandler{}
			handler.Instance = instanceHandlers
			handler.ListInstances(w, req)
			So(w.Body.String(), ShouldEqual, `[{"id":[190,125,210,117,63,141,66,212,89,174,52,5,210,65,131,250],"name":"test"}]`)
		})

		Convey("should generate 500 server error when instance repository produce error ", func() {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/instances", nil)
			instanceHandlers := make(map[string]IInstance)
			instanceHandlers["bad"] = &_BadInstanceRepository{}
			handler := &WebServiceHandler{}
			handler.Instance = instanceHandlers
			handler.ListInstances(w, req)
			So(w.Code, ShouldEqual, http.StatusInternalServerError)
		})
	})
}
