package outline

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	ServerInfoURL = "%s/server" // Check documentation for the actual URL
	ServerNameURL = "%s/name"   // Replace with the actual URL if different
)

type ManagerInterface interface {
	ServerInfo() (ServerInfo, error)
	ChangeHostname(newHostname string) error
	RenameServer(newName string) error
}

type Manager struct {
	apiURL  string
	apiCrt  string
	timeout time.Duration
	client  *http.Client
}

// NewManager initializes new Manager
func NewManager(apiURL string, apiCrt string, timeout time.Duration) *Manager {
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

// ServerInfo retrieves server information
func (m *Manager) ServerInfo() (ServerInfo, error) {
	var info ServerInfo
	url := fmt.Sprintf(ServerInfoURL, m.apiURL)
	resp, err := m.client.Get(url)
	if err != nil {
		return ServerInfo{}, err
	}
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil && err == nil {
			err = fmt.Errorf("failed to close response body: %v", closeErr)
		}
	}()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ServerInfo{}, err
	}
	err = handleHTTPResponse(resp)
	if err != nil {
		return ServerInfo{}, err
	}
	err = json.Unmarshal(body, &info)
	if err != nil {
		return ServerInfo{}, err
	}
	return info, nil
}

// ChangeHostname changes the hostname for all access keys.
func (m *Manager) ChangeHostname(newHostname string) (err error) {
	const HostnameURL = "%s/hostname"
	payload := map[string]string{"hostname": newHostname}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	url := fmt.Sprintf(HostnameURL, m.apiURL)
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := m.client.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil && err == nil {
			err = fmt.Errorf("failed to close response body: %v", closeErr)
		}
	}()
	return handleHTTPResponse(resp)
}

// RenameServer renames the server.
func (m *Manager) RenameServer(newName string) (err error) {
	payload := map[string]string{"name": newName}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	url := fmt.Sprintf(ServerNameURL, m.apiURL)
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := m.client.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil && err == nil {
			err = fmt.Errorf("failed to close response body: %v", closeErr)
		}
	}()
	return handleHTTPResponse(resp)
}
