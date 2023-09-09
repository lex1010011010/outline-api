package outline

type ManagerInterface interface {
	GetServerInfo() (ServerInfo, error)
	UpdateServerHostname(newHostname string) error
	UpdateServerName(newName string) error
	GetMetricsStatus() (MetricsState, error)
	UpdateMetricsStatus(metricsEnabled bool) error

	UpdateDefaultPort(newPort int) error
	UpdateDataLimit(dataLimit int) error
	DeleteDataLimit() error

	CreatesNewAccessKey()
}
