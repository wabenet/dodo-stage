package command

import (
	"fmt"

	"github.com/dodo-cli/dodo-core/pkg/decoder"
	"github.com/dodo-cli/dodo-core/pkg/plugin"
	api "github.com/dodo-cli/dodo-stage/api/v1alpha1"
	"github.com/dodo-cli/dodo-stage/pkg/stage"
	"github.com/dodo-cli/dodo-stage/pkg/types"
	"github.com/oclaussen/go-gimme/configfiles"
	"github.com/oclaussen/go-gimme/ssh"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func NewStageCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:              "stage",
		Short:            "Manage stages",
		TraverseChildren: true,
		SilenceUsage:     true,
	}

	cmd.AddCommand(NewListCommand())
	cmd.AddCommand(NewUpCommand())
	cmd.AddCommand(NewDownCommand())
	cmd.AddCommand(NewSSHCommand())
	return cmd
}

func NewListCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List stages",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			stages := map[string]*api.Stage{}
			configfiles.GimmeConfigFiles(&configfiles.Options{
				Name:                      "dodo",
				Extensions:                []string{"yaml", "yml", "json"},
				IncludeWorkingDirectories: true,
				Filter: func(configFile *configfiles.ConfigFile) bool {
					d := decoder.New(configFile.Path)
					d.DecodeYaml(configFile.Content, &stages, map[string]decoder.Decoding{
						"stages": decoder.Map(types.NewStage(), &stages),
					})
					return false
				},
			})

			for name, conf := range stages {
				s, err := loadPlugin(conf.Type)
				if err != nil {
					return err
				}

				current, err := s.GetStage(name)
				if err != nil {
					return err
				}

                                fmt.Printf("%s (%v)", name, current.Available)
			}

                        return nil
		},
	}
}

func NewUpCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "up",
		Short: "Create or start a stage",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			conf, err := loadStageConfig(args[0])
			if err != nil {
				return err
			}

			s, err := loadPlugin(conf.Type)
			if err != nil {
				return err
			}

			current, err := s.GetStage(args[0])
			if err != nil {
				return err
			}

			if !current.Exist {
				return s.CreateStage(conf)
			}
			return s.StartStage(args[0])
		},
	}
}

type downOptions struct {
	remove  bool
	volumes bool
	force   bool
}

func NewDownCommand() *cobra.Command {
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

			s, err := loadPlugin(conf.Type)
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

func NewSSHCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "ssh",
		Short: "login to the stage",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			conf, err := loadStageConfig(args[0])
			if err != nil {
				return err
			}

			s, err := loadPlugin(conf.Type)
			if err != nil {
				return err
			}

			current, err := s.GetStage(args[0])
			if err != nil {
				return err
			}

			if !current.Available {
				return errors.New("stage is not up")
			}

			return ssh.GimmeShell(&ssh.Options{
				Host:              current.SshOptions.Hostname,
				Port:              int(current.SshOptions.Port),
				User:              current.SshOptions.Username,
				IdentityFileGlobs: []string{current.SshOptions.PrivateKeyFile},
				NonInteractive:    true,
			})
		},
	}
}

func loadStageConfig(name string) (*api.Stage, error) {
	stages := map[string]*api.Stage{}
	configfiles.GimmeConfigFiles(&configfiles.Options{
		Name:                      "dodo",
		Extensions:                []string{"yaml", "yml", "json"},
		IncludeWorkingDirectories: true,
		Filter: func(configFile *configfiles.ConfigFile) bool {
			d := decoder.New(configFile.Path)
			d.DecodeYaml(configFile.Content, &stages, map[string]decoder.Decoding{
				"stages": decoder.Map(types.NewStage(), &stages),
			})
			return false
		},
	})

	if conf, ok := stages[name]; ok {
		conf.Name = name // TODO: figure out where to set defaults like this
		return conf, nil
	}

	return nil, fmt.Errorf("could not find any configuration for stage '%s'", name)
}

func loadPlugin(name string) (stage.Stage, error) {
	for _, p := range plugin.GetPlugins(stage.Type.String()) {
		s := p.(stage.Stage)
		if info, err := s.PluginInfo(); err == nil && info.Name == name {
			return s, nil
		}
	}

	return nil, fmt.Errorf("could not find any stage plugin for type '%s'", name)
}
