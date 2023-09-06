package outline

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"time"
)

// NewManager initializes new Manager
func NewManager(apiURL string, apiCrt string, timeouts ...time.Duration) *Manager {
	// Set a default timeout
	var timeout time.Duration
	if len(timeouts) > 0 {
		timeout = timeouts[0]
	} else {
		timeout = 30 * time.Second // Default timeout
	}

	// Create a Transport to disable SSL verification
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	return &Manager{
		apiURL:  apiURL,
		apiCrt:  apiCrt,
		timeout: timeout,
		client: &http.Client{
			Transport: tr,
			Timeout:   timeout,
		},
	}
}

// handleHTTPResponse handles HTTP responses, returns an error if the status code is unexpected.
func handleHTTPResponse(resp *http.Response) error {
	switch resp.StatusCode {
	case http.StatusOK, http.StatusNoContent:
		return nil
	case http.StatusBadRequest:
		return fmt.Errorf("bad request")
	case http.StatusUnauthorized:
		return fmt.Errorf("unauthorized")
	case http.StatusForbidden:
		return fmt.Errorf("forbidden")
	case http.StatusNotFound:
		return fmt.Errorf("not found")
	case http.StatusInternalServerError:
		return fmt.Errorf("internal server error")
	default:
		return fmt.Errorf("unexpected status code %d", resp.StatusCode)
	}
}
