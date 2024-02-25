package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadConfig(t *testing.T) {
	config, err := ParseConfig("config.yaml")

	assert.NoError(t, err)

	assert.Equal(t, "Example Site", config.Title)
}
