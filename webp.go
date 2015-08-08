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

var (
	nilCBuffer = NewCBuffer(0)
)

func GetInfo(data []byte) (width, height int, hasAlpha bool, err error) {
	return webpGetInfo(data, nilCBuffer)
}

func DecodeGray(data []byte) (m *image.Gray, err error) {
	pix, w, h, err := webpDecodeGray(data, nilCBuffer)
	if err != nil {
		return
	}
	defer pix.Close()
	m = &image.Gray{
		Pix:    append([]byte{}, pix.CData()...),
		Stride: 1 * w,
		Rect:   image.Rect(0, 0, w, h),
	}
	return
}

func DecodeRGB(data []byte) (m *RGBImage, err error) {
	pix, w, h, err := webpDecodeRGB(data, nilCBuffer)
	if err != nil {
		return
	}
	defer pix.Close()
	m = &RGBImage{
		XPix:    append([]byte{}, pix.CData()...),
		XStride: 3 * w,
		XRect:   image.Rect(0, 0, w, h),
	}
	return
}

func DecodeRGBA(data []byte) (m *image.RGBA, err error) {
	pix, w, h, err := webpDecodeRGBA(data, nilCBuffer)
	if err != nil {
		return
	}
	defer pix.Close()
	m = &image.RGBA{
		Pix:    append([]byte{}, pix.CData()...),
		Stride: 4 * w,
		Rect:   image.Rect(0, 0, w, h),
	}
	return
}

func EncodeGray(m image.Image, quality float32) (data []byte, err error) {
	p := toGrayImage(m)
	cbuf, err := webpEncodeGray(p.Pix, p.Rect.Dx(), p.Rect.Dy(), p.Stride, quality, nilCBuffer)
	if err != nil {
		return
	}
	data = append([]byte{}, cbuf.CData()...)
	cbuf.Close()
	return
}

func EncodeRGB(m image.Image, quality float32) (data []byte, err error) {
	p := NewRGBImageFrom(m)
	cbuf, err := webpEncodeRGB(p.XPix, p.XRect.Dx(), p.XRect.Dy(), p.XStride, quality, nilCBuffer)
	if err != nil {
		return
	}
	data = append([]byte{}, cbuf.CData()...)
	cbuf.Close()
	return
}

func EncodeRGBA(m image.Image, quality float32) (data []byte, err error) {
	p := toRGBAImage(m)
	cbuf, err := webpEncodeRGBA(p.Pix, p.Rect.Dx(), p.Rect.Dy(), p.Stride, quality, nilCBuffer)
	if err != nil {
		return
	}
	data = append([]byte{}, cbuf.CData()...)
	cbuf.Close()
	return
}

func EncodeLosslessGray(m image.Image) (data []byte, err error) {
	p := toGrayImage(m)
	cbuf, err := webpEncodeLosslessGray(p.Pix, p.Rect.Dx(), p.Rect.Dy(), p.Stride, nilCBuffer)
	if err != nil {
		return
	}
	data = append([]byte{}, cbuf.CData()...)
	cbuf.Close()
	return
}

func EncodeLosslessRGB(m image.Image) (data []byte, err error) {
	p := NewRGBImageFrom(m)
	cbuf, err := webpEncodeLosslessRGB(p.XPix, p.XRect.Dx(), p.XRect.Dy(), p.XStride, nilCBuffer)
	if err != nil {
		return
	}
	data = append([]byte{}, cbuf.CData()...)
	cbuf.Close()
	return
}

func EncodeLosslessRGBA(m image.Image) (data []byte, err error) {
	p := toRGBAImage(m)
	cbuf, err := webpEncodeLosslessRGBA(p.Pix, p.Rect.Dx(), p.Rect.Dy(), p.Stride, nilCBuffer)
	if err != nil {
		return
	}
	data = append([]byte{}, cbuf.CData()...)
	cbuf.Close()
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
