// Copyright 2016 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package webp

import (
	"testing"
)

func TestWEBP_DECODER_ABI_VERSION(t *testing.T) {
	tAssertEQ(t, _C_WEBP_DECODER_ABI_VERSION, WEBP_DECODER_ABI_VERSION)
}
