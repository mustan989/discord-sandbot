package config_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"pgregory.net/rapid"

	"github.com/mustan989/discord-sandbot/internal/config"
)

const configTemplate = `
http:
  port: %d
`

func TestRead(t *testing.T) {
	t.Run("Default", rapid.MakeCheck(func(t *rapid.T) {
		port := rapid.Int().Draw(t, "http_port")

		content := fmt.Appendf(nil, configTemplate, port)
		want := &config.Config{
			HTTP: config.HTTP{
				Port: port,
			},
		}

		config, err := config.Read(bytes.NewBuffer(content))
		require.NoError(t, err)
		require.Equal(t, want, config)
	}))

	t.Run("Error", rapid.MakeCheck(func(t *rapid.T) {
		http := rapid.Int().Draw(t, "http")

		content := fmt.Appendf(nil, `http: %d`, http)

		config, err := config.Read(bytes.NewBuffer(content))
		require.Error(t, err)
		require.Nil(t, config)
	}))
}
