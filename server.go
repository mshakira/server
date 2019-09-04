/*
Server serves incidents from incidents store based on the request.
ServiceNow is used as incidents store.
SSL is enabled. Api responses are in JSON format.
*/

package main

import (
	log "github.com/Sirupsen/logrus"
	"net/http"
	"server/handlers"
	"server/store/snow"
)

//var serviceNowStore *snow.ServicenowStore

func main() {

	// Initialize http multiplexer
	mux := http.NewServeMux()

	// Initialize log formatter
	Formatter := new(log.TextFormatter)
	Formatter.TimestampFormat = "02-01-2006 15:04:05"
	Formatter.FullTimestamp = true
	log.SetFormatter(Formatter)

	listHandler := &handlers.Handler{}

	// initialize http list handler
	listHandler.SnowStore, _ = snow.Init("")

	// Add the handler for /api/v1/list/incidents api call
	mux.HandleFunc("/api/v1/list/incidents", listHandler.ListHandler)
	mux.HandleFunc("/", handlers.RootHandler)
	//http.ListenAndServe(":3000", nil)

	// enable SSL
	// Add requestLogger
	log.Info("Server starting...")
	err := http.ListenAndServeTLS(":443", "server.crt", "server.key", listHandler.RequestLogger(mux))
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
