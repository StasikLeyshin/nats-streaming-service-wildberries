package startup

type NatsStreamingConfig struct {
	Host      string `yaml:"host"`
	Port      int    `yaml:"port"`
	ClusterID string `yaml:"cluster_id"`
	ClientID  string `yaml:"client_id"`
	Channel   string `yaml:"channel"`
}
