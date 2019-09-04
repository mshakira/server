package handlers

import (
	log "github.com/Sirupsen/logrus"
	"net/http"
	"server/store/snow"
	"time"
)

type Handler struct {
	SnowStore *snow.ServicenowStore
}

func (handler *Handler) ListHandler(w http.ResponseWriter, r *http.Request) {
	// get the response from serviceNow obj
	js, err := handler.SnowStore.GetList()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the content-type header to json
	w.Header().Set("Content-Type", "application/json")
	w.Write(*js)
}

// Log each http request for better tracking
func (handler *Handler) RequestLogger(targetMux http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		targetMux.ServeHTTP(w, r)

		// log request by who(IP address)
		requesterIP := r.RemoteAddr

		log.Printf(
			"%s\t%s\t%s\t%v",
			r.Method,
			r.RequestURI,
			requesterIP,
			time.Since(start),
		)
	})
}
