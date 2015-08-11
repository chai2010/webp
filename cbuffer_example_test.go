// Copyright 2015 ChaiShushan <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package webp

import (
	"fmt"
	"log"
)

func ExampleCBuffer() {
	cbuf := NewCBuffer(100)
	defer cbuf.Close()

	data := cbuf.CData()
	fmt.Printf("CBufMagic: %v\n", cbuf.CBufMagic())
	fmt.Printf("cbuf own data[:]: %v\n", cbuf.Own(data[:]))
	fmt.Printf("cbuf own data[10:]: %v\n", cbuf.Own(data[10:]))
	fmt.Printf("cbuf own []byte{1}: %v\n", cbuf.Own([]byte{1}))

	// Output:
	// CBufMagic: CBufMagic
	// cbuf own data[:]: true
	// cbuf own data[10:]: true
	// cbuf own []byte{1}: false
}

func ExampleCBuffer_resize() {
	cbuf := NewCBuffer(100)
	defer cbuf.Close()

	data := cbuf.CData()
	fmt.Printf("len(data): %d\n", len(data))

	// resize, the old data is invalid!!!
	if err := cbuf.Resize(len(data) * 2); err != nil {
		log.Fatal(err)
	}

	// now size is 200
	data = cbuf.CData()
	fmt.Printf("len(data): %d\n", len(data))

	// Output:
	// len(data): 100
	// len(data): 200
}

func ExampleCBuffer_lockAddress() {
	const dontResize = true
	cbuf := NewCBuffer(100, dontResize)
	defer cbuf.Close()

	fmt.Printf("CanResize: %v\n", cbuf.CanResize())

	// can't resize now
	if err := cbuf.Resize(len(cbuf.CData()) * 2); err == nil {
		log.Fatal("expect not nil")
	}

	// size is still 100
	data := cbuf.CData()
	fmt.Printf("len(data): %d\n", len(data))

	// Output:
	// CanResize: false
	// len(data): 100
}
