package interfaces

import (
	"encoding/json"
	"fmt"
	"github.com/acazau/cloud-manager/domain"
	api "github.com/acazau/cloud-manager/usecases/api/v0"
	"io"
	"net/http"
	"strings"
)

type IInstance interface {
	ListInstances() ([]api.Instance, error)
}

type WebserviceHandler struct {
	IInstance map[string]IInstance
	ILogger   domain.Logger
}

func (handler *WebserviceHandler) ListInstances(w http.ResponseWriter, r *http.Request) {
	handler.ILogger.Info("ListInstances called...")
	var content string
	for i := range handler.IInstance {
		handler.ILogger.Info(fmt.Sprintf("repo ListInstances %s", i))
		instances, err := handler.IInstance[i].ListInstances()
		if err != nil {
			serveError(w, &handler.ILogger)
			return
		}

		jsonBytes, _ := json.Marshal(instances)
		content += strings.Replace(string(jsonBytes), "%", "%%", -1)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(200)
	fmt.Fprintf(w, content)

}

func serve404(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	io.WriteString(w, "404 error")
}

func serveError(w http.ResponseWriter, err *domain.Logger) {
	err.Info(err)
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
}
