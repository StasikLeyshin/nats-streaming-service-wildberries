package startup

import (
	"fmt"
	"github.com/nats-streaming-service-wildberries/internal/http"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Database      DatabaseConfig      `yaml:"database"`
	Http          http.Config         `yaml:"http"`
	NatsStreaming NatsStreamingConfig `yaml:"nats_streaming"`
}

func NewConfig(configFile string) (*Config, error) {
	rawYAML, err := os.ReadFile(configFile)
	if err != nil {
		return nil, fmt.Errorf("reading file error: %w", err)
	}

	cfg := &Config{}
	if err = yaml.Unmarshal(rawYAML, cfg); err != nil {
		return nil, fmt.Errorf("yaml parsing error: %w", err)
	}

	return cfg, nil
}
