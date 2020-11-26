package types

import (
	"testing"

	"github.com/dodo-cli/dodo-core/pkg/decoder"
	api "github.com/dodo-cli/dodo-stage/api/v1alpha1"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

const fullExample = `
type: generic
box:
  user: debian
  name: buster64
  version: 10.0.0
resources:
  cpu: 4
  memory: 8192
  volumes:
    - size: 1GB
  usb:
    - name: 'Test'
      vendorid: 'foo'
      productid: 'bar'
`

func TestFullExample(t *testing.T) {
	config := getExampleConfig(t, fullExample)
	assert.Equal(t, "generic", config.Type)
	assert.NotNil(t, config.Box)
	assert.Equal(t, "debian", config.Box.User)
	assert.Equal(t, "buster64", config.Box.Name)
	assert.Equal(t, "10.0.0", config.Box.Version)
	assert.NotNil(t, config.Resources)
	assert.Equal(t, int64(4), config.Resources.Cpu)
	assert.Equal(t, int64(8192), config.Resources.Memory)
	assert.Equal(t, 1, len(config.Resources.Volumes))
	assert.Equal(t, int64(1000000000), config.Resources.Volumes[0].Size)
	assert.Equal(t, 1, len(config.Resources.UsbFilters))
	assert.Equal(t, "Test", config.Resources.UsbFilters[0].Name)
	assert.Equal(t, "foo", config.Resources.UsbFilters[0].VendorId)
	assert.Equal(t, "bar", config.Resources.UsbFilters[0].ProductId)
}

func getExampleConfig(t *testing.T, yamlConfig string) *api.Stage {
	// TODO: clean up this part
	var mapType map[interface{}]interface{}
	err := yaml.Unmarshal([]byte(yamlConfig), &mapType)
	assert.Nil(t, err)

	produce := NewStage()
	ptr, decode := produce()
	config := *(ptr.(**api.Stage))
	d := decoder.New("test")
	decode(d, mapType)

	return config
}
