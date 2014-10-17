// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package webp

import (
	"image"
	"image/color"
	"reflect"
)

// _RGB is an in-memory image whose At method returns color.RGBA values.
//
// Notes: _RGB use the same struct with image.RGBA!!!
type _RGB struct {
	_Pix    []uint8
	_Stride int
	_Rect   image.Rectangle
}

func (p *_RGB) BaseType() image.Image { return p }
func (p *_RGB) Pix() []byte           { return p._Pix }
func (p *_RGB) Stride() int           { return p._Stride }
func (p *_RGB) Rect() image.Rectangle { return p._Rect }
func (p *_RGB) Channels() int         { return 3 }
func (p *_RGB) Depth() reflect.Kind   { return reflect.Uint8 }

func (p *_RGB) ColorModel() color.Model { return color.RGBAModel }

func (p *_RGB) Bounds() image.Rectangle { return p._Rect }

func (p *_RGB) At(x, y int) color.Color {
	return p.RGBAAt(x, y)
}

func (p *_RGB) RGBAAt(x, y int) color.RGBA {
	if !(image.Point{x, y}.In(p._Rect)) {
		return color.RGBA{}
	}
	i := p.PixOffset(x, y)
	return color.RGBA{p._Pix[i+0], p._Pix[i+1], p._Pix[i+2], 0xFF}
}

// PixOffset returns the index of the first element of _Pix that corresponds to
// the pixel at (x, y).
func (p *_RGB) PixOffset(x, y int) int {
	return (y-p._Rect.Min.Y)*p._Stride + (x-p._Rect.Min.X)*3
}

func (p *_RGB) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(p._Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	c1 := color.RGBAModel.Convert(c).(color.RGBA)
	p._Pix[i+0] = c1.R
	p._Pix[i+1] = c1.G
	p._Pix[i+2] = c1.B
}

func (p *_RGB) SetRGBA(x, y int, c color.RGBA) {
	if !(image.Point{x, y}.In(p._Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	p._Pix[i+0] = c.R
	p._Pix[i+1] = c.G
	p._Pix[i+2] = c.B
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *_RGB) SubImage(r image.Rectangle) image.Image {
	r = r.Intersect(p._Rect)
	// If r1 and r2 are Rectangles, r1.Intersect(r2) is not guaranteed to be inside
	// either r1 or r2 if the intersection is empty. Without explicitly checking for
	// this, the _Pix[i:] expression below can panic.
	if r.Empty() {
		return &_RGB{}
	}
	i := p.PixOffset(r.Min.X, r.Min.Y)
	return &_RGB{
		_Pix:    p._Pix[i:],
		_Stride: p._Stride,
		_Rect:   r,
	}
}

// Opaque scans the entire image and reports whether it is fully opaque.
func (p *_RGB) Opaque() bool {
	if p._Rect.Empty() {
		return true
	}
	i0, i1 := 3, p._Rect.Dx()*3
	for y := p._Rect.Min.Y; y < p._Rect.Max.Y; y++ {
		for i := i0; i < i1; i += 3 {
			if p._Pix[i] != 0xff {
				return false
			}
		}
		i0 += p._Stride
		i1 += p._Stride
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
