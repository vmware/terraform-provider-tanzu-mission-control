/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package transport

import (
	"bytes"
	"crypto/x509"
	"io"
	"net"
	"net/http"
	"time"

	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
)

// Client is the http client implementation.
type Client struct {
	*Config
	client     *http.Client
	timeout    time.Duration
	interval   time.Duration
	retryCount int
}

const (
	defaultRetryCount       = 3
	defaultHTTPTimeout      = 30 * time.Second
	defaultIntervalDuration = 1 * time.Second
)

// NewClient returns a new instance of http Client.
func NewClient(config *helper.TLSConfig) (*Client, error) {
	client := Client{
		Config:     DefaultTransportConfig(),
		timeout:    defaultHTTPTimeout,
		retryCount: defaultRetryCount,
		interval:   defaultIntervalDuration,
	}

	var transport *http.Transport

	// Setup HTTPS client.
	tlsConfig, err := helper.GetConnectorTLSConfig(config)
	if err != nil {
		return nil, err
	}

	if tlsConfig.RootCAs == nil {
		tlsConfig.RootCAs = x509.NewCertPool()
	}

	tlsConfig.RootCAs.AppendCertsFromPEM([]byte(tmcRootCA))

	transport = &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		Proxy:                 http.ProxyFromEnvironment,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		TLSClientConfig:       tlsConfig,
	}

	client.client = &http.Client{
		Timeout:   defaultHTTPTimeout,
		Transport: transport,
	}

	return &client, nil
}

// Get makes a HTTP GET request to provided URL.
func (c *Client) get(url string, headers http.Header) (*http.Response, error) {
	var response *http.Response

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return response, errors.Wrap(err, "GET - request creation failed")
	}

	request.Header = headers

	return c.Do(request)
}

// Post makes a HTTP POST request to provided URL and requestBody.
func (c *Client) post(url string, body io.Reader, headers http.Header) (*http.Response, error) {
	var response *http.Response

	request, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return response, errors.Wrap(err, "POST - request creation failed")
	}

	request.Header = headers

	return c.Do(request)
}

// Put makes a HTTP PUT request to provided URL and requestBody.
func (c *Client) put(url string, body io.Reader, headers http.Header) (*http.Response, error) {
	var response *http.Response

	request, err := http.NewRequest(http.MethodPut, url, body)
	if err != nil {
		return response, errors.Wrap(err, "PUT - request creation failed")
	}

	request.Header = headers

	return c.Do(request)
}

// Patch makes a HTTP PATCH request to provided URL and requestBody.
// nolint: unused
func (c *Client) patch(url string, body io.Reader, headers http.Header) (*http.Response, error) {
	var response *http.Response

	request, err := http.NewRequest(http.MethodPatch, url, body)
	if err != nil {
		return response, errors.Wrap(err, "PATCH - request creation failed")
	}

	request.Header = headers

	return c.Do(request)
}

// Delete makes a HTTP DELETE request with provided URL.
func (c *Client) delete(url string, headers http.Header) (*http.Response, error) {
	var response *http.Response

	request, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return response, errors.Wrap(err, "DELETE - request creation failed")
	}

	request.Header = headers

	return c.Do(request)
}

// Do makes an HTTP request with the native `http.Do` interface.
func (c *Client) Do(request *http.Request) (*http.Response, error) {
	request.Close = true

	var bodyReader *bytes.Reader

	if request.Body != nil {
		reqData, err := io.ReadAll(request.Body)
		if err != nil {
			return nil, err
		}

		bodyReader = bytes.NewReader(reqData)
		request.Body = io.NopCloser(bodyReader) // prevents closing the body between retries
	}

	var (
		err      error
		response *http.Response
	)

	for i := 0; i <= c.retryCount; i++ {
		if response != nil {
			response.Body.Close()
		}

		response, err = c.client.Do(request)

		if bodyReader != nil {
			// Reset the body reader after the request since at this point it's already read
			// Note that it's safe to ignore the error here since the 0,0 position is always valid
			_, _ = bodyReader.Seek(0, 0)
		}

		if err != nil {
			time.Sleep(c.interval)
			continue
		}

		if response.StatusCode >= http.StatusInternalServerError {
			time.Sleep(c.interval)
			continue
		}

		err = nil // Clear errors if any iteration succeeds

		break
	}

	return response, err
}
