// Copyright 2016 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// See https://github.com/dvyukov/go-fuzz
package fuzz

import (
	"bytes"

	"github.com/glados28/webp"
)

func Fuzz(data []byte) int {
	cfg, err := webp.DecodeConfig(bytes.NewReader(data))
	if err != nil {
		return 0
	}
	if cfg.Width*cfg.Height > 1e6 {
		return 0
	}
	if _, err := webp.Decode(bytes.NewReader(data)); err != nil {
		return 0
	}
	return 1
}
