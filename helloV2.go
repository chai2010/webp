// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore
// +build ignore

package main

import (
	"image/png"
	"log"
	"os"

	"github.com/cqzcqq/webp/convertor"
	"github.com/cqzcqq/webp/encoder"
)

func main() {
	file, err := os.Open("testdata/72.png")
	if err != nil {
		log.Fatalln(err)
	}

	img, err := png.Decode(file)
	if err != nil {
		log.Fatalln(err)
	}

	output, err := os.Create("testdata/72.webp")
	if err != nil {
		log.Fatal(err)
	}
	//noinspection GoUnhandledErrorResult
	defer output.Close()

	options, err := encoder.NewLossyEncoderOptions(encoder.PresetDefault, 75)
	if err != nil {
		log.Fatalln(err)
	}

	if err := convertor.Encode(output, img, options); err != nil {
		log.Fatalln(err)
	}
}
