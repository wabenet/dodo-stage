package provision

import (
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"
	log "github.com/hashicorp/go-hclog"
	core "github.com/wabenet/dodo-core/api/core/v1alpha5"
	"github.com/wabenet/dodo-core/pkg/plugin"
	provision "github.com/wabenet/dodo-stage/api/provision/v1alpha2"
	stage "github.com/wabenet/dodo-stage/api/stage/v1alpha4"
	"golang.org/x/net/context"
)

var _ Provisioner = &client{}

type client struct {
	provisionClient provision.PluginClient
}

func (c *client) Type() plugin.Type {
	return Type
}

func (c *client) PluginInfo() *core.PluginInfo {
	info, err := c.provisionClient.GetPluginInfo(context.Background(), &empty.Empty{})
	if err != nil {
		return &core.PluginInfo{
			Name:   &core.PluginName{Type: Type.String(), Name: plugin.FailedPlugin},
			Fields: map[string]string{"error": err.Error()},
		}
	}

	return &core.PluginInfo{
		Name: &core.PluginName{Name: info.Name.Name, Type: info.Name.Type},
	}
}

func (c *client) Init() (plugin.Config, error) {
	resp, err := c.provisionClient.InitPlugin(context.Background(), &empty.Empty{})
	if err != nil {
		return nil, fmt.Errorf("could not initialize plugin: %w", err)
	}

	return resp.Config, nil
}

func (c *client) Cleanup() {
	_, err := c.provisionClient.ResetPlugin(context.Background(), &empty.Empty{})
	if err != nil {
		log.L().Error("plugin reset error", "error", err)
	}
}

func (c *client) ProvisionStage(name string, sshopts *stage.SSHOptions) error {
	if _, err := c.provisionClient.ProvisionStage(context.Background(), &provision.ProvisionStageRequest{
		Name:       name,
		SshOptions: sshopts,
	}); err != nil {
		return fmt.Errorf("could not provision stage: %w", err)
	}

	return nil
}

func (c *client) CleanStage(name string) error {
	if _, err := c.provisionClient.CleanStage(context.Background(), &provision.CleanStageRequest{
		Name: name,
	}); err != nil {
		return fmt.Errorf("could not cleanup stage: %w", err)
	}

	return nil
}
