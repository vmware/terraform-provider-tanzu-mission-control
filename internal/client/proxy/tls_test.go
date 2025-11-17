// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package proxy

import (
	"crypto/tls"
	"crypto/x509"
	"os"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

const (
	clientCertPath                 = "./testdata/client.crt"
	clientKeyPath                  = "./testdata/client.key"
	caCertPath                     = "./testdata/ca.crt"
	invalidClientCertContentInPath = "./testdata/invalid_client.crt"
	invalidClientKeyContentInPath  = "./testdata/invalid_client.key"
	invalidCaCertContentInPath     = "./testdata/invalid_ca.crt"
	clientCert                     = `
-----BEGIN CERTIFICATE-----
MIIEDzCCAvegAwIBAgIJAJZD62ElH7OTMA0GCSqGSIb3DQEBCwUAMGExCzAJBgNV
BAYTAlVTMQ0wCwYDVQQIDAR0ZXN0MQ0wCwYDVQQHDAR0ZXN0MQ4wDAYDVQQKDAVz
cXVpZDEOMAwGA1UECwwFc3F1aWQxFDASBgNVBAMMC3NxdWlkLmxvY2FsMB4XDTIy
MTAyNzA4MTYxMFoXDTI1MDEyOTA4MTYxMFowYTELMAkGA1UEBhMCVVMxDTALBgNV
BAgMBHRlc3QxDTALBgNVBAcMBHRlc3QxDjAMBgNVBAoMBXNxdWlkMQ4wDAYDVQQL
DAVzcXVpZDEUMBIGA1UEAwwLc3F1aWQubG9jYWwwggEiMA0GCSqGSIb3DQEBAQUA
A4IBDwAwggEKAoIBAQDnWYafBZLYclPNt8aAAtOAHPr04wSA8aWOt+/nyU5TnRX5
tTZOuSrmUZdu++cbWZGw2TokmQBFMTWiOyL0Qf9JIr9GxWL/366NN7+y5RiFeyzG
FOrJKRIKd6AQP8/amO00dlRSyq5GmNgn1uCFEfRx28InQMCsHY5ihlWEQsIhjEEO
6F8vM/j4C7RWYH681kJb/lihpVMje+B4bmxM4dfT9mZ1Xv3uvOCHxqlEB97IL2Md
8MAmxbHAvLs0BesrKvE9vOvBiS+1YtU4cEcK52YIzoa2oygqIa+Glbyz8Z6VqgEj
WaHHF7Itj3uKOLTxKaFWIiySemyY5OONtMagH+vBAgMBAAGjgckwgcYwDwYDVR0T
AQH/BAUwAwEB/zAdBgNVHQ4EFgQUC84qNkHPfbiry5TRzrQ0cJZ8wd8wgZMGA1Ud
IwSBizCBiIAUC84qNkHPfbiry5TRzrQ0cJZ8wd+hZaRjMGExCzAJBgNVBAYTAlVT
MQ0wCwYDVQQIDAR0ZXN0MQ0wCwYDVQQHDAR0ZXN0MQ4wDAYDVQQKDAVzcXVpZDEO
MAwGA1UECwwFc3F1aWQxFDASBgNVBAMMC3NxdWlkLmxvY2FsggkAlkPrYSUfs5Mw
DQYJKoZIhvcNAQELBQADggEBAEZ7Ei4A3oQLbUL8xhIUvJXiogJtL1gwVEEt3wtQ
y5ZP+qe3HEXd/F52VbdTEiFtMBa5/nEtu+Bo0OF4fkMkbkzVAkVUvkmBA6BQRLP7
TPbjjT+018MGZCXGNJezUr8yt+By9jeAZE9di17HJ8IAOoK8aY9P1N2BNOoh/zBc
J+guVb41E7ckHez4ENTEj2hrqrYifViGvOaSAG3w1d8PW+wIj3jI6vQk7vTO6mbt
+A42x+js/D4vSO2J/RZkFsfjPDjCmQKFOR7xH+5S64Hg+973soUhYgQbC8h6Kytd
ws3Up/6R4aI0ohB4wjOfUCL6x9L3pyQcDoumDuy8ToyL5Pk=
-----END CERTIFICATE-----
`
	// nolint: gosec
	clientKey = `
-----BEGIN RSA PRIVATE KEY-----
MIIEpQIBAAKCAQEA51mGnwWS2HJTzbfGgALTgBz69OMEgPGljrfv58lOU50V+bU2
Trkq5lGXbvvnG1mRsNk6JJkARTE1ojsi9EH/SSK/RsVi/9+ujTe/suUYhXssxhTq
ySkSCnegED/P2pjtNHZUUsquRpjYJ9bghRH0cdvCJ0DArB2OYoZVhELCIYxBDuhf
LzP4+Au0VmB+vNZCW/5YoaVTI3vgeG5sTOHX0/ZmdV797rzgh8apRAfeyC9jHfDA
JsWxwLy7NAXrKyrxPbzrwYkvtWLVOHBHCudmCM6GtqMoKiGvhpW8s/GelaoBI1mh
xxeyLY97iji08SmhViIsknpsmOTjjbTGoB/rwQIDAQABAoIBAQC5J8KBVPGgv2do
1nNMknZ5KziCc4pwNHX/EiE0Tb2dV7R1xF0xhiraaGSdU4g0MGxUsJZDIhzAQ9Ec
BN5FxOguvVt+pY1FId1Ocla/M3F7qwg7hBQgaPliXTlCp/+PgSMfALEeDc6K+9rL
j8LjXWpZEbuCtOU/GuyQ19hrYQ6Dz3naarBjLw8K3Hgb5mYZfKTGZzQiZo7tVpwS
hjFibrx+SFAikDKx9W0byFIL8/LlbN6fagoOolswhuH7MDoojSdzq3+tKHsdlnWR
zUYDLc9cLwQR7YuAhbB6biSfzxxQMd90ORHnrLZfvmLbLgfv/9XKnktkALMS1O4Q
kTdplliRAoGBAPNqIVEWMXjaeOhanDOg1+Im1c/VD1NtwzoNmA2/DrvtLjXQr2C0
CzWIx8jRVWPLtX9D5xouQPoj/TeTgul0i09ImPEFbjeqo3GAFt09nEZNoNYOWdof
x6Kq9cGmeSrziBQQTLbi8XU6HJy9zSjgEPp/COrAxy7Qofdc9ES4YqMNAoGBAPNP
tBllsaFZUi5GZpGFARBWwetAKJrohvjJ/bkm+NtnAAiRb0jhJDk+kLMaxURjbUXo
zbhgrmjQL5kXmJZwx1Am+mOPa92uRPIL2E8HYUynUR9/12uZEFdY5p/Tq0x3/wLM
U23WhjF9Hv8j9PMGpdGMpvgyTygJNGDletC8MI6FAoGBAIBtlJqFzSBolLZzaErN
KFpIBzOaxHVOSl0M4xcNoSaCI4l9S6sIE4nxWweXeygmSOKW3w4vLVVNO8Lg74dh
WEdClH9GUDrKq2WtIWMlqJhnSN7nv3yYm+o1rWi4/uEskLWVTASKhL9HI+WHNwHE
BvFDqV7Cy9Tley9aOf8wEcrZAoGAMcc8sXV8weXkSlNc6KitbwpQ4jBeHlM0SfIJ
VrgCceDAwQAAJIjrQErsj7gKY9Nzp7nZXL9q70aodkm9jgnEvUE8OOI+zzu4H00N
FB4OagBROICPMhQ+o6AsjsZfZWWnZosnBnG9QqK2lLxmgNH7WsPL5TtltmsrrCdG
2S2nQYECgYEA6SGqvE5TG81wqApUzHcQPwc13W7F8x4uUO3wsNfCS0tr3IhVoTsH
aFv4vWFNv7roN8VUe3WCQGDFQ+3yKi7OwfVDHjJT9Gc8ejfNIYxLuv4pqD/WMlKg
xWu1m8S7h1lyYlqcviWhzNO7qk+W5hRfGqyq27UW3C+08rKJ3ygjWcc=
-----END RSA PRIVATE KEY-----
`
	caCert = `
-----BEGIN CERTIFICATE-----
MIIEDzCCAvegAwIBAgIJAJZD62ElH7OTMA0GCSqGSIb3DQEBCwUAMGExCzAJBgNV
BAYTAlVTMQ0wCwYDVQQIDAR0ZXN0MQ0wCwYDVQQHDAR0ZXN0MQ4wDAYDVQQKDAVz
cXVpZDEOMAwGA1UECwwFc3F1aWQxFDASBgNVBAMMC3NxdWlkLmxvY2FsMB4XDTIy
MTAyNzA4MTYxMFoXDTI1MDEyOTA4MTYxMFowYTELMAkGA1UEBhMCVVMxDTALBgNV
BAgMBHRlc3QxDTALBgNVBAcMBHRlc3QxDjAMBgNVBAoMBXNxdWlkMQ4wDAYDVQQL
DAVzcXVpZDEUMBIGA1UEAwwLc3F1aWQubG9jYWwwggEiMA0GCSqGSIb3DQEBAQUA
A4IBDwAwggEKAoIBAQDnWYafBZLYclPNt8aAAtOAHPr04wSA8aWOt+/nyU5TnRX5
tTZOuSrmUZdu++cbWZGw2TokmQBFMTWiOyL0Qf9JIr9GxWL/366NN7+y5RiFeyzG
FOrJKRIKd6AQP8/amO00dlRSyq5GmNgn1uCFEfRx28InQMCsHY5ihlWEQsIhjEEO
6F8vM/j4C7RWYH681kJb/lihpVMje+B4bmxM4dfT9mZ1Xv3uvOCHxqlEB97IL2Md
8MAmxbHAvLs0BesrKvE9vOvBiS+1YtU4cEcK52YIzoa2oygqIa+Glbyz8Z6VqgEj
WaHHF7Itj3uKOLTxKaFWIiySemyY5OONtMagH+vBAgMBAAGjgckwgcYwDwYDVR0T
AQH/BAUwAwEB/zAdBgNVHQ4EFgQUC84qNkHPfbiry5TRzrQ0cJZ8wd8wgZMGA1Ud
IwSBizCBiIAUC84qNkHPfbiry5TRzrQ0cJZ8wd+hZaRjMGExCzAJBgNVBAYTAlVT
MQ0wCwYDVQQIDAR0ZXN0MQ0wCwYDVQQHDAR0ZXN0MQ4wDAYDVQQKDAVzcXVpZDEO
MAwGA1UECwwFc3F1aWQxFDASBgNVBAMMC3NxdWlkLmxvY2FsggkAlkPrYSUfs5Mw
DQYJKoZIhvcNAQELBQADggEBAEZ7Ei4A3oQLbUL8xhIUvJXiogJtL1gwVEEt3wtQ
y5ZP+qe3HEXd/F52VbdTEiFtMBa5/nEtu+Bo0OF4fkMkbkzVAkVUvkmBA6BQRLP7
TPbjjT+018MGZCXGNJezUr8yt+By9jeAZE9di17HJ8IAOoK8aY9P1N2BNOoh/zBc
J+guVb41E7ckHez4ENTEj2hrqrYifViGvOaSAG3w1d8PW+wIj3jI6vQk7vTO6mbt
+A42x+js/D4vSO2J/RZkFsfjPDjCmQKFOR7xH+5S64Hg+973soUhYgQbC8h6Kytd
ws3Up/6R4aI0ohB4wjOfUCL6x9L3pyQcDoumDuy8ToyL5Pk=
-----END CERTIFICATE-----
`
)

func TestGetConnectorTLSConfig(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name              string
		config            *TLSConfig
		expectedTLSConfig *tls.Config
		expectedErr       error
	}{
		{
			name:   "",
			config: &TLSConfig{},
			//nolint:gosec // (ignore "G402")
			expectedTLSConfig: &tls.Config{},
			expectedErr:       nil,
		},
		{
			name:              "nil TLS config",
			config:            nil,
			expectedTLSConfig: nil,
			expectedErr:       errors.New("please provide TLS config"),
		},
		{
			name:   "empty TLS config",
			config: &TLSConfig{},
			//nolint:gosec // (ignore "G402")
			expectedTLSConfig: &tls.Config{InsecureSkipVerify: false},
			expectedErr:       nil,
		},
		{
			name:   "insecure allow unverified ssl set in TLS config",
			config: &TLSConfig{Insecure: true},
			//nolint:gosec // (ignore "G402")
			expectedTLSConfig: &tls.Config{InsecureSkipVerify: true},
			expectedErr:       nil,
		},
		{
			name: "key file for client certificate not set in TLS config",
			config: &TLSConfig{
				ClientAuthCertFile: clientCertPath,
			},
			expectedTLSConfig: nil,
			expectedErr:       errors.New("please provide key file for client certificate"),
		},
		{
			name: "wrong path to auth files in TLS config",
			config: &TLSConfig{
				ClientAuthCertFile: " ",
				ClientAuthKeyFile:  " ",
			},
			expectedTLSConfig: nil,
			expectedErr:       errors.New("failed to load client cert/key pair: open  : no such file or directory"),
		},
		{
			name: "wrong content in path to auth files in TLS config",
			config: &TLSConfig{
				ClientAuthCert: invalidClientCertContentInPath,
				ClientAuthKey:  invalidClientKeyContentInPath,
			},
			expectedTLSConfig: nil,
			expectedErr:       errors.New("failed to load client cert/key pair: tls: failed to find any PEM data in certificate input"),
		},
		{
			name: "everything set correctly for client auth files in TLS config",
			config: &TLSConfig{
				ClientAuthCertFile: clientCertPath,
				ClientAuthKeyFile:  clientKeyPath,
			},
			//nolint:gosec // (ignore "G402")
			expectedTLSConfig: &tls.Config{
				InsecureSkipVerify: false,
				GetClientCertificate: func(*tls.CertificateRequestInfo) (*tls.Certificate, error) {
					cert, _ := tls.LoadX509KeyPair(clientCertPath, clientKeyPath)
					return &cert, nil
				},
			},
			expectedErr: nil,
		},
		{
			name: "key field for client certificate not set in TLS config",
			config: &TLSConfig{
				ClientAuthCert: clientCert,
			},
			expectedTLSConfig: nil,
			expectedErr:       errors.New("please provide key for client certificate"),
		},
		{
			name: "wrong string as auth fields in TLS config",
			config: &TLSConfig{
				ClientAuthCert: " ",
				ClientAuthKey:  " ",
			},
			expectedTLSConfig: nil,
			expectedErr:       errors.New("failed to load client cert/key pair: tls: failed to find any PEM data in certificate input"),
		},
		{
			name: "everything set correctly for client auth fields in TLS config",
			config: &TLSConfig{
				ClientAuthCert: clientCert,
				ClientAuthKey:  clientKey,
			},
			//nolint:gosec // (ignore "G402")
			expectedTLSConfig: &tls.Config{
				InsecureSkipVerify: false,
				GetClientCertificate: func(*tls.CertificateRequestInfo) (*tls.Certificate, error) {
					cert, _ := tls.X509KeyPair([]byte(clientCert), []byte(clientKey))
					return &cert, nil
				},
			},
			expectedErr: nil,
		},
		{
			name: "wrong path to ca file in TLS config",
			config: &TLSConfig{
				CaFile: " ",
			},
			expectedTLSConfig: nil,
			expectedErr:       errors.New("failed to read CA certificate from file: open  : no such file or directory"),
		},
		{
			name: "wrong content for ca file in TLS config",
			config: &TLSConfig{
				CaFile: invalidCaCertContentInPath,
			},
			expectedTLSConfig: nil,
			expectedErr:       errors.New("failed to append CA certificate from file"),
		},
		{
			name: "correct path to ca file in TLS config",
			config: &TLSConfig{
				CaFile: caCertPath,
			},
			//nolint:gosec // (ignore "G402")
			expectedTLSConfig: &tls.Config{
				InsecureSkipVerify: false,
				RootCAs: func() *x509.CertPool {
					caCert1, _ := os.ReadFile(clientCertPath)
					caCertPool := x509.NewCertPool()
					caCertPool.AppendCertsFromPEM(caCert1)

					return caCertPool
				}(),
			},
			expectedErr: nil,
		},
		{
			name: "wrong content for ca cert in TLS config",
			config: &TLSConfig{
				CaCert: " ",
			},
			expectedTLSConfig: nil,
			expectedErr:       errors.New("failed to append CA certificate from field"),
		},
		{
			name: "correct ca cert in TLS config",
			config: &TLSConfig{
				CaCert: caCert,
			},
			//nolint:gosec // (ignore "G402")
			expectedTLSConfig: &tls.Config{
				InsecureSkipVerify: false,
				RootCAs: func() *x509.CertPool {
					caCertPool := x509.NewCertPool()
					caCertPool.AppendCertsFromPEM([]byte(caCert))

					return caCertPool
				}(),
			},
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actualTLSConfig, actualErr := GetConnectorTLSConfig(tc.config)
			if tc.expectedErr != nil {
				require.EqualError(t, actualErr, tc.expectedErr.Error())
				return
			}

			require.NoError(t, actualErr)

			if tc.expectedTLSConfig.GetClientCertificate != nil {
				require.NotNil(t, actualTLSConfig.GetClientCertificate)
				return
			}

			if tc.expectedTLSConfig.RootCAs != nil {
				require.NotNil(t, actualTLSConfig.RootCAs)
				return
			}

			require.Equal(t, actualTLSConfig, tc.expectedTLSConfig)
		})
	}
}
