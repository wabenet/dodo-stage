package ssh

import (
	"fmt"
	"os"
	"syscall"

	log "github.com/hashicorp/go-hclog"
	"github.com/wabenet/dodo-core/pkg/ui"
	stage "github.com/wabenet/dodo-stage/api/stage/v1alpha4"
	"golang.org/x/crypto/ssh"
)

func Shell(opts *stage.SSHOptions) error {
	client, err := getClient(opts)
	if err != nil {
		return err
	}

	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		return fmt.Errorf("could not start ssh session: %w", err)
	}

	defer session.Close()

	t := ui.NewTerminal()
	t = t.OnSignal(func(s os.Signal, t *ui.Terminal) {
		if s == syscall.SIGWINCH {
			if err := session.WindowChange(int(t.Height), int(t.Width)); err != nil {
				log.L().Warn("could not resize terminal", "error", err)
			}
		}
	})

	if err = t.RunInRaw(func(t *ui.Terminal) error {
		session.Stdin = t.Stdin
		session.Stdout = t.Stdout
		session.Stderr = t.Stderr

		if err := session.RequestPty("xterm", int(t.Height), int(t.Width), ssh.TerminalModes{}); err != nil {
			return fmt.Errorf("could not allocate remote terminal: %w", err)
		}

		if err := session.Shell(); err != nil {
			return fmt.Errorf("could not ssh shell: %w", err)
		}

		if err := session.Wait(); err != nil {
			return fmt.Errorf("error during ssh session: %w", err)
		}

		return nil
	}); err != nil {
		return fmt.Errorf("error during terminal session: %w", err)
	}

	return nil
}
