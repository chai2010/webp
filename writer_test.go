// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package webp

import (
	"bytes"
	_ "image/png"
	"testing"
)

type tTester struct {
	Filename string
	Lossless bool
	Quality  float32 // 0 ~ 100
	MaxDelta int
	MinDelta int
	Exact    bool
}

var tTesterList = []tTester{
	tTester{
		Filename: "video-001.png",
		Lossless: false,
		Quality:  90,
		MaxDelta: 5,
	},
	tTester{
		Filename: "1_webp_ll.png",
		Lossless: false,
		Quality:  90,
		MaxDelta: 5,
	},
	tTester{
		Filename: "2_webp_ll.png",
		Lossless: true,
		Exact:    true,
		Quality:  90,
		MaxDelta: 0,
	},
	tTester{
		Filename: "3_webp_ll.png",
		Lossless: false,
		Quality:  90,
		MaxDelta: 5,
	},
	tTester{
		Filename: "4_webp_ll.png",
		Lossless: true,
		Exact:    true,
		Quality:  90,
		MaxDelta: 0,
	},
	tTester{
		Filename: "5_webp_ll.png",
		Lossless: false,
		Quality:  75,
		MaxDelta: 15,
	},
	tTester{
		Filename: "4_webp_ll.png",
		Lossless: true,
		Quality:  90,
		Exact:    false,
		MaxDelta: 13,
		MinDelta: 10,
	},
}

func TestEncode(t *testing.T) {
	for i, v := range tTesterList {
		img0, err := loadImage(v.Filename)
		if err != nil {
			t.Fatalf("%d: %v", i, err)
		}

		buf := new(bytes.Buffer)
		err = Encode(buf, img0, &Options{
			Lossless: v.Lossless,
			Quality:  v.Quality,
			Exact:    v.Exact,
		})
		if err != nil {
			t.Fatalf("%d: %v", i, err)
		}

		img1, err := Decode(buf)
		if err != nil {
			t.Fatalf("%d: %v", i, err)
		}

		// Compare the average delta to the tolerance level.
		var want int
		if !v.Lossless || !v.Exact {
			want = v.MaxDelta
		}
		got := averageDelta(img0, img1)
		if got > want {
			t.Fatalf("%d: average delta too high; got %d, want <= %d", i, got, want)
		}
		if v.MinDelta > 0 && got < v.MinDelta {
			t.Fatalf("%d: average delta too low; got %d; want >= %d", i, got, v.MinDelta)
		}
	}
}
