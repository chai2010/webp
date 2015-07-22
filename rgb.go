// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package webp

import (
	"image"
	"image/color"
)

var (
	_ color.Color = (*RGBColor)(nil)
	_ image.Image = (*RGBImage)(nil)
)

var RGBModel color.Model = color.ModelFunc(rgbModel)

type RGBColor struct {
	R, G, B uint8
}

func (c RGBColor) RGBA() (r, g, b, a uint32) {
	r = uint32(c.R)
	r |= r << 8
	g = uint32(c.G)
	g |= g << 8
	b = uint32(c.B)
	b |= b << 8
	a = 0xffff
	return
}

func rgbModel(c color.Color) color.Color {
	if _, ok := c.(RGBColor); ok {
		return c
	}
	r, g, b, _ := c.RGBA()
	return RGBColor{R: uint8(r >> 8), G: uint8(g >> 8), B: uint8(b >> 8)}
}

type RGBImage struct {
	Pix    []uint8
	Stride int
	Rect   image.Rectangle
}

func (p *RGBImage) ColorModel() color.Model { return color.RGBAModel }

func (p *RGBImage) Bounds() image.Rectangle { return p.Rect }

func (p *RGBImage) At(x, y int) color.Color {
	return p.RGBAt(x, y)
}

func (p *RGBImage) RGBAt(x, y int) RGBColor {
	if !(image.Point{x, y}.In(p.Rect)) {
		return RGBColor{}
	}
	i := p.PixOffset(x, y)
	return RGBColor{
		R: p.Pix[i+0],
		G: p.Pix[i+1],
		B: p.Pix[i+2],
	}
}

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (p *RGBImage) PixOffset(x, y int) int {
	return (y-p.Rect.Min.Y)*p.Stride + (x-p.Rect.Min.X)*3
}

func (p *RGBImage) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(p.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	c1 := RGBModel.Convert(c).(RGBColor)
	p.Pix[i+0] = c1.R
	p.Pix[i+1] = c1.G
	p.Pix[i+2] = c1.B
	return
}

func (p *RGBImage) SetRGB(x, y int, c RGBColor) {
	if !(image.Point{x, y}.In(p.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	p.Pix[i+0] = c.R
	p.Pix[i+1] = c.G
	p.Pix[i+2] = c.B
	return
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *RGBImage) SubImage(r image.Rectangle) image.Image {
	r = r.Intersect(p.Rect)
	// If r1 and r2 are Rectangles, r1.Intersect(r2) is not guaranteed to be inside
	// either r1 or r2 if the intersection is empty. Without explicitly checking for
	// this, the Pix[i:] expression below can panic.
	if r.Empty() {
		return &RGBImage{}
	}
	i := p.PixOffset(r.Min.X, r.Min.Y)
	return &RGBImage{
		Pix:    p.Pix[i:],
		Stride: p.Stride,
		Rect:   r,
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
		Pix:    pix,
		Stride: 3 * w,
		Rect:   r,
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
			rgb.SetRGB(x, y, RGBColor{
				R: uint8(pr >> 8),
				G: uint8(pg >> 8),
				B: uint8(pb >> 8),
			})
		}
	}
	return rgb
}
