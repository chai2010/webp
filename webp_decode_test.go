// Copyright 2016 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package webp

import (
	"io/ioutil"
	"testing"
)

func TestWEBP_DECODER_ABI_VERSION(t *testing.T) {
	tAssertEQ(t, _C_WEBP_DECODER_ABI_VERSION, WEBP_DECODER_ABI_VERSION)
}

func TestWebPGetInfo(t *testing.T) {
	data, err := ioutil.ReadFile("./testdata/1_webp_ll.webp")
	tAssertNil(t, err)

	w, h, ok := WebPGetInfo(data)
	tAssertEQ(t, 400, w)
	tAssertEQ(t, 301, h)
	tAssert(t, ok)
}
