// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package webp

import (
	"image"
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

func DecodeRGB(data []byte) (m image.Image, err error) {
	pix, w, h, err := webpDecodeRGB(data)
	if err != nil {
		return
	}
	m = &_RGB{
		_Pix:    pix,
		_Stride: 3 * w,
		_Rect:   image.Rect(0, 0, w, h),
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

func EncodeGray(m *image.Gray, quality float32) (data []byte, err error) {
	return webpEncodeGray(m.Pix, m.Rect.Dx(), m.Rect.Dy(), m.Stride, quality)
}

func EncodeRGB(m image.Image, quality float32) (data []byte, err error) {
	p := newRGBFromImage(m)
	return webpEncodeRGB(p._Pix, p._Rect.Dx(), p._Rect.Dy(), p._Stride, quality)
}

func EncodeRGBA(m *image.RGBA, quality float32) (data []byte, err error) {
	return webpEncodeRGBA(m.Pix, m.Rect.Dx(), m.Rect.Dy(), m.Stride, quality)
}

func EncodeLosslessGray(m *image.Gray) (data []byte, err error) {
	return webpEncodeLosslessGray(m.Pix, m.Rect.Dx(), m.Rect.Dy(), m.Stride)
}

func EncodeLosslessRGB(m image.Image) (data []byte, err error) {
	p := newRGBFromImage(m)
	return webpEncodeLosslessRGB(p._Pix, p._Rect.Dx(), p._Rect.Dy(), p._Stride)
}

func EncodeLosslessRGBA(m *image.RGBA) (data []byte, err error) {
	return webpEncodeLosslessRGBA(m.Pix, m.Rect.Dx(), m.Rect.Dy(), m.Stride)
}
