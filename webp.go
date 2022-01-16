// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package webp

import (
	"image"
	"strings"

	"embed"
)

//go:embed internal
var _ embed.FS

const (
	maxWebpHeaderSize = 32
)

func GetInfo(data []byte) (width, height int, hasAlpha bool, err error) {
	return webpGetInfo(data)
}

func DecodeGray(data []byte) (m *image.Gray, err error) {
	pix, w, h, err := webpDecodeGray(data)
	if err != nil {
		return
	}
	m = &image.Gray{
		Pix:    pix,
		Stride: 1 * w,
		Rect:   image.Rect(0, 0, w, h),
	}
	return
}

func DecodeRGB(data []byte) (m *RGBImage, err error) {
	pix, w, h, err := webpDecodeRGB(data)
	if err != nil {
		return
	}
	m = &RGBImage{
		XPix:    pix,
		XStride: 3 * w,
		XRect:   image.Rect(0, 0, w, h),
	}
	return
}

func DecodeRGBA(data []byte) (m *image.RGBA, err error) {
	pix, w, h, err := webpDecodeRGBA(data)
	if err != nil {
		return
	}
	m = &image.RGBA{
		Pix:    pix,
		Stride: 4 * w,
		Rect:   image.Rect(0, 0, w, h),
	}
	return
}

// DecodeGrayToSize decodes a Gray image scaled to the given dimensions. For
// large images, the DecodeXXXToSize methods are significantly faster and
// require less memory compared to decoding a full-size image and then resizing it.
func DecodeGrayToSize(data []byte, width, height int) (m *image.Gray, err error) {
	pix, err := webpDecodeGrayToSize(data, width, height)
	if err != nil {
		return
	}
	m = &image.Gray{
		Pix:    pix,
		Stride: width,
		Rect:   image.Rect(0, 0, width, height),
	}
	return
}

// DecodeRGBToSize decodes an RGB image scaled to the given dimensions.
func DecodeRGBToSize(data []byte, width, height int) (m *RGBImage, err error) {
	pix, err := webpDecodeRGBToSize(data, width, height)
	if err != nil {
		return
	}
	m = &RGBImage{
		XPix:    pix,
		XStride: 3 * width,
		XRect:   image.Rect(0, 0, width, height),
	}
	return
}

// DecodeRGBAToSize decodes a Gray image scaled to the given dimensions.
func DecodeRGBAToSize(data []byte, width, height int) (m *image.RGBA, err error) {
	pix, err := webpDecodeRGBAToSize(data, width, height)
	if err != nil {
		return
	}
	m = &image.RGBA{
		Pix:    pix,
		Stride: 4 * width,
		Rect:   image.Rect(0, 0, width, height),
	}
	return
}

func EncodeGray(m image.Image, quality float32) (data []byte, err error) {
	p := toGrayImage(m)
	data, err = webpEncodeGray(p.Pix, p.Rect.Dx(), p.Rect.Dy(), p.Stride, quality)
	if err != nil {
		return
	}
	return
}

func EncodeRGB(m image.Image, quality float32) (data []byte, err error) {
	p := NewRGBImageFrom(m)
	data, err = webpEncodeRGB(p.XPix, p.XRect.Dx(), p.XRect.Dy(), p.XStride, quality)
	return
}

func EncodeRGBA(m image.Image, quality float32) (data []byte, err error) {
	p := toRGBAImage(m)
	data, err = webpEncodeRGBA(p.Pix, p.Rect.Dx(), p.Rect.Dy(), p.Stride, quality)
	return
}

func EncodeLosslessGray(m image.Image) (data []byte, err error) {
	p := toGrayImage(m)
	data, err = webpEncodeLosslessGray(p.Pix, p.Rect.Dx(), p.Rect.Dy(), p.Stride)
	return
}

func EncodeLosslessRGB(m image.Image) (data []byte, err error) {
	p := NewRGBImageFrom(m)
	data, err = webpEncodeLosslessRGB(p.XPix, p.XRect.Dx(), p.XRect.Dy(), p.XStride)
	return
}

func EncodeLosslessRGBA(m image.Image) (data []byte, err error) {
	p := toRGBAImage(m)
	data, err = webpEncodeLosslessRGBA(0, p.Pix, p.Rect.Dx(), p.Rect.Dy(), p.Stride)
	return
}

// EncodeExactLosslessRGBA Encode lossless RGB mode with exact.
// exact: preserve RGB values in transparent area.
func EncodeExactLosslessRGBA(m image.Image) (data []byte, err error) {
	p := toRGBAImage(m)
	data, err = webpEncodeLosslessRGBA(1, p.Pix, p.Rect.Dx(), p.Rect.Dy(), p.Stride)
	return
}

// GetMetadata return EXIF/ICCP/XMP format metadata.
func GetMetadata(data []byte, format string) (metadata []byte, err error) {
	return webpGetMetadata(data, strings.ToUpper(format))
}

// SetMetadata set EXIF/ICCP/XMP format metadata.
func SetMetadata(data, metadata []byte, format string) (newData []byte, err error) {
	return webpSetMetadata(data, metadata, format)
}
