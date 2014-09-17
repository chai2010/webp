// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ingore

package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/chai2010/webp"
)

func main() {
	var width, height int
	var data []byte
	var err error

	// Load file data
	if data, err = ioutil.ReadFile("./testdata/1_webp_ll.webp"); err != nil {
		log.Fatal(err)
	}

	// GetInfo
	if width, height, _, err = webp.GetInfo(data); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("image size: width = %d, height = %d\n", width, height)

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
	fmt.Printf("Save output.webp ok\n")
}
