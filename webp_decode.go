// Copyright 2016 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package webp

//#include <webp/decode.h>
import "C"
import "unsafe"

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

// Retrieve basic header information: width, height.
// This function will also validate the header and return 0 in
// case of formatting error.
// Pointers 'width' and 'height' can be passed NULL if deemed irrelevant.
func WebPGetInfo(data []byte) (width, height int, ok bool) {
	if len(data) == 0 {
		return 0, 0, false
	}
	var cw, ch C.int
	if C.WebPGetInfo((*C.uint8_t)(unsafe.Pointer(&data[0])), C.size_t(len(data)), &cw, &ch) == 0 {
		return 0, 0, false
	}
	return int(cw), int(ch), true
}
