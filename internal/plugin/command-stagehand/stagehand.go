package command

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
	"github.com/wabenet/dodo-core/pkg/plugin"
	"github.com/wabenet/dodo-stage/pkg/proxy"
	"github.com/wabenet/dodo-stage/pkg/stagehand"
)

func New(m plugin.Manager) *Command {
	cmd := &cobra.Command{
		Use:              "stagehand",
		Short:            "stage helper commands",
		TraverseChildren: true,
		SilenceUsage:     true,
	}

	cmd.AddCommand(NewProvisionCommand(m))
	cmd.AddCommand(NewProxyServerCommand(m))

	return &Command{cmd: cmd}
}

func NewProxyServerCommand(m plugin.Manager) *cobra.Command {
	var c proxy.Config
	cmd := &cobra.Command{
		Use:   "proxyserver",
		Short: "runs a grpc proxy server",
		RunE: func(_ *cobra.Command, _ []string) error {
			s, err := proxy.NewServer(m, &c)
			if err != nil {
				return err
			}

			return s.Listen()
		},
	}

	flags := cmd.Flags()
	flags.StringVarP(&c.Address, "address", "a", "tcp://0.0.0.0:20257", "listen address")
	flags.StringVar(&c.CAFile, "tls-ca-file", "", "ca file")
	flags.StringVar(&c.CertFile, "tls-cert-file", "", "certificate file")
	flags.StringVar(&c.KeyFile, "tls-key-file", "", "private key file")

	return cmd
}

type provisionConfig struct {
	path string
}

func NewProvisionCommand(m plugin.Manager) *cobra.Command {
	var c provisionConfig
	cmd := &cobra.Command{
		Use:   "provision",
		Short: "prepares the stage for usage",
		RunE: func(_ *cobra.Command, _ []string) error {
			configFile, err := ioutil.ReadFile(c.path)
			if err != nil {
				return err
			}

			config, err := stagehand.DecodeConfig(configFile)
			if err != nil {
				return err
			}

			result, err := stagehand.Provision(config)
			if err != nil {
				return err
			}

			output, err := stagehand.EncodeResult(result)
			if err != nil {
				return err
			}

			fmt.Fprintf(os.Stdout, string(output))

			return nil
		},
	}

	flags := cmd.Flags()
	flags.StringVar(&c.path, "config", "c", "config file")

	return cmd
}
