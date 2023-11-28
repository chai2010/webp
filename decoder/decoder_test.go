// The MIT License (MIT)
//
// Copyright (c) 2019 Amangeldy Kadyl
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
// the Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
// FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
// COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
// IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package decoder

import (
	"image"
	"os"
	"testing"

	"github.com/cqzcqq/webp/utils"
	"github.com/stretchr/testify/require"
)

func TestNewDecoder(t *testing.T) {
	t.Run("create success", func(t *testing.T) {
		file, err := os.Open("../test_data/images/m4_q75.webp")
		if err != nil {
			t.Fatal(err)
		}

		if d, err := NewDecoder(file, &Options{
			BypassFiltering:   true,
			NoFancyUpsampling: true,
			Scale: image.Rectangle{
				Max: image.Point{
					X: 300,
					Y: 300,
				},
			},
			UseThreads:             true,
			Flip:                   true,
			DitheringStrength:      0,
			AlphaDitheringStrength: 0,
		}); err != nil {
			t.Fatal(err)
		} else if img, err := d.Decode(); err != nil {
			t.Fatal(err)
		} else if img.Bounds().Max.X <= 0 || img.Bounds().Max.Y <= 0 {
			t.Fatal("invalid decoding result")
		}
	})
	t.Run("empty file", func(t *testing.T) {
		file, err := os.Open("../test_data/images/invalid.webp")
		if err != nil {
			t.Fatal(err)
		}

		if _, err := NewDecoder(file, &Options{
			Crop: image.Rectangle{
				Max: image.Point{
					X: 150,
					Y: 150,
				},
			},
		}); err == nil {
			t.Fatal(err)
		}
	})
}

func TestDecoder_GetFeatures(t *testing.T) {
	file, err := os.Open("../test_data/images/m4_q75.webp")
	if err != nil {
		t.Fatal(err)
	}

	dec, err := NewDecoder(file, &Options{})
	if err != nil {
		t.Fatal(err)
	}

	features := dec.GetFeatures()

	if features.Width != 675 || features.Height != 900 {
		t.Fatal("incorrect dimensions")
	}

	if features.Format != utils.FormatLossy {
		t.Fatal("file format is invalid")
	}

	if features.HasAlpha {
		t.Fatal("file has_alpha is invalid")
	}

	if features.HasAlpha {
		t.Fatal("file has_animation is invalid")
	}
}

func TestDecoder_NilOptions(t *testing.T) {
	file, err := os.Open("../test_data/images/m4_q75.webp")
	require.NoError(t, err)

	_, err = NewDecoder(file, nil)
	require.NoError(t, err)
}
