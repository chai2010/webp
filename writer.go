// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build cgo

package webp

import (
	"image"
	"image/color"
	"io"
)

const DefaulQuality = 90

// Options are the encoding parameters.
type Options struct {
	Lossless bool
	Quality  float32 // 0 ~ 100
}

type colorModeler interface {
	ColorModel() color.Model
}

// Encode writes the image m to w in WEBP format.
func Encode(w io.Writer, m image.Image, opt *Options) (err error) {
	var output []byte
	if opt != nil && opt.Lossless {
		switch m := adjustImage(m).(type) {
		case *image.Gray:
			if output, err = EncodeLosslessGray(m); err != nil {
				return
			}
		case *RGBImage:
			if output, err = EncodeLosslessRGB(m); err != nil {
				return
			}
		case *image.RGBA:
			if output, err = EncodeLosslessRGBA(m); err != nil {
				return
			}
		default:
			panic("image/webp: Encode, unreachable!")
		}
	} else {
		quality := float32(DefaulQuality)
		if opt != nil {
			quality = opt.Quality
		}
		switch m := adjustImage(m).(type) {
		case *image.Gray:
			if output, err = EncodeGray(m, quality); err != nil {
				return
			}
		case *RGBImage:
			if output, err = EncodeRGB(m, quality); err != nil {
				return
			}
		case *image.RGBA:
			if output, err = EncodeRGBA(m, quality); err != nil {
				return
			}
		default:
			panic("image/webp: Encode, unreachable!")
		}
	}
	_, err = w.Write(output)
	return
}

func adjustImage(m image.Image) image.Image {
	switch m := m.(type) {
	case *image.Gray:
		return m
	case *RGBImage:
		return m
	case *image.RGBA:
		return m
	case *image.YCbCr:
		return NewRGBImageFrom(m)

	case *image.Gray16:
		return toGrayImage(m)
	case *image.RGBA64:
		return toRGBAImage(m)
	case *image.NRGBA:
		return toRGBAImage(m)
	case *image.NRGBA64:
		return toRGBAImage(m)

	default:
		return toRGBAImage(m)
	}
}

func toGrayImage(m image.Image) *image.Gray {
	if m, ok := m.(*image.Gray); ok {
		return m
	}
	b := m.Bounds()
	gray := image.NewGray(b)
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			c := color.GrayModel.Convert(m.At(x, y)).(color.Gray)
			gray.SetGray(x, y, c)
		}
	}
	return gray
}

func toRGBAImage(m image.Image) *image.RGBA {
	if m, ok := m.(*image.RGBA); ok {
		return m
	}
	b := m.Bounds()
	rgba := image.NewRGBA(b)
	dstColorRGBA64 := &color.RGBA64{}
	dstColor := color.Color(dstColorRGBA64)
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			pr, pg, pb, pa := m.At(x, y).RGBA()
			dstColorRGBA64.R = uint16(pr)
			dstColorRGBA64.G = uint16(pg)
			dstColorRGBA64.B = uint16(pb)
			dstColorRGBA64.A = uint16(pa)
			rgba.Set(x, y, dstColor)
		}
	}
	return rgba
}
