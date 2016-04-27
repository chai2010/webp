// Copyright 2015 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package webp

import (
	"image"
	"image/color"
	"reflect"
	"runtime"
	"unsafe"
)

const (
	MemPMagic = "MemP" // See https://github.com/chai2010/image
)

const (
	isLittleEndian = (runtime.GOARCH == "386" ||
		runtime.GOARCH == "amd64" ||
		runtime.GOARCH == "arm" ||
		runtime.GOARCH == "arm64")
)

var (
	_ image.Image = (*MemPImage)(nil)
	_ MemP        = (*MemPImage)(nil)
)

// MemP Image Spec (Native Endian), see https://github.com/chai2010/image.
type MemP interface {
	MemPMagic() string
	Bounds() image.Rectangle
	Channels() int
	DataType() reflect.Kind
	Pix() []byte // PixSilce type

	// Stride is the Pix stride (in bytes, must align with SizeofKind(p.DataType))
	// between vertically adjacent pixels.
	Stride() int
}

type MemPImage struct {
	XMemPMagic string // MemP
	XRect      image.Rectangle
	XChannels  int
	XDataType  reflect.Kind
	XPix       PixSlice
	XStride    int
}

func NewMemPImage(r image.Rectangle, channels int, dataType reflect.Kind) *MemPImage {
	m := &MemPImage{
		XMemPMagic: MemPMagic,
		XRect:      r,
		XStride:    r.Dx() * channels * SizeofKind(dataType),
		XChannels:  channels,
		XDataType:  dataType,
	}
	m.XPix = make([]byte, r.Dy()*m.XStride)
	return m
}

// m is MemP or image.Image
func AsMemPImage(m interface{}) (p *MemPImage, ok bool) {
	if m, ok := m.(*MemPImage); ok {
		return m, true
	}
	if m, ok := m.(MemP); ok {
		return &MemPImage{
			XMemPMagic: MemPMagic,
			XRect:      m.Bounds(),
			XChannels:  m.Channels(),
			XDataType:  m.DataType(),
			XPix:       m.Pix(),
			XStride:    m.Stride(),
		}, true
	}
	if m, ok := m.(*image.Gray); ok {
		return &MemPImage{
			XMemPMagic: MemPMagic,
			XRect:      m.Bounds(),
			XChannels:  1,
			XDataType:  reflect.Uint8,
			XPix:       m.Pix,
			XStride:    m.Stride,
		}, true
	}
	if m, ok := m.(*image.RGBA); ok {
		return &MemPImage{
			XMemPMagic: MemPMagic,
			XRect:      m.Bounds(),
			XChannels:  4,
			XDataType:  reflect.Uint8,
			XPix:       m.Pix,
			XStride:    m.Stride,
		}, true
	}
	return nil, false
}

func NewMemPImageFrom(m image.Image) *MemPImage {
	if p, ok := m.(*MemPImage); ok {
		return p.Clone()
	}
	if p, ok := AsMemPImage(m); ok {
		return p.Clone()
	}

	switch m := m.(type) {
	case *image.Gray:
		b := m.Bounds()
		p := NewMemPImage(b, 1, reflect.Uint8)

		for y := b.Min.Y; y < b.Max.Y; y++ {
			off0 := m.PixOffset(0, y)
			off1 := p.PixOffset(0, y)
			copy(p.XPix[off1:][:p.XStride], m.Pix[off0:][:m.Stride])
			off0 += m.Stride
			off1 += p.XStride
		}
		return p

	case *image.Gray16:
		b := m.Bounds()
		p := NewMemPImage(b, 1, reflect.Uint16)

		for y := b.Min.Y; y < b.Max.Y; y++ {
			off0 := m.PixOffset(0, y)
			off1 := p.PixOffset(0, y)
			copy(p.XPix[off1:][:p.XStride], m.Pix[off0:][:m.Stride])
			off0 += m.Stride
			off1 += p.XStride
		}
		if isLittleEndian {
			p.XPix.SwapEndian(p.XDataType)
		}
		return p

	case *image.RGBA:
		b := m.Bounds()
		p := NewMemPImage(b, 4, reflect.Uint8)

		for y := b.Min.Y; y < b.Max.Y; y++ {
			off0 := m.PixOffset(0, y)
			off1 := p.PixOffset(0, y)
			copy(p.XPix[off1:][:p.XStride], m.Pix[off0:][:m.Stride])
			off0 += m.Stride
			off1 += p.XStride
		}
		return p

	case *image.RGBA64:
		b := m.Bounds()
		p := NewMemPImage(b, 4, reflect.Uint16)

		for y := b.Min.Y; y < b.Max.Y; y++ {
			off0 := m.PixOffset(0, y)
			off1 := p.PixOffset(0, y)
			copy(p.XPix[off1:][:p.XStride], m.Pix[off0:][:m.Stride])
			off0 += m.Stride
			off1 += p.XStride
		}
		if isLittleEndian {
			p.XPix.SwapEndian(p.XDataType)
		}
		return p

	case *image.YCbCr:
		b := m.Bounds()
		p := NewMemPImage(b, 4, reflect.Uint8)
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				R, G, B, A := m.At(x, y).RGBA()

				i := p.PixOffset(x, y)
				p.XPix[i+0] = uint8(R >> 8)
				p.XPix[i+1] = uint8(G >> 8)
				p.XPix[i+2] = uint8(B >> 8)
				p.XPix[i+3] = uint8(A >> 8)
			}
		}
		return p

	default:
		b := m.Bounds()
		p := NewMemPImage(b, 4, reflect.Uint16)
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				R, G, B, A := m.At(x, y).RGBA()

				i := p.PixOffset(x, y)
				p.XPix[i+0] = uint8(R >> 8)
				p.XPix[i+1] = uint8(R)
				p.XPix[i+2] = uint8(G >> 8)
				p.XPix[i+3] = uint8(G)
				p.XPix[i+4] = uint8(B >> 8)
				p.XPix[i+5] = uint8(B)
				p.XPix[i+6] = uint8(A >> 8)
				p.XPix[i+7] = uint8(A)
			}
		}
		return p
	}
}

