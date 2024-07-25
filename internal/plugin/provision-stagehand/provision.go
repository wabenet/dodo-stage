package provision

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"net"
	"net/url"
	"os"
	"path/filepath"
	"time"

	log "github.com/hashicorp/go-hclog"
	"github.com/oclaussen/go-gimme/ssl"
	core "github.com/wabenet/dodo-core/api/core/v1alpha5"
	coreconfig "github.com/wabenet/dodo-core/pkg/config"
	"github.com/wabenet/dodo-core/pkg/plugin"
	api "github.com/wabenet/dodo-stage/api/provision/v1alpha1"
	stage "github.com/wabenet/dodo-stage/api/stage/v1alpha3"
	"github.com/wabenet/dodo-stage/internal/config"
	"github.com/wabenet/dodo-stage/pkg/plugin/provision"
	"github.com/wabenet/dodo-stage/pkg/proxy"
	"github.com/wabenet/dodo-stage/pkg/stagehand/installer"
)

const (
	name        = "stagehand"
	DefaultPort = 20257
)

var _ provision.Provisioner = &Provisioner{}

type Provisioner struct {
	proxyClient *proxy.Client
}

func New() *Provisioner {
	return &Provisioner{}
}

func (*Provisioner) Type() plugin.Type {
	return provision.Type
}

func (*Provisioner) PluginInfo() *core.PluginInfo {
	return &core.PluginInfo{
		Name: &core.PluginName{Name: name, Type: provision.Type.String()},
	}
}

func (*Provisioner) Init() (plugin.Config, error) {
	return map[string]string{}, nil
}

func (*Provisioner) Cleanup() {}

func (p *Provisioner) ProvisionStage(info *stage.StageInfo, sshOpts *stage.SSHOptions) error {
	stages, err := config.GetAllStages(coreconfig.GetConfigFiles()...)
	if err != nil {
		return err
	}
	cfg := stages[info.Name]

	inst := installer.SSHInstaller{
		DownloadUrl: cfg.Provision.StagehandURL,
		SSHOptions:  sshOpts,
	}

	certs, _, err := ssl.GimmeCertificates(&ssl.Options{
		Org:   fmt.Sprintf("dodo.%s", info.Name),
		Hosts: []string{info.Hostname, "127.0.0.1", "localhost"},
	})
	if err != nil {
		return err
	}

	// TODO: this is ugly
	conf := map[string]interface{}{
		// TODO: at some point we lost the ability to replace authorized keys??
		//"add-ssh-key": map[string]interface{}{
		//	"default_user":    sshOpts.Username,
		//	"authorized_keys": []string{},
		//},
		"set-hostname": map[string]interface{}{
			"hostname": info.Name,
		},
		"run-script": map[string]interface{}{
			"script": cfg.Provision.Script,
		},
		"install-docker": map[string]interface{}{
			"ca":   string(certs.CA),
			"cert": string(certs.ServerCert),
			"key":  string(certs.ServerKey),
		},
		"install-proxy-service": map[string]interface{}{
			"address": "tcp://0.0.0.0:20257",
			"ca":      string(certs.CA),
			"cert":    string(certs.ServerCert),
			"key":     string(certs.ServerKey),
		},
	}

	provisionConfig, err := json.Marshal(&conf)

	if err := inst.Install(provisionConfig); err != nil {
		return err
	}

	if err := os.MkdirAll(storagePath(info.Name), 0700); err != nil {
		return err
	}

	if err := os.WriteFile(filepath.Join(storagePath(info.Name), "ca.pem"), certs.CA, 0600); err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(storagePath(info.Name), "client.pem"), certs.ClientCert, 0600); err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(storagePath(info.Name), "client-key.pem"), certs.ClientKey, 0600); err != nil {
		return err
	}

	pemData, _ := pem.Decode(certs.CA)
	caCert, err := x509.ParseCertificate(pemData.Bytes)
	if err != nil {
		return err
	}
	certPool := x509.NewCertPool()
	certPool.AddCert(caCert)

	keyPair, err := tls.X509KeyPair(certs.ClientCert, certs.ClientKey)
	if err != nil {
		return err
	}

	parsed, err := url.Parse(fmt.Sprintf("tcp://%s:%d", info.Hostname, DefaultPort))
	if err != nil {
		return fmt.Errorf("could not parse URL: %w", err)
	}

	if _, err = tls.DialWithDialer(
		&net.Dialer{Timeout: 20 * time.Second},
		"tcp",
		parsed.Host,
		&tls.Config{
			RootCAs:      certPool,
			ServerName:   parsed.Hostname(),
			Certificates: []tls.Certificate{keyPair},
		},
	); err != nil {
		return err
	}

	log.L().Info("stage is fully provisioned")

	return nil
}

func (p *Provisioner) GetClient(info *stage.StageInfo) (*proxy.Client, error) {
	if p.proxyClient != nil {
		return p.proxyClient, nil
	}

	pc, err := proxy.NewClient(&api.ProxyConfig{
		Url:      fmt.Sprintf("tcp://%s:%d", info.Hostname, DefaultPort),
		CaPath:   filepath.Join(storagePath(info.Name), "ca.pem"),
		CertPath: filepath.Join(storagePath(info.Name), "client.pem"),
		KeyPath:  filepath.Join(storagePath(info.Name), "client-key.pem"),
	})
	if err != nil {
		return nil, err
	}

	p.proxyClient = pc

	return p.proxyClient, nil

}

func (p *Provisioner) CleanStage(info *stage.StageInfo) error {
	if err := os.Remove(filepath.Join(storagePath(info.Name), "ca.pem")); err != nil {
		return err
	}

	if err := os.Remove(filepath.Join(storagePath(info.Name), "client.pem")); err != nil {
		return err
	}

	if err := os.Remove(filepath.Join(storagePath(info.Name), "client-key.pem")); err != nil {
		return err
	}

	return nil
}

func storagePath(name string) string {
	return filepath.Join(coreconfig.GetAppDir(), "stages", name)
}
