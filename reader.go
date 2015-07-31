// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build cgo

package webp

import (
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
	return DecodeConfig(f)
}

func Load(name string) (m image.Image, err error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return Decode(f)
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
