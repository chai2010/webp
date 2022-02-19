// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build cgo
// +build cgo

package webp

import (
	"errors"
	"image"
	"image/color"
	"io"
	"io/ioutil"
	"os"
)

func LoadConfig(name string) (config image.Config, err error) {
	f, err := os.Open(name)
	if err != nil {
		return image.Config{}, err
	}
	defer f.Close()

	header := make([]byte, maxWebpHeaderSize)
	n, err := f.Read(header)
	if err != nil && err != io.EOF {
		return
	}
	header = header[:n]
	width, height, _, err := GetInfo(header)
	if err != nil {
		return
	}

	config.Width = width
	config.Height = height
	config.ColorModel = color.RGBAModel
	return
}

func Load(name string) (m image.Image, err error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		return nil, err
	}
	if fi.Size() > (2 << 30) {
		return nil, errors.New("webp: Load, file size is too large (> 2GB)!")
	}

	data := make([]byte, int(fi.Size()))
	if _, err = f.Read(data); err != nil {
		return nil, err
	}
	if m, err = DecodeRGBA(data); err != nil {
		return
	}
	return
}

// DecodeConfig returns the color model and dimensions of a WEBP image without
// decoding the entire image.
func DecodeConfig(r io.Reader) (config image.Config, err error) {
	header := make([]byte, maxWebpHeaderSize)
	n, err := r.Read(header)
	if err != nil && err != io.EOF {
		return
	}
	header, err = header[:n], nil
	width, height, _, err := GetInfo(header)
	if err != nil {
		return
	}
	config.Width = width
	config.Height = height
	config.ColorModel = color.RGBAModel
	return
}

// Decode reads a WEBP image from r and returns it as an image.Image.
func Decode(r io.Reader) (m image.Image, err error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return
	}
	if m, err = DecodeRGBA(data); err != nil {
		return
	}
	return
}

func init() {
	image.RegisterFormat("webp", "RIFF????WEBPVP8", Decode, DecodeConfig)
}
