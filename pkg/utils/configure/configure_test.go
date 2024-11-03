package configure_test

import (
	"testing"
	"time"

	"sfw/pkg/utils/configure"
)

func TestNewLoader(t *testing.T) {
	option := &configure.ConfigureOption{
		ConfigName:    `test`,
		ConfigType:    `yaml`,
		ConfigPath:    `.`,
		RegisterParam: []interface{}{},
		Register:      configure.Register,
	}
	err := configure.NewConfLoader(option).Run()
	time.Sleep(2 * time.Second)
	if err != nil {
		t.Error(err)
	}
	t.Log("TestNewLoader passed")
}
