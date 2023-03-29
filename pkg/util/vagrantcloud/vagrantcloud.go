package vagrantcloud

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	baseURL = "https://vagrantcloud.com/api/v1"
)

var ErrAPIError = errors.New("vagrant API answered with non-ok response")

type VagrantCloud struct {
	accessToken string
}

func New(accessToken string) *VagrantCloud {
	return &VagrantCloud{accessToken: accessToken}
}

func (v *VagrantCloud) get(path string) ([]byte, error) {
	return v.request(
		"GET",
		path,
		"application/x-www-form-urlencoded",
		strings.NewReader(""),
	)
}

func (v *VagrantCloud) post(path string, params url.Values) ([]byte, error) {
	return v.request(
		"POST",
		path,
		"application/x-www-form-urlencoded",
		strings.NewReader(params.Encode()),
	)
}

func (v *VagrantCloud) put(path string, params url.Values) ([]byte, error) {
	return v.request(
		"PUT",
		path,
		"application/x-www-form-urlencoded",
		strings.NewReader(params.Encode()),
	)
}

func (v *VagrantCloud) delete(path string) ([]byte, error) {
	return v.request(
		"DELETE",
		path,
		"application/x-www-form-urlencoded",
		strings.NewReader(""),
	)
}

func (v *VagrantCloud) upload(path string, data io.Reader) ([]byte, error) {
	return v.request(
		"PUT",
		path,
		"multipart/form-data",
		data,
	)
}

func (v *VagrantCloud) request(method, path, contentType string, data io.Reader) ([]byte, error) {
	requestURI, err := url.ParseRequestURI(baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("invalid URI: %w", err)
	}

	if v.accessToken != "" {
		query := requestURI.Query()
		query.Set("access_token", v.accessToken)
		requestURI.RawQuery = query.Encode()
	}

	req, err := http.NewRequest(method, requestURI.String(), data)
	if err != nil {
		return nil, fmt.Errorf("error preparing request: %w", err)
	}

	req = req.WithContext(context.Background())
	req.Header.Set("Content-Type", contentType)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error during HTTP request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%w: %d", ErrAPIError, resp.StatusCode)
	}

	return body, nil
}
