// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package webp

import (
	"image"
	"image/draw"
	"reflect"
)

type Image interface {
	// Get original type, such as *image.Gray, *image.RGBA, etc.
	BaseType() image.Image

	// Pix holds the image's pixels, as pixel values in machine order format. The pixel at
	// (x, y) starts at []Type(Pix[(y-Rect.Min.Y)*Stride:])[(x-Rect.Min.X)*Channels:].
	Pix() []byte
	// Stride is the Pix stride (in bytes) between vertically adjacent pixels.
	Stride() int
	// Rect is the image's bounds.
	Rect() image.Rectangle

	// 1:Gray, 2:GrayA, 3:RGB, 4:RGBA
	Channels() int
	// Uint8/Uint16/Int32/Int64/Float32/Float64
	Depth() reflect.Kind

	draw.Image
}
