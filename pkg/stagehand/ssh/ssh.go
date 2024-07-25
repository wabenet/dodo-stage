package ssh

import (
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	Type = "add-ssh-key"
)

type Action struct {
	DefaultUser    string   `mapstructure:"default_user"`
	AuthorizedKeys []string `mapstructure:"authorized_keys"`
}

func (a *Action) Type() string {
	return Type
}

func (a *Action) Execute() error {
	log.Printf("replace SSH key...")

	u, err := user.Lookup(a.DefaultUser)
	if err != nil {
		return err
	}
	uid, err := strconv.Atoi(u.Uid)
	if err != nil {
		return err
	}
	gid, err := strconv.Atoi(u.Gid)
	if err != nil {
		return err
	}

	file := filepath.Join(u.HomeDir, ".ssh", "authorized_keys")
	content := strings.Join(a.AuthorizedKeys, "\n")
	if err := ioutil.WriteFile(file, []byte(content), 0600); err != nil {
		return err
	}
	if err := os.Chown(file, uid, gid); err != nil {
		return err
	}
	return nil
}
