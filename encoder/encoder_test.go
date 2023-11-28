package encoder

import (
	"bytes"
	"image"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewEncoder(t *testing.T) {
	t.Run("encode nrgba image with lossy preset", func(t *testing.T) {
		expected := &bytes.Buffer{}
		img := image.NewNRGBA(image.Rectangle{
			Max: image.Point{
				X: 100,
				Y: 150,
			},
		})

		options, err := NewLossyEncoderOptions(PresetDefault, 0.75)
		require.NoError(t, err)

		e, err := NewEncoder(img, options)
		require.NoError(t, err)

		err = e.Encode(expected)
		require.NoError(t, err)

		actual, err := os.ReadFile("../test_data/images/100x150_lossy.webp")
		require.NoError(t, err)

		assert.Equal(t, actual, expected.Bytes())
	})
	t.Run("encode nrgba image with lossless preset", func(t *testing.T) {
		actuall := &bytes.Buffer{}
		img := image.NewNRGBA(image.Rectangle{
			Max: image.Point{
				X: 100,
				Y: 150,
			},
		})

		options, err := NewLosslessEncoderOptions(PresetDefault, 1)
		require.NoError(t, err)

		e, err := NewEncoder(img, options)
		require.NoError(t, err)

		err = e.Encode(actuall)
		require.NoError(t, err)

		expected, err := os.ReadFile("../test_data/images/100x150_lossless.webp")
		require.NoError(t, err)

		assert.Equal(t, expected, actuall.Bytes())
	})
}

func TestNewEncoder_NilOptions(t *testing.T) {
	t.Run("encode image without passing options should not panic", func(t *testing.T) {
		expected := &bytes.Buffer{}
		img := image.NewNRGBA(image.Rectangle{
			Max: image.Point{
				X: 100,
				Y: 150,
			},
		})

		e, err := NewEncoder(img, nil)
		require.NoError(t, err)

		err = e.Encode(expected)
		require.NoError(t, err)
	})
}
