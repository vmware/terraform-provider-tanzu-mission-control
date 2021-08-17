/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package authctx

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/pkg/errors"
	"google.golang.org/grpc/metadata"
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

func getBearerToken(cspEndpoint, cspToken string) (string, error) {
	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConns:        1000,
		MaxIdleConnsPerHost: 200,
		IdleConnTimeout:     90 * time.Second,
	}

	client := &http.Client{Transport: transport, Timeout: 60 * time.Second}

	var (
		resp *http.Response
		err  error
	)

	for i := 0; i < 10; i++ {
		resp, err = client.Post(
			fmt.Sprintf(
				"https://%s/csp/gateway/am/api/auth/api-tokens/authorize?refresh_token=%s",
				cspEndpoint,
				cspToken,
			),
			"application/json",
			nil,
		)
		if err == nil {
			defer resp.Body.Close()
			break
		}

		// retry for issue of go resolver returning AAAA records
		// https://gitlab.eng.vmware.com/cnabu-sre/triage/issues/286
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

	respJSON, err := ioutil.ReadAll(resp.Body)
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

func getUserAuthCtx(config *TanzuContext) (context.Context, error) {
	var (
		token string
		err   error
	)

	for i := 0; i < 3; i++ {
		token, err = getBearerToken(config.CSPEndPoint, config.Token)
		if err == nil {
			break
		}

		time.Sleep(10 * time.Second)
	}

	if err != nil {
		return nil, errors.Wrap(err, "while getting bearer token")
	}

	endpoint := strings.TrimSuffix(config.ServerEndpoint, ":443")

	md := metadata.MD{
		"x-customer-domain": []string{endpoint},
		"authorization":     []string{"Bearer " + token},
	}

	return metadata.NewOutgoingContext(context.Background(), md), nil
}

func (cfg *TanzuContext) RefreshUserBearerToken() error {
	authCtx, err := getUserAuthCtx(cfg)
	if err != nil {
		return err
	}

	cfg.UserAuthCtx = authCtx

	return nil
}
