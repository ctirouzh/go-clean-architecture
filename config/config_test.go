package config

import (
	"errors"
	"io/fs"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestConfig_Parse(t *testing.T) {
	// Prepare for test
	emptyConfigFile, err := os.Create("./../tmp/empty_config.json")
	if err != nil {
		t.Fatal(err)
	}
	emptyConfigFile.Close()
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
			name:         "unmarshal example.config.json file",
			relativePath: "./",
			fileName:     "example.config.json",
			want: Want{
				err: nil,
				config: &Config{
					Server: Server{
						Host: "127.0.0.1",
						Port: 3000,
					},
					JWT: JWT{
						SecretKey: "your_secret_key",
						TTL:       20 * time.Minute,
					},
					Postgres: Postgres{
						Host: "localhost",
						Port: 5432,
						User: "postgres",
						Pass: "your_db_password",
						Name: "your_db_name",
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
			name:         "empty config file",
			relativePath: "./../tmp/",
			fileName:     emptyConfigFile.Name(),
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

func TestConfig_ConfigFileExists(t *testing.T) {
	_, err := os.Stat("config.json")
	assert.Falsef(t, errors.Is(err, fs.ErrNotExist), "config.json file not exists")
}
