package vagrantcloud

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

type Status string

const (
	Unreleased Status = "unreleased"
	Active     Status = "active"
	Revoked    Status = "revoked"
)

type Version struct {
	Version             string     `json:"version"`
	Status              Status     `json:"status"`
	DescriptionHTML     string     `json:"description_html"`     //nolint: tagliatelle
	DescriptionMarkdown string     `json:"description_markdown"` //nolint: tagliatelle
	CreatedAt           time.Time  `json:"created_at"`           //nolint: tagliatelle
	UpdatedAt           time.Time  `json:"updated_at"`           //nolint: tagliatelle
	Number              string     `json:"number"`
	Downloads           int        `json:"downloads"`
	ReleaseURL          string     `json:"release_url"` //nolint: tagliatelle
	RevokeURL           string     `json:"revoke_url"`  //nolint: tagliatelle
	Providers           []Provider `json:"providers"`
}

type VersionOptions struct {
	Box         *BoxOptions
	Number      string
	Version     string
	Description string
}

func (v *VersionOptions) toPath() string {
	return fmt.Sprintf("%s/version/%s", v.Box.toPath(), v.Number)
}

func (v *VersionOptions) toParams() url.Values {
	params := url.Values{}
	params.Add("version[version]", v.Version)
	params.Add("version[description]", v.Description)

	return params
}

func (v *VagrantCloud) GetVersion(opts *VersionOptions) (*Version, error) {
	body, err := v.get(opts.toPath())
	if err != nil {
		return nil, err
	}

	return parseVersion(body)
}

func (v *VagrantCloud) CreateVersion(opts *VersionOptions) (*Version, error) {
	body, err := v.post(opts.toPath()+"/versions", opts.toParams())
	if err != nil {
		return nil, err
	}

	return parseVersion(body)
}

func (v *VagrantCloud) UpdateVersion(opts *VersionOptions) (*Version, error) {
	body, err := v.put(opts.toPath(), opts.toParams())
	if err != nil {
		return nil, err
	}

	return parseVersion(body)
}

func (v *VagrantCloud) DeleteVersion(opts *VersionOptions) (*Version, error) {
	body, err := v.delete(opts.toPath())
	if err != nil {
		return nil, err
	}

	return parseVersion(body)
}

func (v *VagrantCloud) Release(opts *VersionOptions) (*Version, error) {
	body, err := v.put(opts.toPath()+"/release", url.Values{})
	if err != nil {
		return nil, err
	}

	return parseVersion(body)
}

func (v *VagrantCloud) Revoke(opts *VersionOptions) (*Version, error) {
	body, err := v.put(opts.toPath()+"/revoke", url.Values{})
	if err != nil {
		return nil, err
	}

	return parseVersion(body)
}

func parseVersion(data []byte) (*Version, error) {
	v := &Version{}
	if err := json.Unmarshal(data, v); err != nil {
		return nil, fmt.Errorf("could not parse version json: %w", err)
	}

	return v, nil
}
