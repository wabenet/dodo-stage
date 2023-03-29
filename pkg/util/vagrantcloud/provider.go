package vagrantcloud

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"time"
)

type Provider struct {
	Name        string    `json:"name"`
	Hosted      bool      `json:"hosted"`
	HostedToken string    `json:"hosted_token"` //nolint: tagliatelle
	OriginalURL string    `json:"original_url"` //nolint: tagliatelle
	UploadURL   string    `json:"upload_url"`   //nolint: tagliatelle
	CreatedAt   time.Time `json:"created_at"`   //nolint: tagliatelle
	UpdatedAt   time.Time `json:"updated_at"`   //nolint: tagliatelle
	DownloadURL string    `json:"download_url"` //nolint: tagliatelle
}

type ProviderOptions struct {
	Version *VersionOptions
	Name    string
	URL     string
}

func (p *ProviderOptions) toPath() string {
	return fmt.Sprintf("%s/provider/%s", p.Version.toPath(), p.Name)
}

func (p *ProviderOptions) toParams() url.Values {
	params := url.Values{}
	params.Add("provider[name]", p.Name)
	params.Add("provider[url]", p.URL)

	return params
}

func (v *VagrantCloud) GetProvider(opts *ProviderOptions) (*Provider, error) {
	body, err := v.get(opts.toPath())
	if err != nil {
		return nil, err
	}

	return parseProvider(body)
}

func (v *VagrantCloud) CreateProvider(opts *ProviderOptions) (*Provider, error) {
	body, err := v.post(opts.Version.toPath()+"/providers", opts.toParams())
	if err != nil {
		return nil, err
	}

	return parseProvider(body)
}

func (v *VagrantCloud) UpdateProvider(opts *ProviderOptions) (*Provider, error) {
	body, err := v.put(opts.toPath(), opts.toParams())
	if err != nil {
		return nil, err
	}

	return parseProvider(body)
}

func (v *VagrantCloud) DeleteProvider(opts *ProviderOptions) (*Provider, error) {
	body, err := v.delete(opts.toPath())
	if err != nil {
		return nil, err
	}

	return parseProvider(body)
}

func (v *VagrantCloud) UploadProvider(opts *ProviderOptions, data io.Reader) (*Provider, error) {
	body, err := v.upload(opts.toPath()+"/upload", data)
	if err != nil {
		return nil, err
	}

	return parseProvider(body)
}

func parseProvider(data []byte) (*Provider, error) {
	p := &Provider{}
	if err := json.Unmarshal(data, p); err != nil {
		return nil, fmt.Errorf("could not parse provider json: %w", err)
	}

	return p, nil
}
