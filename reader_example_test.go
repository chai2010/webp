// Copyright 2015 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package webp

import (
	"fmt"
	"log"
)

func ExampleLoadConfig() {
	cfg, err := LoadConfig("./testdata/1_webp_ll.webp")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Width = %d\n", cfg.Width)
	fmt.Printf("Height = %d\n", cfg.Height)
	// Output:
	// Width = 400
	// Height = 301
}

func ExampleLoadConfig_cbuf() {
	cbuf := NewCBuffer(0)
	defer cbuf.Close()

	cfg, err := LoadConfig("./testdata/1_webp_ll.webp", cbuf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Width = %d\n", cfg.Width)
	fmt.Printf("Height = %d\n", cfg.Height)
	// Output:
	// Width = 400
	// Height = 301
}

func ExampleLoad() {
	m, err := Load("./testdata/1_webp_ll.webp")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Bounds = %v\n", m.Bounds())
	// Output:
	// Bounds = (0,0)-(400,301)
}

func ExampleLoad_cbuf() {
	cbuf := NewCBuffer(8 << 20) // 8MB
	defer cbuf.Close()

	m, err := Load("./testdata/1_webp_ll.webp", cbuf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Bounds = %v\n", m.Bounds())
	// Output:
	// Bounds = (0,0)-(400,301)
}
