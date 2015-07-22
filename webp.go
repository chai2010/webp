// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package webp

import (
	"image"
	"strings"
)

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
		Pix:    pix,
		Stride: 3 * w,
		Rect:   image.Rect(0, 0, w, h),
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

func EncodeGray(m image.Image, quality float32) (data []byte, err error) {
	p := toGrayImage(m)
	return webpEncodeGray(p.Pix, p.Rect.Dx(), p.Rect.Dy(), p.Stride, quality)
}

func EncodeRGB(m image.Image, quality float32) (data []byte, err error) {
	p := NewRGBImageFrom(m)
	return webpEncodeRGB(p.Pix, p.Rect.Dx(), p.Rect.Dy(), p.Stride, quality)
}

func EncodeRGBA(m image.Image, quality float32) (data []byte, err error) {
	p := toRGBAImage(m)
	return webpEncodeRGBA(p.Pix, p.Rect.Dx(), p.Rect.Dy(), p.Stride, quality)
}

func EncodeLosslessGray(m image.Image) (data []byte, err error) {
	p := toGrayImage(m)
	return webpEncodeLosslessGray(p.Pix, p.Rect.Dx(), p.Rect.Dy(), p.Stride)
}

func EncodeLosslessRGB(m image.Image) (data []byte, err error) {
	p := NewRGBImageFrom(m)
	return webpEncodeLosslessRGB(p.Pix, p.Rect.Dx(), p.Rect.Dy(), p.Stride)
}

func EncodeLosslessRGBA(m image.Image) (data []byte, err error) {
	p := toRGBAImage(m)
	return webpEncodeLosslessRGBA(p.Pix, p.Rect.Dx(), p.Rect.Dy(), p.Stride)
}

// GetMetadata return EXIF/ICCP/XMP format metadata.
func GetMetadata(data []byte, format string) (metadata []byte, err error) {
	return webpGetMetadata(data, strings.ToUpper(format))
}

// SetMetadata set EXIF/ICCP/XMP format metadata.
func SetMetadata(data, metadata []byte, format string) (newData []byte, err error) {
	return webpSetMetadata(data, metadata, format)
}
