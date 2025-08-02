package traefik_plugin_original_host_header

import (
	"context"
	"net/http"
)

type Config struct {
	HeaderName string `json:"headerName,omitempty"`
}

// initializes the default plugin config
func CreateConfig() *Config {
	return &Config{
		HeaderName: "X-Original-Host",
	}
}

// sets X-Original-Host header.
type OriginalHostMiddleware struct {
	next       http.Handler
	name       string
	headerName string
}

// creates a new middleware instance
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	return &OriginalHostMiddleware{
		next:       next,
		name:       name,
		headerName: config.HeaderName,
	}, nil
}

// sets the X-Original-Host header from the request's Host
func (m *OriginalHostMiddleware) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	originalHost := req.Header.Get(m.headerName)
	if originalHost == "" {
		// Fallback to X-Forwarded-Host
		originalHost = req.Header.Get("X-Forwarded-Host")
		if originalHost == "" {
			// Finally fallback to the current Host header
			originalHost = req.Host
		}
	}

	req.Header.Set(m.headerName, originalHost)
	m.next.ServeHTTP(rw, req)
}
