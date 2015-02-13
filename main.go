package main

import (
	"flag"
	"fmt"
	"github.com/acazau/cloud-manager/domain"
	"github.com/acazau/cloud-manager/infrastructure"
	console "github.com/acazau/cloud-manager/infrastructure/logger"
	"github.com/acazau/cloud-manager/interfaces"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"os/signal"
)

var (
	_port              = flag.Int("port", 8090, "Port for web server")
	_host              = flag.String("host", "localhost", "Address for web server")
	_logger            = new(domain.Logger)
	_webserviceHandler = new(interfaces.WebserviceHandler)
)

func init() {
	_logger.ILogger = &console.Logger{}
	aws := &infrastructure.AWSRepository{}
	do := &infrastructure.DigitalOceanRepository{}
	// Inject logger into aws service
	aws.Logger = *_logger
	// Inject aws service into webservice
	instanceHandlers := make(map[string]interfaces.IInstance)
	instanceHandlers["aws"] = aws
	instanceHandlers["digitalocean"] = do
	_webserviceHandler.IInstance = instanceHandlers
	// Inject logger into webservice
	_webserviceHandler.ILogger = *_logger
}

func main() {
	/*
	 * Handle SIGINT (CTRL+C)
	 */
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		_logger.Info("Shutting down cloud-manager service...")
		os.Exit(0)
	}()

	router := mux.NewRouter()
	router.HandleFunc("/instances", ListInstances).Methods("GET")

	_logger.Info(fmt.Sprintf("cloud-manager service started on %s:%d\n\n", *_host, *_port))
	http.ListenAndServe(fmt.Sprintf("%s:%d", *_host, *_port), router)
}

func ListInstances(w http.ResponseWriter, r *http.Request) {
	_webserviceHandler.ListInstances(w, r)
}
