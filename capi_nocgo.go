// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !cgo

package webp

import (
	"errors"
)

func webpGetInfo(data []byte) (width, height int, has_alpha bool, err error) {
	err = errors.New("webpGetInfo: cgo is disabled!")
	return
}

func webpDecodeGray(data []byte) (pix []byte, width, height int, err error) {
	err = errors.New("webpDecodeGray: cgo is disabled!")
	return
}

func webpDecodeRGB(data []byte) (pix []byte, width, height int, err error) {
	err = errors.New("webpDecodeRGB: cgo is disabled!")
	return
}

func webpDecodeRGBA(data []byte) (pix []byte, width, height int, err error) {
	err = errors.New("webpDecodeRGBA: cgo is disabled!")
	return
}

func webpEncodeGray(
	pix []byte, width, height, stride int,
	quality_factor float32,
) (output []byte, err error) {
	err = errors.New("webpEncodeGray: cgo is disabled!")
	return
}

func webpEncodeRGB(
	pix []byte, width, height, stride int,
	quality_factor float32,
) (output []byte, err error) {
	err = errors.New("webpEncodeRGB: cgo is disabled!")
	return
}

func webpEncodeRGBA(
	pix []byte, width, height, stride int,
	quality_factor float32,
) (output []byte, err error) {
	err = errors.New("webpEncodeRGBA: cgo is disabled!")
	return
}

func webpEncodeLosslessGray(
	pix []byte, width, height, stride int,
) (output []byte, err error) {
	err = errors.New("webpEncodeLosslessGray: cgo is disabled!")
	return
}

func webpEncodeLosslessRGB(
	pix []byte, width, height, stride int,
) (output []byte, err error) {
	err = errors.New("webpEncodeLosslessRGB: cgo is disabled!")
	return
}

func webpEncodeLosslessRGBA(
	pix []byte, width, height, stride int,
) (output []byte, err error) {
	err = errors.New("webpEncodeLosslessRGBA: cgo is disabled!")
	return
}

func webpGetEXIF(data []byte) (metadata []byte, err error) {
	err = errors.New("webpGetEXIF: cgo is disabled!")
	return
}
func webpGetICCP(data []byte) (metadata []byte, err error) {
	err = errors.New("webpGetICCP: cgo is disabled!")
	return
}
func webpGetXMP(data []byte) (metadata []byte, err error) {
	err = errors.New("webpGetXMP: cgo is disabled!")
	return
}
func webpGetMetadata(data []byte, format string) (metadata []byte, err error) {
	err = errors.New("webpGetMetadata: cgo is disabled!")
	return
}

func webpSetEXIF(data, metadata []byte) (newData []byte, err error) {
	err = errors.New("webpSetEXIF: cgo is disabled!")
	return
}
func webpSetICCP(data, metadata []byte) (newData []byte, err error) {
	err = errors.New("webpSetICCP: cgo is disabled!")
	return
}
func webpSetXMP(data, metadata []byte) (newData []byte, err error) {
	err = errors.New("webpSetXMP: cgo is disabled!")
	return
}
func webpSetMetadata(data, metadata []byte, format string) (newData []byte, err error) {
	err = errors.New("webpSetMetadata: cgo is disabled!")
	return
}

func webpDelEXIF(data []byte) (newData []byte, err error) {
	err = errors.New("webpDelEXIF: cgo is disabled!")
	return
}
func webpDelICCP(data []byte) (newData []byte, err error) {
	err = errors.New("webpDelICCP: cgo is disabled!")
	return
}
func webpDelXMP(data []byte) (newData []byte, err error) {
	err = errors.New("webpDelXMP: cgo is disabled!")
	return
}
