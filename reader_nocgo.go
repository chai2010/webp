// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !cgo

package webp // import "github.com/chai2010/webp"

import (
	"image"
	"io"

	"github.com/chai2010/webp/internal/webp"
)

// DecodeConfig returns the color model and dimensions of a WEBP image without
// decoding the entire image.
func DecodeConfig(r io.Reader) (config image.Config, err error) {
	return webp.DecodeConfig(r)
}

// Decode reads a WEBP image from r and returns it as an image.Image.
func Decode(r io.Reader) (m image.Image, err error) {
	return webp.Decode(r)
}

func init() {
	image.RegisterFormat("webp", "RIFF????WEBPVP8", Decode, DecodeConfig)
}
