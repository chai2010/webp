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
	_ image.Image = (*RGBImage)(nil)
	_ MemP        = (*RGBImage)(nil)
)

type RGBImage struct {
	XPix    []uint8
	XStride int
	XRect   image.Rectangle
}

func (p *RGBImage) MemPMagic() string {
	return MemPMagic
}

func (p *RGBImage) Bounds() image.Rectangle {
	return p.XRect
}

func (p *RGBImage) Channels() int {
	return 3
}

func (p *RGBImage) DataType() reflect.Kind {
	return reflect.Uint8
}

func (p *RGBImage) Pix() []byte {
	return p.XPix
}

func (p *RGBImage) Stride() int {
	return p.XStride
}

func (p *RGBImage) ColorModel() color.Model { return color.RGBAModel }

func (p *RGBImage) At(x, y int) color.Color {
	if !(image.Point{x, y}.In(p.XRect)) {
		return color.RGBA{}
	}
	i := p.PixOffset(x, y)
	return color.RGBA{
		R: p.XPix[i+0],
		G: p.XPix[i+1],
		B: p.XPix[i+2],
		A: 0xff,
	}
}

func (p *RGBImage) RGBAt(x, y int) [3]uint8 {
	if !(image.Point{x, y}.In(p.XRect)) {
		return [3]uint8{}
	}
	i := p.PixOffset(x, y)
	return [3]uint8{
		p.XPix[i+0],
		p.XPix[i+1],
		p.XPix[i+2],
	}
}

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (p *RGBImage) PixOffset(x, y int) int {
	return (y-p.XRect.Min.Y)*p.XStride + (x-p.XRect.Min.X)*3
}

func (p *RGBImage) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(p.XRect)) {
		return
	}
	i := p.PixOffset(x, y)
	c1 := color.RGBAModel.Convert(c).(color.RGBA)
	p.XPix[i+0] = c1.R
	p.XPix[i+1] = c1.G
	p.XPix[i+2] = c1.B
	return
}

func (p *RGBImage) SetRGB(x, y int, c [3]uint8) {
	if !(image.Point{x, y}.In(p.XRect)) {
		return
	}
	i := p.PixOffset(x, y)
	p.XPix[i+0] = c[0]
	p.XPix[i+1] = c[1]
	p.XPix[i+2] = c[2]
	return
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *RGBImage) SubImage(r image.Rectangle) image.Image {
	r = r.Intersect(p.XRect)
	// If r1 and r2 are Rectangles, r1.Intersect(r2) is not guaranteed to be inside
	// either r1 or r2 if the intersection is empty. Without explicitly checking for
	// this, the Pix[i:] expression below can panic.
	if r.Empty() {
		return &RGBImage{}
	}
	i := p.PixOffset(r.Min.X, r.Min.Y)
	return &RGBImage{
		XPix:    p.XPix[i:],
		XStride: p.XStride,
		XRect:   r,
	}
}

// Opaque scans the entire image and reports whether it is fully opaque.
func (p *RGBImage) Opaque() bool {
	return true
}

// NewRGBImage returns a new RGBImage with the given bounds.
func NewRGBImage(r image.Rectangle) *RGBImage {
	w, h := r.Dx(), r.Dy()
	pix := make([]uint8, 3*w*h)
	return &RGBImage{
		XPix:    pix,
		XStride: 3 * w,
		XRect:   r,
	}
}

func NewRGBImageFrom(m image.Image) *RGBImage {
	if m, ok := m.(*RGBImage); ok {
		return m
	}

	// convert to RGBImage
	b := m.Bounds()
	rgb := NewRGBImage(b)
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			pr, pg, pb, _ := m.At(x, y).RGBA()
			rgb.SetRGB(x, y, [3]uint8{
				uint8(pr >> 8),
				uint8(pg >> 8),
				uint8(pb >> 8),
			})
		}
	}
	return rgb
}
