package outline

import (
	"net/http"
	"time"
)

// Manager structure for Manager instance
type Manager struct {
	apiURL  string
	apiCrt  string
	timeout time.Duration
	client  *http.Client
}

// ServerInfo structure for storing server-related data
type ServerInfo struct {
	Name               string `json:"name"`
	ServerId           string `json:"serverId"`
	MetricsEnabled     bool   `json:"metricsEnabled"`
	CreatedTimestampMs int64  `json:"createdTimestampMs"`
	Version            string `json:"version"`
	AccessKeyDataLimit struct {
		Bytes int `json:"bytes"`
	} `json:"accessKeyDataLimit"`
	PortForNewAccessKeys  int    `json:"portForNewAccessKeys"`
	HostnameForAccessKeys string `json:"hostnameForAccessKeys"`
}

// MetricsState structure for storing metrics
type MetricsState struct {
	MetricsEnabled bool `json:"metricsEnabled"`
}
