package simple_config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigGet_Exist(t *testing.T) {
	config := NewSimpleConfig("./config.test", "yml")
	value := config.Get("test-key")
	assert.Equal(t, value, "test value")
}

func TestConfigGet_NonExist(t *testing.T) {
	config := NewSimpleConfig("./config.test", "yml")
	value := config.Get("allauakbarrr")
	assert.Nil(t, value)
}

func TestConfigGet_String(t *testing.T) {
	config := NewSimpleConfig("./config.test", "yml")
	value := config.GetString("test-key")
	assert.Equal(t, value, "test value")
}
