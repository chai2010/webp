// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package webp

/*
#cgo CFLAGS: -I./libwebp/include  -I./libwebp/src
#cgo !windows LDFLAGS: -lm

#include "webp.h"
*/
import "C"
import (
	"errors"
	"unsafe"
)

func webpGetInfo(data []byte) (width, height int, has_alpha bool, err error) {
	if len(data) == 0 {
		err = errors.New("webpGetInfo: bad arguments")
		return
	}
	var c_width, c_height, c_has_alpha C.int
	rv := C.webpGetInfo(
		(*C.uint8_t)(unsafe.Pointer(&data[0])), C.size_t(len(data)),
		&c_width, &c_height, &c_has_alpha,
	)
	if rv == 0 {
		err = errors.New("webpGetInfo: failed")
		return
	}
	width, height = int(c_width), int(c_height)
	has_alpha = (c_has_alpha != 0)
	return
}

func webpDecodeGray(data []byte) (pix []byte, width, height int, err error) {
	if len(data) == 0 {
		err = errors.New("webpDecodeGray: bad arguments")
		return
	}
	var c_width, c_height C.int
	d := C.webpDecodeGray(
		(*C.uint8_t)(unsafe.Pointer(&data[0])), C.size_t(len(data)),
		&c_width, &c_height,
	)
	if d == nil {
		err = errors.New("webpDecodeGray: failed")
		return
	}
	width, height = int(c_width), int(c_height)
	pix = make([]byte, width*height*1)
	copy(pix, ((*[1 << 30]byte)(unsafe.Pointer(d)))[0:len(pix):len(pix)])
	C.webpFree(unsafe.Pointer(d))
	return
}

func webpDecodeRGB(data []byte) (pix []byte, width, height int, err error) {
	if len(data) == 0 {
		err = errors.New("webpDecodeRGB: bad arguments")
		return
	}
	var c_width, c_height C.int
	d := C.webpDecodeRGB(
		(*C.uint8_t)(unsafe.Pointer(&data[0])), C.size_t(len(data)),
		&c_width, &c_height,
	)
	if d == nil {
		err = errors.New("webpDecodeRGB: failed")
		return
	}
	width, height = int(c_width), int(c_height)
	pix = make([]byte, width*height*3)
	copy(pix, ((*[1 << 30]byte)(unsafe.Pointer(d)))[0:len(pix):len(pix)])
	C.webpFree(unsafe.Pointer(d))
	return
}

func webpDecodeRGBA(data []byte) (pix []byte, width, height int, err error) {
	if len(data) == 0 {
		err = errors.New("webpDecodeRGBA: bad arguments")
		return
	}
	var c_width, c_height C.int
	d := C.webpDecodeRGBA(
		(*C.uint8_t)(unsafe.Pointer(&data[0])), C.size_t(len(data)),
		&c_width, &c_height,
	)
	if d == nil {
		err = errors.New("webpDecodeRGBA: failed")
		return
	}
	width, height = int(c_width), int(c_height)
	pix = make([]byte, width*height*4)
	copy(pix, ((*[1 << 30]byte)(unsafe.Pointer(d)))[0:len(pix):len(pix)])
	C.webpFree(unsafe.Pointer(d))
	return
}

func webpEncodeGray(
	pix []byte, width, height, stride int,
	quality_factor float32,
) (output []byte, err error) {
	if len(pix) == 0 || width <= 0 || height <= 0 || stride <= 0 || quality_factor < 0.0 {
		err = errors.New("webpEncodeGray: bad arguments")
		return
	}
	if stride < width*1 && len(pix) < height*stride {
		err = errors.New("webpEncodeGray: bad arguments")
		return
	}
	var d *C.uint8_t
	d_size := C.webpEncodeGray(
		(*C.uint8_t)(unsafe.Pointer(&pix[0])), C.int(width), C.int(height),
		C.int(stride), C.float(quality_factor),
		&d,
	)
	if d_size == 0 {
		err = errors.New("webpEncodeGray: failed")
		return
	}
	output = make([]byte, int(d_size))
	copy(output, ((*[1 << 30]byte)(unsafe.Pointer(d)))[0:len(output):len(output)])
	C.webpFree(unsafe.Pointer(d))
	return
}

