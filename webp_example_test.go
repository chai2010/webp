// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package webp

import (
	"fmt"
	"io/ioutil"
	"log"
)

func xLoadData(filename string) []byte {
	data, err := ioutil.ReadFile("./testdata/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func xLoadCBuffer(filename string) CBuffer {
	data, err := ioutil.ReadFile("./testdata/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	cbuf := NewCBuffer(len(data))
	copy(cbuf.CData(), data)
	return cbuf
}

func ExampleGetInfo() {
	data := xLoadData("1_webp_a.webp")

	width, height, hasAlpha, err := GetInfo(data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("width: %v\n", width)
	fmt.Printf("height: %v\n", height)
	fmt.Printf("hasAlpha: %v\n", hasAlpha)

	// Output:
	// width: 400
	// height: 301
	// hasAlpha: true
}

func ExampleGetInfo_noAlpha() {
	data := xLoadData("video-001.webp")

	width, height, hasAlpha, err := GetInfo(data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("width: %v\n", width)
	fmt.Printf("height: %v\n", height)
	fmt.Printf("hasAlpha: %v\n", hasAlpha)

	// Output:
	// width: 150
	// height: 103
	// hasAlpha: false
}
