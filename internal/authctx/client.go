/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package authctx

import (
	"context"
	"crypto/tls"
	"fmt"
	"strings"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	grpc_credentials "google.golang.org/grpc/credentials"
)

const (
	ServerEndpointEnvVar = "TMC_ENDPOINT"
	CSPTokenEnvVar       = "TMC_CSP_TOKEN"
)

type TanzuContext struct {
	ServerEndpoint string
	Token          string
	CSPEndPoint    string
	UserAuthCtx    context.Context
	TMCConnection  *grpc.ClientConn
}

func (cfg *TanzuContext) Setup() error {
	serverEndpoint := cfg.ServerEndpoint
	if !strings.HasSuffix(cfg.ServerEndpoint, "443") {
		serverEndpoint = fmt.Sprintf("%s:%d", cfg.ServerEndpoint, 443)
	}

	var err error

	cfg.UserAuthCtx, err = getUserAuthCtx(cfg)
	if err != nil {
		return errors.Wrap(err, "while getting user ctx")
	}

	tlsCreds := grpc_credentials.NewTLS(&tls.Config{})

	unaryInterceptors := []grpc.UnaryClientInterceptor{
		grpc_retry.UnaryClientInterceptor(),
	}

	dialOpts := []grpc.DialOption{
		grpc.WithTransportCredentials(tlsCreds),
		grpc.WithUnaryInterceptor(grpc_middleware.ChainUnaryClient(unaryInterceptors...)),
	}

	cfg.TMCConnection, err = grpc.Dial(serverEndpoint, dialOpts...)
	if err != nil {
		return errors.Wrap(err, "while dailing grpc connection")
	}

	return nil
}
