package proxy

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io/ioutil"
	"net/url"
	"path/filepath"

	"google.golang.org/grpc/credentials"
)

type Config struct {
	Address  string
	CAFile   string
	CertFile string
	KeyFile  string
}

func (c *Config) DialOptions() (string, string, error) {
	u, err := url.Parse(c.Address)
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

func (c *Config) TLSServerOptions() (credentials.TransportCredentials, error) {
	tlsConfig := &tls.Config{
		MinVersion:   tls.VersionTLS12,
		ClientAuth:   tls.RequireAndVerifyClientCert,
		Certificates: []tls.Certificate{},
		ClientCAs:    x509.NewCertPool(),
	}

	certificate, err := tls.LoadX509KeyPair(c.CertFile, c.KeyFile)
	if err != nil {
		return nil, fmt.Errorf("could not load certificate file: %w", err)
	}

	tlsConfig.Certificates = append(tlsConfig.Certificates, certificate)

	bs, err := ioutil.ReadFile(c.CAFile)
	if err != nil {
		return nil, fmt.Errorf("could not read ca file: %w", err)
	}

	if ok := tlsConfig.ClientCAs.AppendCertsFromPEM(bs); !ok {
		return nil, errors.New("invalid ca")
	}

	return credentials.NewTLS(tlsConfig), nil
}

func (c *Config) TLSClientOptions() (credentials.TransportCredentials, error) {
	tlsConfig := &tls.Config{
		MinVersion:   tls.VersionTLS12,
		Certificates: []tls.Certificate{},
		RootCAs:      x509.NewCertPool(),
	}

	u, err := url.Parse(c.Address)
	if err != nil {
		return nil, fmt.Errorf("invalid address: %w", err)
	}

	if u.Scheme == "tcp" {
		tlsConfig.ServerName = u.Hostname()
	}

	certificate, err := tls.LoadX509KeyPair(c.CertFile, c.KeyFile)
	if err != nil {
		return nil, fmt.Errorf("could not load certificate file: %w", err)
	}

	tlsConfig.Certificates = append(tlsConfig.Certificates, certificate)

	bs, err := ioutil.ReadFile(c.CAFile)
	if err != nil {
		return nil, fmt.Errorf("could not read ca file: %w", err)
	}

	if ok := tlsConfig.RootCAs.AppendCertsFromPEM(bs); !ok {
		return nil, errors.New("invalid ca")
	}

	return credentials.NewTLS(tlsConfig), nil
}
