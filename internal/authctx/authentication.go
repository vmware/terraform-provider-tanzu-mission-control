/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package authctx

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
)

type tokenResponse struct {
	IDToken      string `json:"id_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	Scope        string `json:"scope"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type AgentTokenInfo struct {
	AccessToken string `json:"access_token"`
}

func getBearerToken(cspEndpoint, cspToken string, config *helper.TLSConfig) (string, error) {
	var (
		transport *http.Transport
		resp      *http.Response
		err       error
	)

	tlsConfig, err := helper.GetConnectorTLSConfig(config)
	if err != nil {
		return "", err
	}

	transport = &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		Proxy:               http.ProxyFromEnvironment,
		MaxIdleConns:        1000,
		MaxIdleConnsPerHost: 200,
		IdleConnTimeout:     90 * time.Second,
		TLSClientConfig:     tlsConfig,
	}

	client := &http.Client{Transport: transport, Timeout: 60 * time.Second}

	data := url.Values{}
	data.Set("refresh_token", cspToken)
	encodedToken := strings.NewReader(data.Encode())

	for i := 0; i < 10; i++ {
		resp, err = client.Post(
			fmt.Sprintf("https://%s/csp/gateway/am/api/auth/api-tokens/authorize", cspEndpoint),
			"application/x-www-form-urlencoded",
			encodedToken,
		)

		if err == nil {
			defer resp.Body.Close()
			break
		}

		// retry for issue of go resolver returning AAAA records
		if urlErr, ok := err.(*url.Error); ok {
			if netErr, ok := urlErr.Err.(*net.OpError); ok {
				if osErr, ok := netErr.Err.(*os.SyscallError); ok {
					if osErr.Syscall == "connect" {
						continue
					}
				}
			}
		}

		return "", err
	}

	if err != nil {
		return "", err
	}

	respJSON, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	token := &tokenResponse{}

	err = json.Unmarshal(respJSON, token)
	if err != nil {
		return "", err
	}

	return token.AccessToken, nil
}

func getUserAuthCtx(config *TanzuContext) (map[string]string, error) {
	var (
		token string
		err   error
	)

	for i := 0; i < 3; i++ {
		token, err = getBearerToken(config.VMWCloudEndPoint, config.Token, config.TLSConfig)
		if err == nil {
			break
		}

		time.Sleep(10 * time.Second)
	}

	if err != nil {
		return nil, errors.Wrap(err, "while getting bearer token from VMware Cloud API Token")
	}

	md := map[string]string{
		"authorization": "Bearer " + token,
	}

	return md, nil
}
