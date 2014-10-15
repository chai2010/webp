// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package webp

import (
	"image"
	"image/color"
)

// _RGB is an in-memory image whose At method returns color.RGBA values.
//
// Notes: _RGB use the same struct with image.RGBA!!!
type _RGB struct {
	// Pix holds the image's pixels, in R, G, B order. The pixel at
	// (x, y) starts at Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*3].
	Pix []uint8
	// Stride is the Pix stride (in bytes) between vertically adjacent pixels.
	Stride int
	// Rect is the image's bounds.
	Rect image.Rectangle
}

func (p *_RGB) ColorModel() color.Model { return color.RGBAModel }

func (p *_RGB) Bounds() image.Rectangle { return p.Rect }

func (p *_RGB) At(x, y int) color.Color {
	return p.RGBAAt(x, y)
}

func (p *_RGB) RGBAAt(x, y int) color.RGBA {
	if !(image.Point{x, y}.In(p.Rect)) {
		return color.RGBA{}
	}
	i := p.PixOffset(x, y)
	return color.RGBA{p.Pix[i+0], p.Pix[i+1], p.Pix[i+2], 0xFF}
}

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (p *_RGB) PixOffset(x, y int) int {
	return (y-p.Rect.Min.Y)*p.Stride + (x-p.Rect.Min.X)*3
}

func (p *_RGB) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(p.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	c1 := color.RGBAModel.Convert(c).(color.RGBA)
	p.Pix[i+0] = c1.R
	p.Pix[i+1] = c1.G
	p.Pix[i+2] = c1.B
}

func (p *_RGB) SetRGBA(x, y int, c color.RGBA) {
	if !(image.Point{x, y}.In(p.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	p.Pix[i+0] = c.R
	p.Pix[i+1] = c.G
	p.Pix[i+2] = c.B
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *_RGB) SubImage(r image.Rectangle) image.Image {
	r = r.Intersect(p.Rect)
	// If r1 and r2 are Rectangles, r1.Intersect(r2) is not guaranteed to be inside
	// either r1 or r2 if the intersection is empty. Without explicitly checking for
	// this, the Pix[i:] expression below can panic.
	if r.Empty() {
		return &_RGB{}
	}
	i := p.PixOffset(r.Min.X, r.Min.Y)
	return &_RGB{
		Pix:    p.Pix[i:],
		Stride: p.Stride,
		Rect:   r,
	}
}

// Opaque scans the entire image and reports whether it is fully opaque.
func (p *_RGB) Opaque() bool {
	if p.Rect.Empty() {
		return true
	}
	i0, i1 := 3, p.Rect.Dx()*3
	for y := p.Rect.Min.Y; y < p.Rect.Max.Y; y++ {
		for i := i0; i < i1; i += 3 {
			if p.Pix[i] != 0xff {
				return false
			}
		}
		i0 += p.Stride
		i1 += p.Stride
	}
	return true
}

// newRGB returns a new _RGB with the given bounds.
func newRGB(r image.Rectangle) *_RGB {
	w, h := r.Dx(), r.Dy()
	buf := make([]uint8, 3*w*h)
	return &_RGB{buf, 3 * w, r}
}

func newRGBFromImage(m image.Image) *_RGB {
	switch m := m.(type) {
	case *_RGB:
		return m
	default:
		b := m.Bounds()
		rgb := newRGB(b)
		dstColorRGBA64 := &color.RGBA64{}
		dstColor := color.Color(dstColorRGBA64)
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				pr, pg, pb, _ := m.At(x, y).RGBA()
				dstColorRGBA64.R = uint16(pr)
				dstColorRGBA64.G = uint16(pg)
				dstColorRGBA64.B = uint16(pb)
				rgb.Set(x, y, dstColor)
			}
		}
		return rgb
	}
}
