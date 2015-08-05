// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package webp

import (
	"image"
	"image/color"
	"reflect"
)

var (
	_ image.Image = (*RGB48Image)(nil)
	_ MemP        = (*RGB48Image)(nil)
)

type RGB48Image struct {
	XPix    []uint8 // XPix use Native Endian (same as MemP) !!!
	XStride int
	XRect   image.Rectangle
}

func (p *RGB48Image) MemPMagic() string {
	return MemPMagic
}

func (p *RGB48Image) Bounds() image.Rectangle {
	return p.XRect
}

func (p *RGB48Image) Channels() int {
	return 3
}

func (p *RGB48Image) DataType() reflect.Kind {
	return reflect.Uint16
}

func (p *RGB48Image) Pix() []byte {
	return p.XPix
}

func (p *RGB48Image) Stride() int {
	return p.XStride
}

func (p *RGB48Image) ColorModel() color.Model { return color.RGBA64Model }

func (p *RGB48Image) At(x, y int) color.Color {
	if !(image.Point{x, y}.In(p.XRect)) {
		return color.RGBA64{}
	}
	i := p.PixOffset(x, y)
	if isLittleEndian {
		return color.RGBA64{
			R: uint16(p.XPix[i+1])<<8 | uint16(p.XPix[i+0]),
			G: uint16(p.XPix[i+3])<<8 | uint16(p.XPix[i+2]),
			B: uint16(p.XPix[i+4])<<8 | uint16(p.XPix[i+4]),
			A: 0xffff,
		}
	} else {
		return color.RGBA64{
			R: uint16(p.XPix[i+0])<<8 | uint16(p.XPix[i+1]),
			G: uint16(p.XPix[i+2])<<8 | uint16(p.XPix[i+3]),
			B: uint16(p.XPix[i+4])<<8 | uint16(p.XPix[i+5]),
			A: 0xffff,
		}
	}
}

func (p *RGB48Image) RGB48At(x, y int) [3]uint16 {
	if !(image.Point{x, y}.In(p.XRect)) {
		return [3]uint16{}
	}
	i := p.PixOffset(x, y)
	if isLittleEndian {
		return [3]uint16{
			uint16(p.XPix[i+1])<<8 | uint16(p.XPix[i+0]),
			uint16(p.XPix[i+3])<<8 | uint16(p.XPix[i+2]),
			uint16(p.XPix[i+5])<<8 | uint16(p.XPix[i+4]),
		}
	} else {
		return [3]uint16{
			uint16(p.XPix[i+0])<<8 | uint16(p.XPix[i+1]),
			uint16(p.XPix[i+2])<<8 | uint16(p.XPix[i+3]),
			uint16(p.XPix[i+4])<<8 | uint16(p.XPix[i+5]),
		}
	}
}

// PixOffset returns the index of the first element of XPix that corresponds to
// the pixel at (x, y).
func (p *RGB48Image) PixOffset(x, y int) int {
	return (y-p.XRect.Min.Y)*p.XStride + (x-p.XRect.Min.X)*3
}

func (p *RGB48Image) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(p.XRect)) {
		return
	}
	i := p.PixOffset(x, y)
	c1 := color.RGBA64Model.Convert(c).(color.RGBA64)
	if isLittleEndian {
		p.XPix[i+1] = uint8(c1.R >> 8)
		p.XPix[i+0] = uint8(c1.R)
		p.XPix[i+3] = uint8(c1.G >> 8)
		p.XPix[i+2] = uint8(c1.G)
		p.XPix[i+5] = uint8(c1.B >> 8)
		p.XPix[i+4] = uint8(c1.B)
	} else {
		p.XPix[i+0] = uint8(c1.R >> 8)
		p.XPix[i+1] = uint8(c1.R)
		p.XPix[i+2] = uint8(c1.G >> 8)
		p.XPix[i+3] = uint8(c1.G)
		p.XPix[i+4] = uint8(c1.B >> 8)
		p.XPix[i+5] = uint8(c1.B)
	}
	return
}

func (p *RGB48Image) SetRGB48(x, y int, c [3]uint16) {
	if !(image.Point{x, y}.In(p.XRect)) {
		return
	}
	i := p.PixOffset(x, y)
	if isLittleEndian {
		p.XPix[i+1] = uint8(c[0] >> 8)
		p.XPix[i+0] = uint8(c[0])
		p.XPix[i+3] = uint8(c[1] >> 8)
		p.XPix[i+2] = uint8(c[1])
		p.XPix[i+5] = uint8(c[2] >> 8)
		p.XPix[i+4] = uint8(c[2])
	} else {
		p.XPix[i+0] = uint8(c[0] >> 8)
		p.XPix[i+1] = uint8(c[0])
		p.XPix[i+2] = uint8(c[1] >> 8)
		p.XPix[i+3] = uint8(c[1])
		p.XPix[i+4] = uint8(c[2] >> 8)
		p.XPix[i+5] = uint8(c[2])
	}
	return
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *RGB48Image) SubImage(r image.Rectangle) image.Image {
	r = r.Intersect(p.XRect)
	// If r1 and r2 are Rectangles, r1.Intersect(r2) is not guaranteed to be inside
	// either r1 or r2 if the intersection is empty. Without explicitly checking for
	// this, the XPix[i:] expression below can panic.
	if r.Empty() {
		return &RGB48Image{}
	}
	i := p.PixOffset(r.Min.X, r.Min.Y)
	return &RGB48Image{
		XPix:    p.XPix[i:],
		XStride: p.XStride,
		XRect:   r,
	}
}

// Opaque scans the entire image and reports whether it is fully opaque.
func (p *RGB48Image) Opaque() bool {
	return true
}

// NewRGB48Image returns a new RGB48Image with the given bounds.
func NewRGB48Image(r image.Rectangle) *RGB48Image {
	w, h := r.Dx(), r.Dy()
	pix := make([]uint8, 6*w*h)
	return &RGB48Image{
		XPix:    pix,
		XStride: 6 * w,
		XRect:   r,
	}
}

func NewRGB48ImageFrom(m image.Image) *RGB48Image {
	if m, ok := m.(*RGB48Image); ok {
		return m
	}

	// convert to RGB48Image
	b := m.Bounds()
	rgb := NewRGB48Image(b)
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			pr, pg, pb, _ := m.At(x, y).RGBA()
			rgb.SetRGB48(x, y, [3]uint16{
				uint16(pr),
				uint16(pg),
				uint16(pb),
			})
		}
	}
	return rgb
}
