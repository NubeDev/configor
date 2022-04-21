package configor

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type DummyWithPtr struct {
	IsExist *bool   `default:"true"`
	Name    *string `default:"RaiBnod"`
	Value   *int    `default:"22"`
}

type DummyWithoutPtr struct {
	IsExist bool   `default:"true"`
	Name    string `default:"RaiBnod"`
	Value   int    `default:"22"`
}

// Use of pointer:
// - when we need default value as `true`, but want to override sometimes from .yml config as `false`
// - when we need default value as `number != 0`, but want to override sometimes from .yml config as `0`
// - when we need default value as `string != ""`, but want to override sometimes from .yml config as `""`
func TestWithPointerConfig(t *testing.T) {
	adminConfig := DummyWithPtr{}
	New(&Config{}).Load(&adminConfig, "dummy_test.yml")
	assert.Equal(t, *adminConfig.IsExist, false)
	assert.Equal(t, *adminConfig.Name, "")
	assert.Equal(t, *adminConfig.Value, 0)
	fmt.Println("IsExist", *adminConfig.IsExist)
	fmt.Println("Name", *adminConfig.Name)
	fmt.Println("Value", *adminConfig.Value)
}

func TestWithoutPointerConfig(t *testing.T) {
	adminConfig := DummyWithoutPtr{}
	New(&Config{}).Load(&adminConfig, "dummy_test.yml")
	assert.Equal(t, adminConfig.IsExist, true)
	assert.Equal(t, adminConfig.Name, "RaiBnod")
	assert.Equal(t, adminConfig.Value, 22)
	fmt.Println("IsExist", adminConfig.IsExist)
	fmt.Println("Name", adminConfig.Name)
	fmt.Println("Value", adminConfig.Value)
}
