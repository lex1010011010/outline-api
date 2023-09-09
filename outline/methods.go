package outline

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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
