package encoder

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNilOptions(t *testing.T) {
	t.Run("config from nil Options succeeds", func(t *testing.T) {
		var options *Options = nil
		cfg, err := options.GetConfig()
		require.NoError(t, err)
		require.NotNil(t, cfg)
	})
}

func TestNewLossyEncoderOptions(t *testing.T) {
	t.Run("invalid quality", func(t *testing.T) {
		options, err := NewLossyEncoderOptions(PresetDefault, -1)
		require.Error(t, err)
		assert.Nil(t, options)
	})
	t.Run("create lossy encoder options is success", func(t *testing.T) {
		options, err := NewLossyEncoderOptions(PresetDefault, 75)
		require.Nil(t, err)
		require.NotNil(t, options)

		assert.False(t, options.Lossless)
		assert.Equal(t, float32(75), options.Quality)

		cfg, err := options.GetConfig()
		require.NoError(t, err)
		assert.NotNil(t, cfg)
	})
}

func TestNewLosslessEncoderOptions(t *testing.T) {
	t.Run("invalid level of lossless encoder", func(t *testing.T) {
		options, err := NewLosslessEncoderOptions(PresetDefault, -1)
		require.Error(t, err)
		assert.Nil(t, options)
	})
	t.Run("create lossless encoder options is success", func(t *testing.T) {
		options, err := NewLosslessEncoderOptions(PresetDefault, 1)
		require.Nil(t, err)
		require.NotNil(t, options)

		assert.True(t, options.Lossless)
		assert.Equal(t, 1, options.Method)
		assert.Equal(t, float32(20), options.Quality)

		cfg, err := options.GetConfig()
		require.NoError(t, err)
		assert.NotNil(t, cfg)
	})
}
