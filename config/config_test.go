package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig_Parse(t *testing.T) {
	type Want struct {
		err    error
		config *Config
	}
	testCases := []struct {
		name         string
		relativePath string
		fileName     string
		want         Want
	}{
		{
			name:         "unmarshal config.json file",
			relativePath: "./",
			fileName:     "config.json",
			want: Want{
				err: nil,
				config: &Config{
					Server: Server{
						Host: "127.0.0.1",
						Port: 3000,
					},
				},
			},
		},
		{
			name:         "wrong directory",
			relativePath: "./../",
			fileName:     "config.json",
			want: Want{
				err:    ErrFailedToReadFile,
				config: nil,
			},
		},
		{
			name:         "invalid file format",
			relativePath: "./../tmp/",
			fileName:     "example.txt",
			want: Want{
				err:    ErrFailedToUnmarshal,
				config: nil,
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gotConfig, gotErr := Parse(tc.relativePath + tc.fileName)
			assert.Equal(t, tc.want.err, gotErr)
			assert.Equal(t, tc.want.config, gotConfig)
		})
	}

}
