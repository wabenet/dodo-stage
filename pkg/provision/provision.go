//go:generate env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o assets/stagehand github.com/wabenet/dodo-stage/pkg/provision/stagehand

package provision

import (
	"bytes"
	_ "embed"
	"fmt"
	"path"
	"strings"

	log "github.com/hashicorp/go-hclog"
	"github.com/oclaussen/go-gimme/ssh"
	"github.com/pkg/errors"
	api "github.com/wabenet/dodo-stage/api/v1alpha2"
	"github.com/wabenet/dodo-stage/pkg/stagehand"
)

//go:embed assets/stagehand
var StagehandBinary string

const targetPath = "/tmp/dodo/"

func Provision(sshOpts *api.SSHOptions, config *stagehand.Config) (*stagehand.ProvisionResult, error) {
	executor, err := ssh.GimmeExecutor(&ssh.Options{
		Host:              sshOpts.Hostname,
		Port:              int(sshOpts.Port),
		User:              sshOpts.Username,
		IdentityFileGlobs: []string{sshOpts.PrivateKeyFile},
		NonInteractive:    true,
	})
	if err != nil {
		return nil, fmt.Errorf("could not get SSH connection to stage: %w", err)
	}
	defer executor.Close()

	log.L().Debug("write stagehand to stage")
	if out, err := executor.Execute(fmt.Sprintf("mkdir -p %s", targetPath)); err != nil {
		return nil, errors.Wrap(err, out)
	}

	// TODO: needs root
	if err := executor.WriteFile(&ssh.FileOptions{
		Path:   path.Join(targetPath, "stagehand"),
		Reader: strings.NewReader(StagehandBinary),
		Size:   int64(len(StagehandBinary)),
		Mode:   0755,
	}); err != nil {
		return nil, fmt.Errorf("could not write stagehand binary: %w", err)
	}

	encoded, err := stagehand.EncodeConfig(config)
	if err != nil {
		return nil, fmt.Errorf("could not generate stagehand config: %w", err)
	}
	if err := executor.WriteFile(&ssh.FileOptions{
		Path: path.Join(targetPath, "config.json"),
		// TODO: this is a bug in gimme. Fix it!
		//Content: encoded,
		Reader: bytes.NewReader(encoded),
		Size:   int64(len(encoded)),
		Mode:   0644,
	}); err != nil {
		return nil, fmt.Errorf("could not write stagehand config: %w", err)
	}

	// TODO: figure out whether we have/need sudo
	log.L().Debug("executing stagehand")
	out, err := executor.Execute(fmt.Sprintf(
		"sudo %s provision --config %s",
		path.Join(targetPath, "stagehand"),
		path.Join(targetPath, "config.json"),
	))
	if err != nil {
		return nil, errors.Wrap(err, out)
	}

	result, err := stagehand.DecodeResult([]byte(out))
	if err != nil {
		return nil, fmt.Errorf("could not read stagehand output: %w", err)
	}

	return result, nil
}
