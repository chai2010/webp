// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package webp

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func BenchmarkGetInfo(b *testing.B) {
	data, err := ioutil.ReadFile("./testdata/1_webp_ll.webp")
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GetInfo(data)
	}
	b.StopTimer()
}

func BenchmarkDecodeGray(b *testing.B) {
	data, err := ioutil.ReadFile("./testdata/1_webp_ll.webp")
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := DecodeGray(data)
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
	b.StopTimer()
}

func BenchmarkDecodeRGB(b *testing.B) {
	data, err := ioutil.ReadFile("./testdata/1_webp_ll.webp")
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := DecodeRGB(data)
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
	b.StopTimer()
}

func BenchmarkDecodeRGBA(b *testing.B) {
	data, err := ioutil.ReadFile("./testdata/1_webp_ll.webp")
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := DecodeRGBA(data)
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
	b.StopTimer()
}

func BenchmarkDecodeGrayToSize(b *testing.B) {
	data, err := ioutil.ReadFile("./testdata/1_webp_ll.webp")
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := DecodeGrayToSize(data, 32, 32)
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
	b.StopTimer()
}

func BenchmarkDecodeRGBToSize(b *testing.B) {
	data, err := ioutil.ReadFile("./testdata/1_webp_ll.webp")
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := DecodeRGBToSize(data, 32, 32)
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
	b.StopTimer()
}

func BenchmarkDecodeRGBAToSize(b *testing.B) {
	data, err := ioutil.ReadFile("./testdata/1_webp_ll.webp")
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := DecodeRGBAToSize(data, 32, 32)
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
	b.StopTimer()
}

func BenchmarkEncodeAndDecode(b *testing.B) {
	var buf bytes.Buffer

	img, err := loadImage("1_webp_ll.png")
	if err != nil {
		b.Fatal(err)
	}
	s := img.Bounds().Size()
	b.SetBytes(int64(s.X * s.Y * 4))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err = Encode(&buf, img, nil); err != nil {
			b.Fatal(err)
		}
		if _, err = Decode(&buf); err != nil {
			b.Fatal(err)
		}
	}
}
