package command

import (
	"fmt"

	log "github.com/hashicorp/go-hclog"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	core "github.com/wabenet/dodo-core/pkg/config"
	"github.com/wabenet/dodo-core/pkg/plugin"
	api "github.com/wabenet/dodo-stage/api/stage/v1alpha4"
	"github.com/wabenet/dodo-stage/internal/plugin/command/config"
	"github.com/wabenet/dodo-stage/pkg/plugin/provision"
	"github.com/wabenet/dodo-stage/pkg/plugin/stage"
	"github.com/wabenet/dodo-stage/pkg/util/ssh"
)

func New(m plugin.Manager) *Command {
	cmd := &cobra.Command{
		Use:              "stage",
		Short:            "Manage stages",
		TraverseChildren: true,
		SilenceUsage:     true,
	}

	cmd.AddCommand(NewListCommand(m))
	cmd.AddCommand(NewUpCommand(m))
	cmd.AddCommand(NewDownCommand(m))
	cmd.AddCommand(NewProvisionCommand(m))
	cmd.AddCommand(NewSSHCommand(m))

	return &Command{cmd: cmd}
}

func NewListCommand(m plugin.Manager) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List stages",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			stages, err := config.GetAllStages(core.GetConfigFiles()...)
			if err != nil {
				log.L().Error(err.Error())
			}

			for name, conf := range stages {
				s, err := loadStagePlugin(m, conf.Type)
				if err != nil {
					return err
				}

				current, err := s.GetStage(name)
				if err != nil {
					return err
				}

				fmt.Printf("%s (%v)", name, current.Info.Status)
			}

			return nil
		},
	}
}

func NewUpCommand(m plugin.Manager) *cobra.Command {
	return &cobra.Command{
		Use:   "up",
		Short: "Create or start a stage",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			conf, err := loadStageConfig(args[0])
			if err != nil {
				return err
			}

			s, err := loadStagePlugin(m, conf.Type)
			if err != nil {
				return err
			}

			p, err := loadProvisionPlugin(m, conf.Provision.Type)
			if err != nil {
				return err
			}

			current, err := s.GetStage(args[0])
			if err != nil {
				return err
			}

			if current.Info.Status == api.StageStatus_NONE {
				if err := s.CreateStage(args[0]); err != nil {
					return err
				}
			}

			if err := s.StartStage(args[0]); err != nil {
				return err
			}

			status, err := s.GetStage(args[0])
			if err != nil {
				return err
			}

			if err := p.ProvisionStage(args[0], status.SshOptions); err != nil {
				return err
			}

			return nil
		},
	}
}

type downOptions struct {
	remove  bool
	volumes bool
	force   bool
}

func NewDownCommand(m plugin.Manager) *cobra.Command {
	var opts downOptions
	cmd := &cobra.Command{
		Use:   "down",
		Short: "Remove or pause a stage",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			conf, err := loadStageConfig(args[0])
			if err != nil {
				return err
			}

			s, err := loadStagePlugin(m, conf.Type)
			if err != nil {
				return err
			}

			if opts.remove {
				return s.DeleteStage(args[0], opts.force, opts.volumes)
			}

			return s.StopStage(args[0])
		},
	}
	flags := cmd.Flags()
	flags.BoolVarP(&opts.remove, "rm", "", false, "remove the stage instead of pausing")
	flags.BoolVarP(&opts.volumes, "volumes", "", false, "when used with '--rm', also delete persistent volumes")
	flags.BoolVarP(&opts.force, "force", "f", false, "when used with '--rm', don't stop on errors")

	return cmd
}

func NewProvisionCommand(m plugin.Manager) *cobra.Command {
	return &cobra.Command{
		Use:   "provision",
		Short: "provision a stage",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			conf, err := loadStageConfig(args[0])
			if err != nil {
				return err
			}

			s, err := loadStagePlugin(m, conf.Type)
			if err != nil {
				return err
			}

			p, err := loadProvisionPlugin(m, conf.Provision.Type)
			if err != nil {
				return err
			}

			current, err := s.GetStage(args[0])
			if err != nil {
				return err
			}

			if current.Info.Status != api.StageStatus_UP {
				return errors.New("stage is not up")
			}

			status, err := s.GetStage(args[0])
			if err != nil {
				return err
			}

			if err := p.ProvisionStage(args[0], status.SshOptions); err != nil {
				return err
			}

			return nil
		},
	}
}

func NewSSHCommand(m plugin.Manager) *cobra.Command {
	return &cobra.Command{
		Use:   "ssh",
		Short: "login to the stage",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			conf, err := loadStageConfig(args[0])
			if err != nil {
				return err
			}

			s, err := loadStagePlugin(m, conf.Type)
			if err != nil {
				return err
			}

			current, err := s.GetStage(args[0])
			if err != nil {
				return err
			}

			if current.Info.Status != api.StageStatus_UP {
				return errors.New("stage is not up")
			}

			return ssh.Shell(current.SshOptions)
		},
	}
}

func loadStageConfig(name string) (*config.Stage, error) {
	stages, err := config.GetAllStages(core.GetConfigFiles()...)
	if err != nil {
		log.L().Error(err.Error())
	}

	if conf, ok := stages[name]; ok {
		conf.Name = name // TODO: figure out where to set defaults like this

		return conf, nil
	}

	return nil, fmt.Errorf("could not find any configuration for stage '%s'", name)
}

func loadStagePlugin(m plugin.Manager, name string) (stage.Stage, error) {
	for _, p := range m.GetPlugins(stage.Type.String()) {
		s := p.(stage.Stage)
		if info := s.PluginInfo(); info.Name.Name == name {
			return s, nil
		}
	}

	return nil, fmt.Errorf("could not find any stage plugin for type '%s'", name)
}

func loadProvisionPlugin(m plugin.Manager, name string) (provision.Provisioner, error) {
	for _, p := range m.GetPlugins(provision.Type.String()) {
		s := p.(provision.Provisioner)
		if info := s.PluginInfo(); info.Name.Name == name {
			return s, nil
		}
	}

	return nil, fmt.Errorf("could not find any stage provisioner plugin for type '%s'", name)
}
