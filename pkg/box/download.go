package box

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/cavaliergopher/grab/v3"
	log "github.com/hashicorp/go-hclog"
	"github.com/mholt/archiver"
	"github.com/pkg/errors"
)

// TODO: automatically check for updates and download if necessary

func (box *Box) Download() error {
	if _, err := os.Stat(filepath.Join(box.Path(), "metadata.json")); err == nil {
		log.L().Info("box already exist")
		return nil
	}

	filename := fmt.Sprintf(
		"%s-%s-%s-%s.box",
		box.metadata.Username,
		box.metadata.Name,
		box.version.Version,
		box.provider.Name,
	)

	client := grab.NewClient()
	req, err := grab.NewRequest(filepath.Join(box.tmpPath, filename), box.provider.DownloadUrl)
	if err != nil {
		return err
	}

	log.L().Info("downloading box...")
	resp := client.Do(req)
	tick(resp)
	if err := resp.Err(); err != nil {
		return err
	}

	log.L().Info("extracting box...")
	if err := unarchive(filepath.Join(box.tmpPath, filename), box.Path()); err != nil {
		return err
	}

	if err := os.Remove(filepath.Join(box.tmpPath, filename)); err != nil {
		return err
	}

	return nil
}

func tick(resp *grab.Response) {
	t := time.NewTicker(500 * time.Millisecond)
	defer t.Stop()

	for {
		select {
		case <-t.C:
			fmt.Fprintf(os.Stderr, "%v / %v bytes (%.2f%%)\n", resp.BytesComplete(), resp.Size(), 100*resp.Progress())
		case <-resp.Done:
			return
		}
	}
}

func unarchive(source string, target string) error {
	unarchivers := []archiver.Unarchiver{
		&archiver.Tar{OverwriteExisting: true, MkdirAll: true, ImplicitTopLevelFolder: false},
		&archiver.TarGz{Tar: &archiver.Tar{OverwriteExisting: true, MkdirAll: true, ImplicitTopLevelFolder: false}},
		&archiver.Zip{OverwriteExisting: true, MkdirAll: true, ImplicitTopLevelFolder: false},
	}
	for _, unarchiver := range unarchivers {
		if err := unarchiver.Unarchive(source, target); err == nil {
			return nil
		}
	}
	return errors.New("unsupported archive format")
}
