package healthz

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

var (
	fatalErrors    []HealthError
	nonFatalErrors []HealthError
	mutex          sync.Mutex
	fatalmutex     sync.Mutex
)

type response struct {
	FatalErrors    []HealthError `json:"fatal_errors"`
	NonFatalErrors []HealthError `json:"non_fatal_errors"`
}

// HealthError represents a single error, fatal or non, to be created by the package consumer and submitted through one of the two setter functions
type HealthError struct {
	Description string            `json:"description"`
	Error       string            `json:"error"`
	Metadata    map[string]string `json:"metadata"`
	Type        string            `json:"type"`
}

// NewFatalError appends a fatal error to the global fatal error slice
func NewFatalError(err HealthError) {
	fatalmutex.Lock()
	fatalErrors = append(fatalErrors, err)
	fatalmutex.Unlock()
	return
}

// NewNonFatalError appends a non fatal error to the global non fatal error slice
func NewNonFatalError(err HealthError) {
	mutex.Lock()
	nonFatalErrors = append(nonFatalErrors, err)
	mutex.Unlock()
	return
}

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	var statusCode int

	if len(fatalErrors) > 0 {
		statusCode = http.StatusInternalServerError
	} else {
		statusCode = http.StatusOK
	}

	response := response{FatalErrors: fatalErrors, NonFatalErrors: nonFatalErrors}

	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	return
}

// Serve serves the healthz endpoint on the specified endpoint and url, defaults to localhost:8080/healthz
func Serve(url string, endpoint string) {
	if endpoint == "" {
		fmt.Printf("Endpoint not specified, defaulting to /healthz")
		http.HandleFunc("/healthz", healthzHandler)
	} else {
		http.HandleFunc(endpoint, healthzHandler)
	}

	if url == "" {
		fmt.Printf("Url/Port not specified, defaulting to localhost:8080")
		go http.ListenAndServe("localhost:8080", nil)
	} else {
		go http.ListenAndServe(url, nil)
	}
}
