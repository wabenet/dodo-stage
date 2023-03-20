package installer

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/cavaliergopher/grab/v3"
	log "github.com/hashicorp/go-hclog"
	"github.com/oclaussen/go-gimme/ssh"
	"github.com/pkg/errors"
	"github.com/wabenet/dodo-core/pkg/config"
	api "github.com/wabenet/dodo-stage/api/stage/v1alpha3"
	"github.com/wabenet/dodo-stage/pkg/stagehand"
)

const (
	targetPath = "/tmp/dodo/"
)

type SSHInstaller struct {
	DownloadUrl string
	SSHOptions  *api.SSHOptions
}

func (i *SSHInstaller) Install(cfg *stagehand.Config) (*stagehand.ProvisionResult, error) {
	localFile := filepath.Join(config.GetAppDir(), "tmp", "stagehand")

	if strings.HasPrefix(i.DownloadUrl, "file://") {
		localFile = i.DownloadUrl[7:]
	} else {
		_, err := grab.Get(localFile, i.DownloadUrl)
		if err != nil {
			return nil, fmt.Errorf("could not download stagehand from %s: %w", i.DownloadUrl, err)
		}
	}

	f, err := os.Open(localFile)
	if err != nil {
		return nil, fmt.Errorf("could not read file: %w", err)
	}

	stat, err := f.Stat()
	if err != nil {
		return nil, fmt.Errorf("could not read file: %w", err)
	}

	executor, err := ssh.GimmeExecutor(&ssh.Options{
		Host:              i.SSHOptions.Hostname,
		Port:              int(i.SSHOptions.Port),
		User:              i.SSHOptions.Username,
		IdentityFileGlobs: []string{i.SSHOptions.PrivateKeyFile},
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

	if err := executor.WriteFile(&ssh.FileOptions{
		Path:   path.Join(targetPath, "stagehand"),
		Reader: bufio.NewReader(f),
		Size:   stat.Size(),
		Mode:   0755,
	}); err != nil {
		return nil, fmt.Errorf("could not write stagehand binary: %w", err)
	}

	encoded, err := stagehand.EncodeConfig(cfg)
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
