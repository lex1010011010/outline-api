package outline

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// GetServerInfo get server information
func (m *Manager) GetServerInfo() (ServerInfo, error) {
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

// UpdateServerHostname changes the hostname for all access keys.
func (m *Manager) UpdateServerHostname(newHostname string) (err error) {
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

// UpdateServerName renames the server.
func (m *Manager) UpdateServerName(newName string) (err error) {
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

// GetMetricsStatus get server information
func (m *Manager) GetMetricsStatus() (MetricsState, error) {
	var info MetricsState
	url := fmt.Sprintf(MetricsURL, m.apiURL)
	resp, err := m.client.Get(url)
	if err != nil {
		return MetricsState{}, err
	}
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil && err == nil {
			err = fmt.Errorf("failed to close response body: %v", closeErr)
		}
	}()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return MetricsState{}, err
	}
	err = handleHTTPResponse(resp)
	if err != nil {
		return MetricsState{}, err
	}
	err = json.Unmarshal(body, &info)
	if err != nil {
		return MetricsState{}, err
	}
	return info, nil
}

// UpdateMetricsStatus enables or disables the sharing of metrics.
func (m *Manager) UpdateMetricsStatus(metricsEnabled bool) (err error) {
	payload := map[string]bool{"metricsEnabled": metricsEnabled}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	url := fmt.Sprintf(MetricsURL, m.apiURL)
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

// UpdateDefaultPort changes the default port for newly created access
func (m *Manager) UpdateDefaultPort(newPort int) error {
	// Check if the requested port is within the valid range (1-65535)
	if newPort < 1 || newPort > 65535 {
		return fmt.Errorf("invalid port value, must be between 1 and 65535")
	}

	// Create JSON payload
	payload := map[string]int{"port": newPort}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	// Create a new HTTP request
	url := fmt.Sprintf(PortForNewAccessKeysURL, m.apiURL)
	log.Println(url)
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	// Perform the request
	resp, err := m.client.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil && err == nil {
			err = fmt.Errorf("failed to close response body: %v", closeErr)
		}
	}()

	// Handle the response based on the status code
	switch resp.StatusCode {
	case http.StatusNoContent:
		return nil
	case http.StatusBadRequest:
		return fmt.Errorf("invalid port parameter or missing port parameter")
	case http.StatusConflict:
		return fmt.Errorf("the requested port is already in use by another service")
	default:
		return fmt.Errorf("unexpected status code %d", resp.StatusCode)
	}
}

// UpdateDataLimit sets a data transfer limit for all access keys in bytes.
func (m *Manager) UpdateDataLimit(dataLimit int) error {
	// Check if the data limit is valid (non-negative)
	if dataLimit < 0 {
		return fmt.Errorf("invalid data limit, must be a non-negative integer")
	}

	// Create JSON payload
	payload := map[string]interface{}{"limit": map[string]int{"bytes": dataLimit}}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	// Create a new HTTP request
	url := fmt.Sprintf(AccessKeyDataLimitURL, m.apiURL)
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	// Perform the request
	resp, err := m.client.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil && err == nil {
			err = fmt.Errorf("failed to close response body: %v", closeErr)
		}
	}()

	// Handle the response based on the status code
	switch resp.StatusCode {
	case http.StatusNoContent:
		return nil
	case http.StatusBadRequest:
		return fmt.Errorf("invalid data limit")
	default:
		return fmt.Errorf("unexpected status code %d", resp.StatusCode)
	}
}

// DeleteDataLimit removes the access key data limit, lifting data transfer restrictions on all access keys.
func (m *Manager) DeleteDataLimit() error {
	// Create a new HTTP request to send a DELETE request to the specified URL
	url := fmt.Sprintf(AccessKeyDataLimitURL, m.apiURL)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	// Perform the DELETE request
	resp, err := m.client.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil && err == nil {
			err = fmt.Errorf("failed to close response body: %v", closeErr)
		}
	}()

	// Handle the response based on the status code
	switch resp.StatusCode {
	case http.StatusNoContent:
		return nil
	default:
		return fmt.Errorf("unexpected status code %d", resp.StatusCode)
	}
}
