package vagrantcloud

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"time"
)

type Box struct {
	CreatedAt           time.Time `json:"created_at"` //nolint: tagliatelle
	UpdatedAt           time.Time `json:"updated_at"` //nolint: tagliatelle
	Tag                 string    `json:"tag"`
	Name                string    `json:"name"`
	ShortDescription    string    `json:"short_description"`    //nolint: tagliatelle
	DescriptionHTML     string    `json:"description_html"`     //nolint: tagliatelle
	DescriptionMarkdown string    `json:"description_markdown"` //nolint: tagliatelle
	Username            string    `json:"username"`
	Private             bool      `json:"private"`
	CurrentVersion      Version   `json:"current_version"` //nolint: tagliatelle
	Versions            []Version `json:"versions"`
}

type BoxOptions struct {
	Username         string
	Name             string
	ShortDescription string
	Description      string
	IsPrivate        bool
}

func (b *BoxOptions) toPath() string {
	return fmt.Sprintf("/box/%s/%s", b.Username, b.Name)
}

func (b *BoxOptions) toParams() url.Values {
	params := url.Values{}
	params.Add("box[name]", b.Name)
	params.Add("box[username]", b.Username)
	params.Add("box[short_description]", b.ShortDescription)
	params.Add("box[description]", b.Description)
	params.Add("box[is_private]", strconv.FormatBool(b.IsPrivate))

	return params
}

func (v *VagrantCloud) GetBox(opts *BoxOptions) (*Box, error) {
	body, err := v.get(opts.toPath())
	if err != nil {
		return nil, err
	}

	return parseBox(body)
}

func (v *VagrantCloud) CreateBox(opts *BoxOptions) (*Box, error) {
	body, err := v.post("/boxes", opts.toParams())
	if err != nil {
		return nil, err
	}

	return parseBox(body)
}

func (v *VagrantCloud) UpdateBox(opts *BoxOptions) (*Box, error) {
	body, err := v.put(opts.toPath(), opts.toParams())
	if err != nil {
		return nil, err
	}

	return parseBox(body)
}

func (v *VagrantCloud) DeleteBox(opts *BoxOptions) (*Box, error) {
	body, err := v.delete(opts.toPath())
	if err != nil {
		return nil, err
	}

	return parseBox(body)
}

func parseBox(data []byte) (*Box, error) {
	box := &Box{}
	if err := json.Unmarshal(data, box); err != nil {
		return nil, fmt.Errorf("could not parse box json: %w", err)
	}

	return box, nil
}
