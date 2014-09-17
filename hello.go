// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ingore

package main

import (
	"fmt"
	"image"
	"io/ioutil"
	"log"

	"github.com/chai2010/webp"
)

func main() {
	var cfg image.Config
	var data []byte
	var err error

	// Load file data
	if data, err = ioutil.ReadFile("./testdata/1_webp_ll.webp"); err != nil {
		log.Fatal(err)
	}

	// GetInfo
	if cfg.Width, cfg.Height, _, err = webp.GetInfo(data); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("width = %d, height = %d\n", cfg.Width, cfg.Height)

	// Decode webp
	rgba, err := webp.DecodeRGBA(data)
	if err != nil {
		log.Fatal(err)
	}

	// Encode lossless webp
	if data, err = webp.EncodeLosslessRGBA(rgba); err != nil {
		log.Fatalf("saveWebp: webp.EncodeLosslessRGBA, err = %v", err)
	}
	if err = ioutil.WriteFile("output.webp", data, 0666); err != nil {
		log.Fatalf("saveWebp: ioutil.WriteFile, err = %v", err)
	}
}
