package config_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wabenet/dodo-stage/internal/config"
)

func TestStage(t *testing.T) {
	cfg, err := config.GetAllStages("test/dodo.yaml")

	assert.Nil(t, err)

	stage, ok := cfg["test"]

	assert.True(t, ok)
	assert.Equal(t, "teststage", stage.Name)
	assert.Equal(t, "someplugin", stage.Type)
}
