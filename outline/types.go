package outline

type OutlineAPI interface {
	All() ([]map[string]interface{}, error)
	ChangePassword(id int, newPassword string) (bool, error)
	Rename(id int, newName string) (bool, error)
	AllActive() (map[string]int64, error)
	New(label string) (map[string]interface{}, error)
}

// ServerInfo structure for storing server-related data
type ServerInfo struct {
	Name                  string `json:"name"`
	ServerId              string `json:"serverId"`
	MetricsEnabled        bool   `json:"metricsEnabled"`
	CreatedTimestampMs    int64  `json:"createdTimestampMs"`
	Version               string `json:"version"`
	PortForNewAccessKeys  int    `json:"portForNewAccessKeys"`
	HostnameForAccessKeys string `json:"hostnameForAccessKeys"`
}
