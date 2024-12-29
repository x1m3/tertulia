package config

import (
	"context"
	"time"

	"github.com/kelseyhightower/envconfig"

	"github.com/x1m3/tertulia/internal/pkg/log"
)

// Config holds the project configuration
type Config struct {
	ApiPort       int    `envconfig:"TERTULIA_HTTP_PORT" default:"3000"`
	MetricsPort   int    `envconfig:"TERTULIA_METRICS_PORT" default:"3001"`
	DatabaseUrl   string `envconfig:"TERTULIA_DB_URL"`
	RunTimeLimits RunTimeLimits
}

// RunTimeLimits holds the runtime limits configuration
type RunTimeLimits struct {
	// HttpClientTimeOut is the maximum time to wait for a response when calling pinata
	HttpClientTimeOut time.Duration `envconfig:"TERTULIA_HTTP_CLIENT_TIMEOUT" default:"10s"`
	// HttpRequestSize is the maximum size of any http request
	HttpRequestSize int64 `envconfig:"TERTULIA_HTTP_REQUEST_SIZE" default:"1048576"` // 1MB
	// HttpHeaderSize is the maximum header size of any http request
	HttpHeaderSize int `envconfig:"TERTULIA_HTTP_HEADER_SIZE" default:"16384"` // 16kB
	// HttpReadHeaderTimeout is the maximum time to read the request headers
	HttpReadHeaderTimeout time.Duration `envconfig:"TERTULIA_HTTP_READ_HEADER_TIMEOUT" default:"5s"`
	// HttpReadTimeout is the maximum time to read the whole request
	HttpReadTimeout time.Duration `envconfig:"TERTULIA_HTTP_READ_TIMEOUT" default:"60s"`
	// HttpWriteTimeout is the maximum time between the request is read and the response is written
	HttpWriteTimeout time.Duration `envconfig:"TERTULIA_HTTP_WRITE_TIMEOUT" default:"90s"`
	// MaxFileSize is the maximum size of the file that can be downloaded
	MaxFileSize int `envconfig:"TERTULIA_MAX_FILE_SIZE" default:"204800"` // 200 kB
}

// Load loads the configuration from the environment
func Load(ctx context.Context) (*Config, error) {
	const length = 10
	cfg := &Config{}
	if err := envconfig.Process("", cfg); err != nil {
		return nil, err
	}
	log.Debug(ctx, "database url", "url", obfuscate(cfg.DatabaseUrl[:min(length, len(cfg.DatabaseUrl))]+"***"))
	return cfg, nil
}

func obfuscate(s string) string {
	const minLen = 10
	if len(s) < minLen {
		return s
	}
	return s[:10] + "***"
}