func (p *MemPImage) Clone() *MemPImage {
	q := new(MemPImage)
	*q = *p
	q.XPix = append([]byte(nil), p.XPix...)
	return q
}

func (p *MemPImage) MemPMagic() string {
	return p.XMemPMagic
}

func (p *MemPImage) Bounds() image.Rectangle {
	return p.XRect
}

func (p *MemPImage) Channels() int {
	return p.XChannels
}

func (p *MemPImage) DataType() reflect.Kind {
	return p.XDataType
}

func (p *MemPImage) Pix() []byte {
	return p.XPix
}

func (p *MemPImage) Stride() int {
	return p.XStride
}

func (p *MemPImage) ColorModel() color.Model {
	return ColorModel(p.XChannels, p.XDataType)
}

func (p *MemPImage) At(x, y int) color.Color {
	if !(image.Point{x, y}.In(p.XRect)) {
		return MemPColor{
			Channels: p.XChannels,
			DataType: p.XDataType,
		}
	}
	i := p.PixOffset(x, y)
	n := SizeofPixel(p.XChannels, p.XDataType)
	return MemPColor{
		Channels: p.XChannels,
		DataType: p.XDataType,
		Pix:      p.XPix[i:][:n],
	}
}

func (p *MemPImage) PixelAt(x, y int) []byte {
	if !(image.Point{x, y}.In(p.XRect)) {
		return nil
	}
	i := p.PixOffset(x, y)
	n := SizeofPixel(p.XChannels, p.XDataType)
	return p.XPix[i:][:n]
}

func (p *MemPImage) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(p.XRect)) {
		return
	}
	i := p.PixOffset(x, y)
	n := SizeofPixel(p.XChannels, p.XDataType)
	v := p.ColorModel().Convert(c).(MemPColor)
	copy(p.XPix[i:][:n], v.Pix)
}

func (p *MemPImage) SetPixel(x, y int, c []byte) {
	if !(image.Point{x, y}.In(p.XRect)) {
		return
	}
	i := p.PixOffset(x, y)
	n := SizeofPixel(p.XChannels, p.XDataType)
	copy(p.XPix[i:][:n], c)
}

func (p *MemPImage) PixOffset(x, y int) int {
	return (y-p.XRect.Min.Y)*p.XStride + (x-p.XRect.Min.X)*SizeofPixel(p.XChannels, p.XDataType)
}

func (p *MemPImage) SubImage(r image.Rectangle) image.Image {
	r = r.Intersect(p.XRect)
	// If r1 and r2 are Rectangles, r1.Intersect(r2) is not guaranteed to be inside
	// either r1 or r2 if the intersection is empty. Without explicitly checking for
	// this, the Pix[i:] expression below can panic.
	if r.Empty() {
		return &MemPImage{}
	}
	i := p.PixOffset(r.Min.X, r.Min.Y)
	return &MemPImage{
		XRect:     r,
		XChannels: p.XChannels,
		XDataType: p.XDataType,
		XPix:      p.XPix[i:],
		XStride:   p.XStride,
	}
}

