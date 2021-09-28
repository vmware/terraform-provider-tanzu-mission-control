/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package transport

import (
	"fmt"
	"net/http"
	"strings"
)

const (
	// DefaultHost is the default Host.
	DefaultHost string = "https://<your_org>.tmc.cloud.vmware.com"
	// DefaultBasePath is the default BasePath.
	DefaultBasePath string = "/"

	contentLengthKey = "Content-Length"
)

// DefaultTransportConfig creates a Config with default values.
func DefaultTransportConfig() *Config {
	return &Config{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Headers:  http.Header{},
	}
}

// Config contains the transport related info.
type Config struct {
	Host     string
	BasePath string
	Headers  http.Header
}

// WithHost overrides the default host.
func (cfg *Config) WithHost(host string) *Config {
	host = strings.TrimPrefix(host, "http://")
	host = strings.TrimSuffix(host, "/")

	if !strings.HasPrefix(host, "https://") {
		host = fmt.Sprintf("%s%s", "https://", host)
	}

	cfg.Host = host

	return cfg
}

// WithBasePath overrides the default basePath.
func (cfg *Config) WithBasePath(basePath string) *Config {
	cfg.BasePath = basePath
	return cfg
}

// WithHeader overrides the default header.
func (cfg *Config) AddHeaders(header http.Header) *Config {
	if cfg.Headers == nil {
		cfg.Headers = header
	}

	for key, value := range header {
		cfg.Headers[key] = value
	}

	return cfg
}
