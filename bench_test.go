// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package webp

import (
	"bytes"
	"testing"
)

func BenchmarkEncodeAndDecode_Safe(b *testing.B) {
	old := cgoIsUnsafePtr
	cgoIsUnsafePtr = false
	benchmarkEncodeAndDecode(b)
	cgoIsUnsafePtr = old
}

func BenchmarkEncodeAndDecode_Unsafe(b *testing.B) {
	old := cgoIsUnsafePtr
	cgoIsUnsafePtr = true
	benchmarkEncodeAndDecode(b)
	cgoIsUnsafePtr = old
}

func benchmarkEncodeAndDecode(b *testing.B) {
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
