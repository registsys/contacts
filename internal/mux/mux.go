package mux

import (
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/registsys/contacts/internal/handlers/contacts"
	"github.com/registsys/contacts/internal/services"
)

var buckets = prometheus.ExponentialBuckets(0.1, 1.5, 5)
var requestDuration = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Name:    "http_request_duration_seconds",
		Help:    "Tracks the latencies for HTTP requests.",
		Buckets: buckets,
	},
	[]string{"status", "path", "method"},
)

func prometheusMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		recorder := &statusRecorder{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		start := time.Now()
		next.ServeHTTP(recorder, r)
		elapsed := time.Since(start)

		method := r.Method
		path := r.URL.Path
		status := strconv.Itoa(recorder.statusCode)

		requestDuration.WithLabelValues(status, path, method).Observe(elapsed.Seconds())
	})
}

type statusRecorder struct {
	http.ResponseWriter
	statusCode int
}

func (rec *statusRecorder) WriteHeader(code int) {
	rec.statusCode = code
	rec.ResponseWriter.WriteHeader(code)
}

func New(s *services.Services) http.Handler {

	reg := prometheus.NewRegistry()

	reg.MustRegister(requestDuration)

	handler := promhttp.HandlerFor(
		reg,
		promhttp.HandlerOpts{})

	mux := http.NewServeMux()

	contactsH := contacts.NewContactsHandler(s)

	mux.HandleFunc("/contacts", contactsH.ContactsHandler)
	mux.Handle("/metrics", handler)

	promHandler := prometheusMiddleware(mux)

	return promHandler
}
