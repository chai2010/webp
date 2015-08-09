// Copyright 2015 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package webp_bench

import (
	"bytes"
	"io/ioutil"
	"testing"

	chai2010_webp "github.com/chai2010/webp"
	x_image_webp "golang.org/x/image/webp"
)

type CBuffer interface {
	chai2010_webp.CBuffer
}

func tbLoadData(tb testing.TB, filename string) []byte {
	data, err := ioutil.ReadFile("../testdata/" + filename)
	if err != nil {
		tb.Fatal(err)
	}
	return data
}

func tbLoadCData(tb testing.TB, filename string) CBuffer {
	data, err := ioutil.ReadFile("../testdata/" + filename)
	if err != nil {
		tb.Fatal(err)
	}
	cbuf := chai2010_webp.NewCBuffer(len(data))
	copy(cbuf.CData(), data)
	return cbuf
}

func BenchmarkDecode_1_a_chai2010_webp(b *testing.B) {
	data := tbLoadData(b, "1_webp_a.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := chai2010_webp.Decode(bytes.NewReader(data))
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}
func BenchmarkDecode_1_a_chai2010_webp_cbuf(b *testing.B) {
	cbuf := tbLoadCData(b, "1_webp_a.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, pix, err := chai2010_webp.DecodeRGBAEx(cbuf.CData(), cbuf)
		if err != nil {
			b.Fatal(err)
		}
		_ = m
		pix.Close()
	}
}
func BenchmarkDecode_1_a_x_image_webp(b *testing.B) {
	data := tbLoadData(b, "1_webp_a.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := x_image_webp.Decode(bytes.NewReader(data))
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}

func BenchmarkDecode_1_ll(b *testing.B) {
	//
}
func BenchmarkDecode_1_ll_x_image_webp(b *testing.B) {
	//
}

func BenchmarkDecode_2_a(b *testing.B) {
	//
}
func BenchmarkDecode_2_a_x_image_webp(b *testing.B) {
	//
}

func BenchmarkDecode_2_ll(b *testing.B) {
	//
}
func BenchmarkDecode_2_ll_x_image_webp(b *testing.B) {
	//
}

func BenchmarkDecode_3_a(b *testing.B) {
	//
}
func BenchmarkDecode_3_a_x_image_webp(b *testing.B) {
	//
}

func BenchmarkDecode_3_ll(b *testing.B) {
	//
}
func BenchmarkDecode_3_ll_x_image_webp(b *testing.B) {
	//
}

func BenchmarkDecode_4_a(b *testing.B) {
	//
}
func BenchmarkDecode_4_a_x_image_webp(b *testing.B) {
	//
}

func BenchmarkDecode_4_ll(b *testing.B) {
	//
}
func BenchmarkDecode_4_ll_x_image_webp(b *testing.B) {
	//
}

func BenchmarkDecode_5_a(b *testing.B) {
	//
}
func BenchmarkDecode_5_a_x_image_webp(b *testing.B) {
	//
}

func BenchmarkDecode_5_ll(b *testing.B) {
	//
}
func BenchmarkDecode_5_ll_x_image_webp(b *testing.B) {
	//
}
