package proxy

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io/ioutil"
	"net/url"
	"path/filepath"

	provision "github.com/wabenet/dodo-stage/api/provision/v1alpha1"
	"google.golang.org/grpc/credentials"
)

func DialOptions(c *provision.ProxyConfig) (string, string, error) {
	u, err := url.Parse(c.Url)
	if err != nil {
		return "", "", fmt.Errorf("invalid address: %w", err)
	}

	switch u.Scheme {
	case "tcp":
		return u.Scheme, u.Host, nil

	case "unix":
		addr, err := filepath.Abs(u.Host)
		if err != nil {
			return "", "", fmt.Errorf("could not get socket path: %w", err)
		}

		return u.Scheme, addr, nil
	}

	return "", "", errors.New("invalid protocol")
}

func TLSServerOptions(c *provision.ProxyConfig) (credentials.TransportCredentials, error) {
	tlsConfig := &tls.Config{
		MinVersion:   tls.VersionTLS12,
		ClientAuth:   tls.RequireAndVerifyClientCert,
		Certificates: []tls.Certificate{},
		ClientCAs:    x509.NewCertPool(),
	}

	certificate, err := tls.LoadX509KeyPair(c.CertPath, c.KeyPath)
	if err != nil {
		return nil, fmt.Errorf("could not load certificate file: %w", err)
	}

	tlsConfig.Certificates = append(tlsConfig.Certificates, certificate)

	bs, err := ioutil.ReadFile(c.CaPath)
	if err != nil {
		return nil, fmt.Errorf("could not read ca file: %w", err)
	}

	if ok := tlsConfig.ClientCAs.AppendCertsFromPEM(bs); !ok {
		return nil, errors.New("invalid ca")
	}

	return credentials.NewTLS(tlsConfig), nil
}

func TLSClientOptions(c *provision.ProxyConfig) (credentials.TransportCredentials, error) {
	tlsConfig := &tls.Config{
		MinVersion:   tls.VersionTLS12,
		Certificates: []tls.Certificate{},
		RootCAs:      x509.NewCertPool(),
	}

	u, err := url.Parse(c.Url)
	if err != nil {
		return nil, fmt.Errorf("invalid address: %w", err)
	}

	if u.Scheme == "tcp" {
		tlsConfig.ServerName = u.Hostname()
	}

	certificate, err := tls.LoadX509KeyPair(c.CertPath, c.KeyPath)
	if err != nil {
		return nil, fmt.Errorf("could not load certificate file: %w", err)
	}

	tlsConfig.Certificates = append(tlsConfig.Certificates, certificate)

	bs, err := ioutil.ReadFile(c.CaPath)
	if err != nil {
		return nil, fmt.Errorf("could not read ca file: %w", err)
	}

	if ok := tlsConfig.RootCAs.AppendCertsFromPEM(bs); !ok {
		return nil, errors.New("invalid ca")
	}

	return credentials.NewTLS(tlsConfig), nil
}
