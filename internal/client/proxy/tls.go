/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package proxy

import (
	"crypto/tls"
	"crypto/x509"
	"os"

	"github.com/pkg/errors"
)

type TLSConfig struct {
	Insecure           bool
	ClientAuthCertFile string
	ClientAuthKeyFile  string
	CaFile             string
	ClientAuthCert     string
	ClientAuthKey      string
	CaCert             string
}

func GetConnectorTLSConfig(config *TLSConfig) (*tls.Config, error) {
	if config == nil {
		return nil, errors.New("please provide TLS config")
	}

	//nolint:gosec // (ignore "G402")
	tlsConfig := tls.Config{InsecureSkipVerify: config.Insecure}

	if len(config.ClientAuthCertFile) > 0 {
		// cert and key are passed via filesystem
		if len(config.ClientAuthKeyFile) == 0 {
			return nil, errors.New("please provide key file for client certificate")
		}

		cert, err := tls.LoadX509KeyPair(config.ClientAuthCertFile, config.ClientAuthKeyFile)
		if err != nil {
			return nil, errors.Wrap(err, "failed to load client cert/key pair")
		}

		tlsConfig.GetClientCertificate = func(*tls.CertificateRequestInfo) (*tls.Certificate, error) {
			return &cert, nil
		}
	}

	if len(config.ClientAuthCert) > 0 {
		// cert and key are passed as strings
		if len(config.ClientAuthKey) == 0 {
			return nil, errors.New("please provide key for client certificate")
		}

		cert, err := tls.X509KeyPair([]byte(config.ClientAuthCert), []byte(config.ClientAuthKey))
		if err != nil {
			return nil, errors.Wrap(err, "failed to load client cert/key pair")
		}

		tlsConfig.GetClientCertificate = func(*tls.CertificateRequestInfo) (*tls.Certificate, error) {
			return &cert, nil
		}
	}

	if len(config.CaFile) > 0 {
		caCert, err := os.ReadFile(config.CaFile)
		if err != nil {
			return nil, errors.Wrap(err, "failed to read CA certificate from file")
		}

		caCertPool := x509.NewCertPool()

		ok := caCertPool.AppendCertsFromPEM(caCert)
		if !ok {
			return nil, errors.New("failed to append CA certificate from file")
		}

		tlsConfig.RootCAs = caCertPool
	}

	if len(config.CaCert) > 0 {
		caCertPool := x509.NewCertPool()

		ok := caCertPool.AppendCertsFromPEM([]byte(config.CaCert))
		if !ok {
			return nil, errors.New("failed to append CA certificate from field")
		}

		tlsConfig.RootCAs = caCertPool
	}

	return &tlsConfig, nil
}
