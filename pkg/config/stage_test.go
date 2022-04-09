package config_test

import (
	"testing"

	"github.com/dodo-cli/dodo-stage/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestStage(t *testing.T) {
	cfg, err := config.GetAllStages("test/dodo.yaml")

	assert.Nil(t, err)

	stage, ok := cfg["test"]

	assert.True(t, ok)
	assert.Equal(t, "someplugin", stage.Type)
	assert.Equal(t, "debian", stage.Box.User)
	assert.Equal(t, "buster64", stage.Box.Name)
	assert.Equal(t, "10.0.0", stage.Box.Version)
	assert.Equal(t, int64(1), stage.Resources.Cpu)
	assert.Equal(t, int64(200*1000*1000), stage.Resources.Memory)
	assert.Equal(t, 1, len(stage.Resources.Volumes))
	assert.Equal(t, int64(100*1000*1000*1000), stage.Resources.Volumes[0].Size)
}
