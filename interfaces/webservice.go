package interfaces

import (
	"encoding/json"
	"fmt"
	"github.com/acazau/cloud-manager/domain"
	api "github.com/acazau/cloud-manager/usecases/api/v0"
	"net/http"
	"strings"
)

type IInstance interface {
	ListInstances() ([]api.Instance, error)
}

type WebServiceHandler struct {
	Instance   map[string]IInstance
	LogManager domain.LogManager
}

func (handler *WebServiceHandler) ListInstances(w http.ResponseWriter, r *http.Request) {
	handler.LogManager.Log(domain.Info, "ListInstances called...")
	var content string
	for i := range handler.Instance {
		handler.LogManager.Log(domain.Info, fmt.Sprintf("repo ListInstances %s", i))
		instances, err := handler.Instance[i].ListInstances()
		if err != nil {
			serveError(w, &handler.LogManager)
			return
		}

		jsonBytes, _ := json.Marshal(instances)
		content += strings.Replace(string(jsonBytes), "%", "%%", -1)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(200)
	fmt.Fprintf(w, content)
}

func serveError(w http.ResponseWriter, err *domain.LogManager) {
	err.Log(domain.Info, err)
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
}
