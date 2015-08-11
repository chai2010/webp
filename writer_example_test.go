// Copyright 2015 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package webp

import (
	"bytes"
	"image"
	"log"
	"os"
	"reflect"
)

func ExampleSave() {
	tmpname := "z_test_ExampleSave.webp"
	defer os.Remove(tmpname)

	gray := NewMemPImage(image.Rect(0, 0, 400, 300), 1, reflect.Uint8)
	if err := Save(tmpname, gray, &Options{Quality: 75}); err != nil {
		log.Fatal(err)
	}
}

func ExampleSave_cbuf() {
	tmpname := "z_test_ExampleSave.webp"
	defer os.Remove(tmpname)

	b := image.Rect(0, 0, 400, 300)
	cbuf := NewCBuffer(b.Dx() * b.Dy() * 4)
	defer cbuf.Close()

	rgba := &image.RGBA{
		Pix:    cbuf.CData(),
		Stride: b.Dx() * 4,
		Rect:   b,
	}
	if err := Save(tmpname, rgba, nil, cbuf); err != nil {
		log.Fatal(err)
	}

	gray := &image.Gray{
		Pix:    cbuf.CData(),
		Stride: b.Dx() * 1,
		Rect:   b,
	}
	if err := Save(tmpname, gray, nil, cbuf); err != nil {
		log.Fatal(err)
	}
}

func ExampleEncode() {
	m, err := Load("./testdata/1_webp_ll.webp")
	if err != nil {
		log.Fatal(err)
	}

	var buf bytes.Buffer
	if err := Encode(&buf, m, nil); err != nil {
		log.Fatal(err)
	}
	_ = buf.Bytes()
}

func ExampleEncode_lossless() {
	m, err := Load("./testdata/1_webp_ll.webp")
	if err != nil {
		log.Fatal(err)
	}

	var buf bytes.Buffer
	if err := Encode(&buf, m, &Options{Lossless: true}); err != nil {
		log.Fatal(err)
	}
	_ = buf.Bytes()
}

func ExampleEncode_rgb() {
	rgb := NewRGBImage(image.Rect(0, 0, 400, 300))

	var buf bytes.Buffer
	if err := Encode(&buf, rgb, nil); err != nil {
		log.Fatal(err)
	}
	_ = buf.Bytes()
}

func ExampleEncode_rgb48MemP() {
	rgb48 := NewMemPImage(image.Rect(0, 0, 400, 300), 3, reflect.Uint16)

	var buf bytes.Buffer
	if err := Encode(&buf, rgb48, nil); err != nil {
		log.Fatal(err)
	}
	_ = buf.Bytes()
}
