// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package webp

import (
	"io/ioutil"
	"testing"
)

type tGetInfoTester struct {
	Filename string
	HdrSize  int
	Width    int
	Height   int
	HasAlpha bool
}

func TestGetInfo(t *testing.T) {
	for i, v := range tGetInfoTesterList {
		data, err := ioutil.ReadFile(testdataDir + v.Filename)
		if err != nil {
			t.Fatalf("%d: %v", i, err)
		}
		width, height, hasAlpha, err := GetInfo(data)
		if err != nil {
			t.Fatalf("%d: %v", i, err)
		}
		if width != v.Width {
			t.Fatalf("%d: expect = %v, got = %v", i, v.Width, width)
		}
		if height != v.Height {
			t.Fatalf("%d: expect = %v, got = %v", i, v.Height, height)
		}
		if hasAlpha != v.HasAlpha {
			t.Fatalf("%d: expect = %v, got = %v", i, v.HasAlpha, hasAlpha)
		}
	}
}

var tGetInfoTesterList = []tGetInfoTester{
	tGetInfoTester{
		Filename: "1_webp_ll.webp",
		Width:    400,
		Height:   301,
		HasAlpha: true,
	},
}
