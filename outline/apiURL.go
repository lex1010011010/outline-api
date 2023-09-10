package outline

const (
	ServerInfoURL           = "%s/server"                          //Returns information about the server
	HostnameURL             = "%s/server/hostname-for-access-keys" //Changes the hostname for access keys
	ServerNameURL           = "%s/name"                            //Renames the server
	MetricsURL              = "%s/metrics/enabled"                 //Enables or disables sharing of metrics
	PortForNewAccessKeysURL = "%s/server/port-for-new-access-keys" //Changes the default port for newly created access
	AccessKeyDataLimitURL   = "%s/server/access-key-data-limit"    //Access Key data limits actions
	AccessKeysURL           = "%s/access-keys"
)