func (p *MemPImage) AsStdImage() (m image.Image, ok bool) {
	switch {
	case p.XChannels == 1 && p.XDataType == reflect.Uint8:
		return &image.Gray{
			Pix:    p.XPix,
			Stride: p.XStride,
			Rect:   p.XRect,
		}, true
	case p.XChannels == 4 && p.XDataType == reflect.Uint8:
		return &image.RGBA{
			Pix:    p.XPix,
			Stride: p.XStride,
			Rect:   p.XRect,
		}, true
	default:
		return nil, false
	}
}

func (p *MemPImage) StdImage() image.Image {
	switch {
	case p.XChannels == 1 && p.XDataType == reflect.Uint8:
		return &image.Gray{
			Pix:    p.XPix,
			Stride: p.XStride,
			Rect:   p.XRect,
		}
	case p.XChannels == 1 && p.XDataType == reflect.Uint16:
		m := &image.Gray16{
			Pix:    p.XPix,
			Stride: p.XStride,
			Rect:   p.XRect,
		}
		if isLittleEndian {
			m.Pix = append([]byte(nil), m.Pix...)
			PixSlice(m.Pix).SwapEndian(p.XDataType)
		}
		return m
	case p.XChannels == 4 && p.XDataType == reflect.Uint8:
		return &image.RGBA{
			Pix:    p.XPix,
			Stride: p.XStride,
			Rect:   p.XRect,
		}
	case p.XChannels == 4 && p.XDataType == reflect.Uint16:
		m := &image.RGBA64{
			Pix:    p.XPix,
			Stride: p.XStride,
			Rect:   p.XRect,
		}
		if isLittleEndian {
			m.Pix = append([]byte(nil), m.Pix...)
			PixSlice(m.Pix).SwapEndian(p.XDataType)
		}
		return m
	}

	return p
}

func ChannelsOf(m image.Image) int {
	if m, ok := AsMemPImage(m); ok {
		return m.XChannels
	}
	switch m.(type) {
	case *image.Gray:
		return 1
	case *image.Gray16:
		return 1
	case *image.YCbCr:
		return 3
	}
	return 4
}

func DepthOf(m image.Image) int {
	if m, ok := m.(*MemPImage); ok {
		return SizeofKind(m.XDataType) * 8
	}
	if m, ok := m.(MemP); ok {
		return SizeofKind(m.DataType() * 8)
	}
	switch m.(type) {
	case *image.Gray:
		return 1 * 8
	case *image.Gray16:
		return 2 * 8
	case *image.NRGBA:
		return 1 * 8
	case *image.NRGBA64:
		return 2 * 8
	case *image.RGBA:
		return 1 * 8
	case *image.RGBA64:
		return 2 * 8
	case *image.YCbCr:
		return 1 * 8
	}
	return 2 * 8
}

type SizeofImager interface {
	SizeofImage() int
}

func SizeofImage(m image.Image) int {
	if m, ok := m.(SizeofImager); ok {
		return m.SizeofImage()
	}
	if m, ok := AsMemPImage(m); ok {
		return int(unsafe.Sizeof(*m)) + len(m.XPix)
	}

	b := m.Bounds()
	switch m := m.(type) {
	case *image.Alpha:
		return int(unsafe.Sizeof(*m)) + b.Dx()*b.Dy()*1
	case *image.Alpha16:
		return int(unsafe.Sizeof(*m)) + b.Dx()*b.Dy()*2
	case *image.Gray:
		return int(unsafe.Sizeof(*m)) + b.Dx()*b.Dy()*1
	case *image.Gray16:
		return int(unsafe.Sizeof(*m)) + b.Dx()*b.Dy()*2
	case *image.NRGBA:
		return int(unsafe.Sizeof(*m)) + b.Dx()*b.Dy()*4
	case *image.NRGBA64:
		return int(unsafe.Sizeof(*m)) + b.Dx()*b.Dy()*8
	case *image.RGBA:
		return int(unsafe.Sizeof(*m)) + b.Dx()*b.Dy()*4
	case *image.RGBA64:
		return int(unsafe.Sizeof(*m)) + b.Dx()*b.Dy()*8
	case *image.Uniform:
		return int(unsafe.Sizeof(*m))
	case *image.YCbCr:
		return int(unsafe.Sizeof(*m)) + len(m.Y) + len(m.Cb) + len(m.Cr)
	}

	// return same as RGBA64 size
	return int(unsafe.Sizeof((*image.RGBA64)(nil))) + b.Dx()*b.Dy()*8
}
