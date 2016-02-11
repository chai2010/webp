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

func DecodeRGBEx(data []byte, cbuf CBuffer) (m *RGBImage, pix CBuffer, err error) {
	if cbuf == nil {
		cbuf = nilCBuffer
	}
	pix, w, h, err := webpDecodeRGB(data, cbuf)
	if err != nil {
		return
	}
	m = &RGBImage{
		XPix:    pix.CData(),
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

func DecodeRGBAEx(data []byte, cbuf CBuffer) (m *image.RGBA, pix CBuffer, err error) {
	if cbuf == nil {
		cbuf = nilCBuffer
	}
	pix, w, h, err := webpDecodeRGBA(data, cbuf)
	if err != nil {
		return
	}
	m = &image.RGBA{
		Pix:    pix.CData(),
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

func EncodeGrayEx(m image.Image, quality float32, cbuf CBuffer) (data CBuffer, err error) {
	if cbuf == nil {
		cbuf = nilCBuffer
	}
	p := toGrayImage(m)
	data, err = webpEncodeGray(p.Pix, p.Rect.Dx(), p.Rect.Dy(), p.Stride, quality, cbuf)
	if err != nil {
		return
	}
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

func EncodeRGBEx(m image.Image, quality float32, cbuf CBuffer) (data CBuffer, err error) {
	if cbuf == nil {
		cbuf = nilCBuffer
	}
	p := NewRGBImageFrom(m)
	data, err = webpEncodeRGB(p.XPix, p.XRect.Dx(), p.XRect.Dy(), p.XStride, quality, cbuf)
	if err != nil {
		return
	}
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

func EncodeRGBAEx(m image.Image, quality float32, cbuf CBuffer) (data CBuffer, err error) {
	if cbuf == nil {
		cbuf = nilCBuffer
	}
	p := toRGBAImage(m)
	data, err = webpEncodeRGBA(p.Pix, p.Rect.Dx(), p.Rect.Dy(), p.Stride, quality, cbuf)
	if err != nil {
		return
	}
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

func EncodeLosslessGrayEx(m image.Image, cbuf CBuffer) (data CBuffer, err error) {
	if cbuf == nil {
		cbuf = nilCBuffer
	}
	p := toGrayImage(m)
	data, err = webpEncodeLosslessGray(p.Pix, p.Rect.Dx(), p.Rect.Dy(), p.Stride, cbuf)
	if err != nil {
		return
	}
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

func EncodeLosslessRGBEx(m image.Image, cbuf CBuffer) (data CBuffer, err error) {
	if cbuf == nil {
		cbuf = nilCBuffer
	}
	p := NewRGBImageFrom(m)
	data, err = webpEncodeLosslessRGB(p.XPix, p.XRect.Dx(), p.XRect.Dy(), p.XStride, cbuf)
	if err != nil {
		return
	}
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

func EncodeLosslessRGBAEx(m image.Image, cbuf CBuffer) (data CBuffer, err error) {
	if cbuf == nil {
		cbuf = nilCBuffer
	}
	p := toRGBAImage(m)
	data, err = webpEncodeLosslessRGBA(p.Pix, p.Rect.Dx(), p.Rect.Dy(), p.Stride, cbuf)
	if err != nil {
		return
	}
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
