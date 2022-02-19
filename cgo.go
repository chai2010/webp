// Copyright 2016 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package webp

//#include "webp.h"
import "C"
import "unsafe"

type (
	C_int    C.int
	C_uint   C.uint
	C_float  C.float
	C_double C.double
	C_size_t C.size_t

	C_uint8_t  C.uint8_t
	C_uint16_t C.uint16_t
	C_uint32_t C.uint32_t
	C_uint64_t C.uint64_t

	C_int8_t  C.int8_t
	C_int16_t C.int16_t
	C_int32_t C.int32_t
	C_int64_t C.int64_t
)

func C_webpGetInfo(
	data *C_uint8_t, data_size C_size_t,
	width *C_int, height *C_int,
	has_alpha *C_int,
) C_int {
	return C_int(C.webpGetInfo(
		(*C.uint8_t)(data), (C.size_t)(data_size),
		(*C.int)(width), (*C.int)(height),
		(*C.int)(has_alpha),
	))
}

func C_webpDecodeGray(
	data *C_uint8_t, data_size C_size_t,
	width *C_int, height *C_int,
) *C_uint8_t {
	return (*C_uint8_t)(C.webpDecodeGray(
		(*C.uint8_t)(data), (C.size_t)(data_size),
		(*C.int)(width), (*C.int)(height),
	))
}

func C_webpDecodeRGB(
	data *C_uint8_t, data_size C_size_t,
	width *C_int, height *C_int,
) *C_uint8_t {
	return (*C_uint8_t)(C.webpDecodeRGB(
		(*C.uint8_t)(data), (C.size_t)(data_size),
		(*C.int)(width), (*C.int)(height),
	))
}

func C_webpDecodeRGBA(
	data *C_uint8_t, data_size C_size_t,
	width *C_int, height *C_int,
) *C_uint8_t {
	return (*C_uint8_t)(C.webpDecodeRGBA(
		(*C.uint8_t)(data), (C.size_t)(data_size),
		(*C.int)(width), (*C.int)(height),
	))
}

func C_webpDecodeGrayToSize(
	data *C_uint8_t, data_size C_size_t,
	width C_int, height C_int, outStride C_int,
	out *C_uint8_t,
) C_int {
	return (C_int)(C.webpDecodeGrayToSize(
		(*C.uint8_t)(data), (C.size_t)(data_size),
		(C.int)(width), (C.int)(height),
		(C.int)(outStride),
		(*C.uint8_t)(out),
	))
}

func C_webpDecodeRGBToSize(
	data *C_uint8_t, data_size C_size_t,
	width C_int, height C_int, outStride C_int,
	out *C_uint8_t,
) C_int {
	return (C_int)(C.webpDecodeRGBToSize(
		(*C.uint8_t)(data), (C.size_t)(data_size),
		(C.int)(width), (C.int)(height),
		(C.int)(outStride),
		(*C.uint8_t)(out),
	))
}

func C_webpDecodeRGBAToSize(
	data *C_uint8_t, data_size C_size_t,
	width C_int, height C_int, outStride C_int,
	out *C_uint8_t,
) C_int {
	return (C_int)(C.webpDecodeRGBAToSize(
		(*C.uint8_t)(data), (C.size_t)(data_size),
		(C.int)(width), (C.int)(height),
		(C.int)(outStride),
		(*C.uint8_t)(out),
	))
}

func C_webpEncodeGray(
	pix *C_uint8_t,
	width C_int, height C_int, stride C_int,
	quality_factor C_float,
	output_size *C_size_t,
) *C_uint8_t {
	return (*C_uint8_t)(C.webpEncodeGray(
		(*C.uint8_t)(pix),
		(C.int)(width), (C.int)(height), (C.int)(stride),
		(C.float)(quality_factor),
		(*C.size_t)(output_size),
	))
}

func C_webpEncodeRGB(
	pix *C_uint8_t,
	width C_int, height C_int, stride C_int,
	quality_factor C_float,
	output_size *C_size_t,
) *C_uint8_t {
	return (*C_uint8_t)(C.webpEncodeRGB(
		(*C.uint8_t)(pix),
		(C.int)(width), (C.int)(height), (C.int)(stride),
		(C.float)(quality_factor),
		(*C.size_t)(output_size),
	))
}

func C_webpEncodeRGBA(
	pix *C_uint8_t,
	width C_int, height C_int, stride C_int,
	quality_factor C_float,
	output_size *C_size_t,
) *C_uint8_t {
	return (*C_uint8_t)(C.webpEncodeRGBA(
		(*C.uint8_t)(pix),
		(C.int)(width), (C.int)(height), (C.int)(stride),
		(C.float)(quality_factor),
		(*C.size_t)(output_size),
	))
}

func C_webpEncodeLosslessGray(
	pix *C_uint8_t,
	width C_int, height C_int, stride C_int,
	output_size *C_size_t,
) *C_uint8_t {
	return (*C_uint8_t)(C.webpEncodeLosslessGray(
		(*C.uint8_t)(pix),
		(C.int)(width), (C.int)(height), (C.int)(stride),
		(*C.size_t)(output_size),
	))
}

func C_webpEncodeLosslessRGB(
	pix *C_uint8_t,
	width C_int, height C_int, stride C_int,
	output_size *C_size_t,
) *C_uint8_t {
	return (*C_uint8_t)(C.webpEncodeLosslessRGB(
		(*C.uint8_t)(pix),
		(C.int)(width), (C.int)(height), (C.int)(stride),
		(*C.size_t)(output_size),
	))
}

func C_webpEncodeLosslessRGBA(
	exact C_int,
	pix *C_uint8_t,
	width C_int, height C_int, stride C_int,
	output_size *C_size_t,
) *C_uint8_t {
	return (*C_uint8_t)(C.webpEncodeLosslessRGBA(
		(C.int)(exact),
		(*C.uint8_t)(pix),
		(C.int)(width), (C.int)(height), (C.int)(stride),
		(*C.size_t)(output_size),
	))
}

func C_webpMalloc(size C_size_t) unsafe.Pointer {
	return C.webpMalloc(C.size_t(size))
}

func C_webpFree(p unsafe.Pointer) {
	C.webpFree(p)
}
