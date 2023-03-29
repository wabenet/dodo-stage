package state

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/wabenet/dodo-core/pkg/config"
)

const (
	stateFilename = "state.json"
	permStateFile = 0o644
	permDirectory = 0o700
)

func Location(name string) string {
	return filepath.Join(config.GetAppDir(), "stages", name)
}

func Exist(name string) (bool, error) {
	if _, err := os.Stat(Location(name)); err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}

		return false, fmt.Errorf("could not check state file for %s: %w", name, err)
	}

	return true, nil
}

func Load[S json.Marshaler](name string) (*S, error) {
	filename := filepath.Join(Location(name), stateFilename)

	var state S

	if _, err := os.Stat(filename); err != nil {
		if os.IsNotExist(err) {
			return &state, nil
		}

		return nil, fmt.Errorf("could not check state file for %s: %w", name, err)
	}

	stateFile, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("could read check state file for %s: %w", name, err)
	}

	if err := json.Unmarshal(stateFile, &state); err != nil {
		return nil, fmt.Errorf("could not parse state file for %s: %w", name, err)
	}

	return &state, nil
}

func Save[S json.Marshaler](name string, state *S) error {
	if err := os.MkdirAll(Location(name), permDirectory); err != nil {
		return fmt.Errorf("could not create state directory for %s: %w", name, err)
	}

	filename := filepath.Join(Location(name), stateFilename)

	stateFile, err := json.Marshal(state)
	if err != nil {
		return fmt.Errorf("could not marshal json state for %s: %w", name, err)
	}

	if err := os.WriteFile(filename, stateFile, permStateFile); err != nil {
		return fmt.Errorf("could not write state file for %s: %w", name, err)
	}

	return nil
}

func Delete(name string) error {
	if err := os.RemoveAll(Location(name)); err != nil {
		return fmt.Errorf("could not delete state file for %s: %w", name, err)
	}

	return nil
}
