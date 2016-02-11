// Copyright 2016 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package webp

//#include <webp/decode.h>
import "C"

const (
	WEBP_DECODER_ABI_VERSION = 0x0203 // MAJOR(8b) + MINOR(8b)
)

const (
	_C_WEBP_DECODER_ABI_VERSION = C.WEBP_DECODER_ABI_VERSION // for test
)

// Return the decoder's version number, packed in hexadecimal using 8bits for
// each of major/minor/revision. E.g: v2.5.7 is 0x020507.
func WebPGetDecoderVersion() uint {
	return uint(C.WebPGetDecoderVersion())
}
