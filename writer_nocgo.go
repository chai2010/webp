// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !cgo

package webp

import (
	"errors"
	"image"
	"io"
)

const DefaulQuality = 90

// Options are the encoding parameters.
type Options struct {
	Lossless bool
	Quality  float32 // 0 ~ 100
}

// Encode writes the image m to w in WEBP format.
func Encode(w io.Writer, m image.Image, opt *Options) (err error) {
	err = errors.New("webp.Encode: cgo is disabled!")
	return
}
