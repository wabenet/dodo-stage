//go:generate env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o assets/stage-designer github.com/wabenet/dodo-stage/cmd/stage-designer

package stage

import (
	"bytes"
	_ "embed"
	"fmt"
	"strings"

	log "github.com/hashicorp/go-hclog"
	"github.com/oclaussen/go-gimme/ssh"
	"github.com/pkg/errors"
	api "github.com/wabenet/dodo-stage/api/v1alpha2"
	"github.com/wabenet/dodo-stage/pkg/stagedesigner"
)

//go:embed assets/stage-designer
var StageDesignerBinary string

func Provision(sshOpts *api.SSHOptions, config *stagedesigner.Config) (*stagedesigner.ProvisionResult, error) {
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

	log.L().Debug("write stage designer to stage")
	if err := executor.WriteFile(&ssh.FileOptions{
		Path:   "/tmp/stage-designer",
		Reader: strings.NewReader(StageDesignerBinary),
		Size:   int64(len(StageDesignerBinary)),
		Mode:   0755,
	}); err != nil {
		return nil, fmt.Errorf("could not write stage designer binary: %w", err)
	}

	encoded, err := stagedesigner.EncodeConfig(config)
	if err != nil {
		return nil, fmt.Errorf("could not generate stage designer config: %w", err)
	}
	if err := executor.WriteFile(&ssh.FileOptions{
		Path: "/tmp/stage-designer-config",
		// TODO: this is a bug in gimme. Fix it!
		//Content: encoded,
		Reader: bytes.NewReader(encoded),
		Size:   int64(len(encoded)),
		Mode:   0644,
	}); err != nil {
		return nil, fmt.Errorf("could not write stage designer config: %w", err)
	}

	// TODO: figure out whether we have/need sudo
	log.L().Debug("executing stage designer")
	out, err := executor.Execute("sudo /tmp/stage-designer /tmp/stage-designer-config")
	if err != nil {
		return nil, errors.Wrap(err, out)
	}

	result, err := stagedesigner.DecodeResult([]byte(out))
	if err != nil {
		return nil, fmt.Errorf("could not read stage designer output: %w", err)
	}

	return result, nil
}
