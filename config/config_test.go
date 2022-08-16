package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig_Parse(t *testing.T) {
	want := &Config{
		Server: Server{
			Host: "127.0.0.1",
			Port: 3000,
		},
	}
	got, err := Parse("config.json")
	assert.Equal(t, nil, err)
	assert.NotEmpty(t, got)
	assert.Equal(t, want, got)
}