func webpEncodeRGB(
	pix []byte, width, height, stride int,
	quality_factor float32,
) (output []byte, err error) {
	if len(pix) == 0 || width <= 0 || height <= 0 || stride <= 0 || quality_factor < 0.0 {
		err = errors.New("webpEncodeRGB: bad arguments")
		return
	}
	if stride < width*3 && len(pix) < height*stride {
		err = errors.New("webpEncodeRGB: bad arguments")
		return
	}
	var d *C.uint8_t
	d_size := C.webpEncodeRGB(
		(*C.uint8_t)(unsafe.Pointer(&pix[0])), C.int(width), C.int(height),
		C.int(stride), C.float(quality_factor),
		&d,
	)
	if d_size == 0 {
		err = errors.New("webpEncodeRGB: failed")
		return
	}
	output = make([]byte, int(d_size))
	copy(output, ((*[1 << 30]byte)(unsafe.Pointer(d)))[0:len(output):len(output)])
	C.webpFree(unsafe.Pointer(d))
	return
}

func webpEncodeRGBA(
	pix []byte, width, height, stride int,
	quality_factor float32,
) (output []byte, err error) {
	if len(pix) == 0 || width <= 0 || height <= 0 || stride <= 0 || quality_factor < 0.0 {
		err = errors.New("webpEncodeRGBA: bad arguments")
		return
	}
	if stride < width*4 && len(pix) < height*stride {
		err = errors.New("webpEncodeRGBA: bad arguments")
		return
	}
	var d *C.uint8_t
	d_size := C.webpEncodeRGBA(
		(*C.uint8_t)(unsafe.Pointer(&pix[0])), C.int(width), C.int(height),
		C.int(stride), C.float(quality_factor),
		&d,
	)
	if d_size == 0 {
		err = errors.New("webpEncodeRGBA: failed")
		return
	}
	output = make([]byte, int(d_size))
	copy(output, ((*[1 << 30]byte)(unsafe.Pointer(d)))[0:len(output):len(output)])
	C.webpFree(unsafe.Pointer(d))
	return
}

func webpEncodeLosslessGray(
	pix []byte, width, height, stride int,
) (output []byte, err error) {
	if len(pix) == 0 || width <= 0 || height <= 0 || stride <= 0 {
		err = errors.New("webpEncodeLosslessGray: bad arguments")
		return
	}
	if stride < width*1 && len(pix) < height*stride {
		err = errors.New("webpEncodeLosslessGray: bad arguments")
		return
	}
	var d *C.uint8_t
	d_size := C.webpEncodeLosslessGray(
		(*C.uint8_t)(unsafe.Pointer(&pix[0])), C.int(width), C.int(height),
		C.int(stride),
		&d,
	)
	if d_size == 0 {
		err = errors.New("webpEncodeLosslessGray: failed")
		return
	}
	output = make([]byte, int(d_size))
	copy(output, ((*[1 << 30]byte)(unsafe.Pointer(d)))[0:len(output):len(output)])
	C.webpFree(unsafe.Pointer(d))
	return
}

func webpEncodeLosslessRGB(
	pix []byte, width, height, stride int,
) (output []byte, err error) {
	if len(pix) == 0 || width <= 0 || height <= 0 || stride <= 0 {
		err = errors.New("webpEncodeLosslessRGB: bad arguments")
		return
	}
	if stride < width*3 && len(pix) < height*stride {
		err = errors.New("webpEncodeLosslessRGB: bad arguments")
		return
	}
	var d *C.uint8_t
	d_size := C.webpEncodeLosslessRGB(
		(*C.uint8_t)(unsafe.Pointer(&pix[0])), C.int(width), C.int(height),
		C.int(stride),
		&d,
	)
	if d_size == 0 {
		err = errors.New("webpEncodeLosslessRGB: failed")
		return
	}
	output = make([]byte, int(d_size))
	copy(output, ((*[1 << 30]byte)(unsafe.Pointer(d)))[0:len(output):len(output)])
	C.webpFree(unsafe.Pointer(d))
	return
}

func webpEncodeLosslessRGBA(
	pix []byte, width, height, stride int,
) (output []byte, err error) {
	if len(pix) == 0 || width <= 0 || height <= 0 || stride <= 0 {
		err = errors.New("webpEncodeLosslessRGBA: bad arguments")
		return
	}
	if stride < width*4 && len(pix) < height*stride {
		err = errors.New("webpEncodeLosslessRGBA: bad arguments")
		return
	}
	var d *C.uint8_t
	d_size := C.webpEncodeLosslessRGBA(
		(*C.uint8_t)(unsafe.Pointer(&pix[0])), C.int(width), C.int(height),
		C.int(stride),
		&d,
	)
	if d_size == 0 {
		err = errors.New("webpEncodeLosslessRGBA: failed")
		return
	}
	output = make([]byte, int(d_size))
	copy(output, ((*[1 << 30]byte)(unsafe.Pointer(d)))[0:len(output):len(output)])
	C.webpFree(unsafe.Pointer(d))
	return
}
