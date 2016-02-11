// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//
// TODO(chai2010): simplify cgo pointer:
//
// CBuffer/cgoSafePtr/cgoFreePtr removed!
//
// Go1.3: Changes to the garbage collector
// http://golang.org/doc/go1.3#garbage_collector
//
// Go1.6:
// https://github.com/golang/proposal/blob/master/design/12416-cgo-pointers.md
//

package webp

/*
#cgo CFLAGS: -I./internal/libwebp/include  -I./internal/libwebp/src -Wno-pointer-sign -DWEBP_USE_THREAD
#cgo !windows LDFLAGS: -lm

#include <webp.h>
#include <webp/decode.h>

#include <stdlib.h>
#include <string.h>

struct cgoWebpDecodeRGBAReturn {
	int ok;
	int width;
	int height;
	uint8_t* ptr;
} cgoWebpDecodeRGBA(const uint8_t* data, size_t data_size) {
	struct cgoWebpDecodeRGBAReturn t;
	t.ptr = webpDecodeRGBA(data, data_size, &t.width, &t.height);
	t.ok = (t.ptr != NULL)? 1: 0;
	return t;
}

struct cgoWebpEncodeGrayReturn {
	int ok;
	size_t size;
	uint8_t* ptr;
} cgoWebpEncodeGray(const uint8_t* data, int width, int height, int stride, float quality_factor) {
	struct cgoWebpEncodeGrayReturn t;
	t.size = webpEncodeGray(data, width, height, stride, quality_factor, &t.ptr);
	t.ok = (t.size != 0)? 1: 0;
	return t;
}

struct cgoWebpEncodeRGBReturn {
	int ok;
	size_t size;
	uint8_t* ptr;
} cgoWebpEncodeRGB(const uint8_t* data, int width, int height, int stride, float quality_factor) {
	struct cgoWebpEncodeRGBReturn t;
	t.size = webpEncodeRGB(data, width, height, stride, quality_factor, &t.ptr);
	t.ok = (t.size != 0)? 1: 0;
	return t;
}

struct cgoWebpEncodeRGBAReturn {
	int ok;
	size_t size;
	uint8_t* ptr;
} cgoWebpEncodeRGBA(const uint8_t* data, int width, int height, int stride, float quality_factor) {
	struct cgoWebpEncodeRGBAReturn t;
	t.size = webpEncodeRGBA(data, width, height, stride, quality_factor, &t.ptr);
	t.ok = (t.size != 0)? 1: 0;
	return t;
}

struct cgoWebpEncodeLosslessGrayReturn {
	int ok;
	size_t size;
	uint8_t* ptr;
} cgoWebpEncodeLosslessGray(const uint8_t* data, int width, int height, int stride) {
	struct cgoWebpEncodeLosslessGrayReturn t;
	t.size = webpEncodeLosslessGray(data, width, height, stride, &t.ptr);
	t.ok = (t.size != 0)? 1: 0;
	return t;
}

struct cgoWebpEncodeLosslessRGBReturn {
	int ok;
	size_t size;
	uint8_t* ptr;
} cgoWebpEncodeLosslessRGB(const uint8_t* data, int width, int height, int stride) {
	struct cgoWebpEncodeLosslessRGBReturn t;
	t.size = webpEncodeLosslessRGB(data, width, height, stride, &t.ptr);
	t.ok = (t.size != 0)? 1: 0;
	return t;
}

struct cgoWebpEncodeLosslessRGBAReturn {
	int ok;
	size_t size;
	uint8_t* ptr;
} cgoWebpEncodeLosslessRGBA(const uint8_t* data, int width, int height, int stride) {
	struct cgoWebpEncodeLosslessRGBAReturn t;
	t.size = webpEncodeLosslessRGBA(data, width, height, stride, &t.ptr);
	t.ok = (t.size != 0)? 1: 0;
	return t;
}

struct cgoWebpGetEXIFReturn {
	int ok;
	size_t size;
	uint8_t* ptr;
} cgoWebpGetEXIF(const uint8_t* data, size_t data_size) {
	struct cgoWebpGetEXIFReturn t;
	t.ptr = webpGetEXIF(data, data_size, &t.size);
	t.ok = (t.size != 0)? 1: 0;
	return t;
}
struct cgoWebpGetICCPReturn {
	int ok;
	size_t size;
	uint8_t* ptr;
} cgoWebpGetICCP(const uint8_t* data, size_t data_size) {
	struct cgoWebpGetICCPReturn t;
	t.ptr = webpGetICCP(data, data_size, &t.size);
	t.ok = (t.size != 0)? 1: 0;
	return t;
}
struct cgoWebpGetXMPReturn {
	int ok;
	size_t size;
	uint8_t* ptr;
} cgoWebpGetXMP(const uint8_t* data, size_t data_size) {
	struct cgoWebpGetXMPReturn t;
	t.ptr = webpGetXMP(data, data_size, &t.size);
	t.ok = (t.size != 0)? 1: 0;
	return t;
}

struct cgoWebpSetEXIFReturn {
	int ok;
	size_t size;
	uint8_t* ptr;
} cgoWebpSetEXIF(const uint8_t* data, size_t data_size, const char* metadata, size_t metadata_size) {
	struct cgoWebpSetEXIFReturn t;
	t.ptr = webpSetEXIF(data, data_size, metadata, metadata_size, &t.size);
	t.ok = (t.size != 0)? 1: 0;
	return t;
}
struct cgoWebpSetICCPReturn {
	int ok;
	size_t size;
	uint8_t* ptr;
} cgoWebpSetICCP(const uint8_t* data, size_t data_size, const char* metadata, size_t metadata_size) {
	struct cgoWebpSetICCPReturn t;
	t.ptr = webpSetICCP(data, data_size, metadata, metadata_size, &t.size);
	t.ok = (t.size != 0)? 1: 0;
	return t;
}
struct cgoWebpSetXMPReturn {
	int ok;
	size_t size;
	uint8_t* ptr;
} cgoWebpSetXMP(const uint8_t* data, size_t data_size, const char* metadata, size_t metadata_size) {
	struct cgoWebpSetXMPReturn t;
	t.ptr = webpSetXMP(data, data_size, metadata, metadata_size, &t.size);
	t.ok = (t.size != 0)? 1: 0;
	return t;
}

struct cgoWebpDelEXIFReturn {
	int ok;
	size_t size;
	uint8_t* ptr;
} cgoWebpDelEXIF(const uint8_t* data, size_t data_size) {
	struct cgoWebpDelEXIFReturn t;
	t.ptr = webpDelEXIF(data, data_size, &t.size);
	t.ok = (t.size != 0)? 1: 0;
	return t;
}
struct cgoWebpDelICCPReturn {
	int ok;
	size_t size;
	uint8_t* ptr;
} cgoWebpDelICCP(const uint8_t* data, size_t data_size) {
	struct cgoWebpDelICCPReturn t;
	t.ptr = webpDelICCP(data, data_size, &t.size);
	t.ok = (t.size != 0)? 1: 0;
	return t;
}
struct cgoWebpDelXMPReturn {
	int ok;
	size_t size;
	uint8_t* ptr;
} cgoWebpDelXMP(const uint8_t* data, size_t data_size) {
	struct cgoWebpDelXMPReturn t;
	t.ptr = webpDelXMP(data, data_size, &t.size);
	t.ok = (t.size != 0)? 1: 0;
	return t;
}

*/
import "C"
import (
	"errors"
	"unsafe"
)

func webpGetInfo(data []byte) (width, height int, hasAlpha bool, err error) {
	if len(data) == 0 {
		err = errors.New("webpGetInfo: bad arguments, data is empty")
		return
	}
	if len(data) > maxWebpHeaderSize {
		data = data[:maxWebpHeaderSize]
	}

	var features C.WebPBitstreamFeatures
	if C.WebPGetFeatures((*C.uint8_t)(unsafe.Pointer(&data[0])), C.size_t(len(data)), &features) != C.VP8_STATUS_OK {
		err = errors.New("C.WebPGetFeatures: failed")
		return
	}
	width, height = int(features.width), int(features.height)
	hasAlpha = (features.has_alpha != 0)
	return
}

func webpDecodeGray(data []byte) (pix []byte, width, height int, err error) {
	if len(data) == 0 {
		err = errors.New("webpDecodeGray: bad arguments")
		return
	}

	var cw, ch C.int
	var cptr = C.webpDecodeGray((*C.uint8_t)(unsafe.Pointer(&data[0])), C.size_t(len(data)), &cw, &ch)
	if cptr == nil {
		err = errors.New("webpDecodeGray: failed")
		return
	}
	defer C.free(unsafe.Pointer(cptr))

	pix = append([]byte{}, ((*[1 << 30]byte)(unsafe.Pointer(cptr)))[0:len(pix):len(pix)]...)
	width, height = int(cw), int(ch)
	return
}

func webpDecodeRGB(data []byte) (pix []byte, width, height int, err error) {
	if len(data) == 0 {
		err = errors.New("webpDecodeRGB: bad arguments")
		return
	}

	var cw, ch C.int
	var cptr = C.webpDecodeRGB((*C.uint8_t)(unsafe.Pointer(&data[0])), C.size_t(len(data)), &cw, &ch)
	if cptr == nil {
		err = errors.New("webpDecodeRGB: failed")
		return
	}
	defer C.free(unsafe.Pointer(cptr))

	pix = append([]byte{}, ((*[1 << 30]byte)(unsafe.Pointer(cptr)))[0:len(pix):len(pix)]...)
	width, height = int(cw), int(ch)
	return
}

func webpDecodeRGBA(data []byte, cbuf CBuffer) (pix CBuffer, width, height int, err error) {
	if len(data) == 0 {
		err = errors.New("webpDecodeRGBA: bad arguments")
		return
	}
	isCBuf := cbuf.Own(data)
	cData := cgoSafePtr(data, isCBuf)
	defer cgoFreePtr(cData, isCBuf)

	rv := C.cgoWebpDecodeRGBA((*C.uint8_t)(cData), C.size_t(len(data)))
	if rv.ok != 1 {
		err = errors.New("webpDecodeRGBA: failed")
		return
	}

	width, height = int(rv.width), int(rv.height)
	pix = newCBufferFrom(unsafe.Pointer(rv.ptr), width*height*4)
	return
}

func webpEncodeGray(pix []byte, width, height, stride int, quality float32, cbuf CBuffer) (output CBuffer, err error) {
	if len(pix) == 0 || width <= 0 || height <= 0 || stride <= 0 || quality < 0.0 {
		err = errors.New("webpEncodeGray: bad arguments")
		return
	}
	if stride < width*1 && len(pix) < height*stride {
		err = errors.New("webpEncodeGray: bad arguments")
		return
	}
	isCBuf := cbuf.Own(pix)
	cPix := cgoSafePtr(pix, isCBuf)
	defer cgoFreePtr(cPix, isCBuf)

	rv := C.cgoWebpEncodeGray(
		(*C.uint8_t)(cPix), C.int(width), C.int(height),
		C.int(stride), C.float(quality),
	)
	if rv.ok != 1 {
		err = errors.New("webpEncodeGray: failed")
		return
	}

	output = newCBufferFrom(unsafe.Pointer(rv.ptr), int(rv.size))
	return
}

func webpEncodeRGB(pix []byte, width, height, stride int, quality float32, cbuf CBuffer) (output CBuffer, err error) {
	if len(pix) == 0 || width <= 0 || height <= 0 || stride <= 0 || quality < 0.0 {
		err = errors.New("webpEncodeRGB: bad arguments")
		return
	}
	if stride < width*3 && len(pix) < height*stride {
		err = errors.New("webpEncodeRGB: bad arguments")
		return
	}
	isCBuf := cbuf.Own(pix)
	cPix := cgoSafePtr(pix, isCBuf)
	defer cgoFreePtr(cPix, isCBuf)

	rv := C.cgoWebpEncodeRGB(
		(*C.uint8_t)(cPix), C.int(width), C.int(height),
		C.int(stride), C.float(quality),
	)
	if rv.ok != 1 {
		err = errors.New("webpEncodeRGB: failed")
		return
	}

	output = newCBufferFrom(unsafe.Pointer(rv.ptr), int(rv.size))
	return
}

func webpEncodeRGBA(pix []byte, width, height, stride int, quality float32, cbuf CBuffer) (output CBuffer, err error) {
	if len(pix) == 0 || width <= 0 || height <= 0 || stride <= 0 || quality < 0.0 {
		err = errors.New("webpEncodeRGBA: bad arguments")
		return
	}
	if stride < width*4 && len(pix) < height*stride {
		err = errors.New("webpEncodeRGBA: bad arguments")
		return
	}
	isCBuf := cbuf.Own(pix)
	cPix := cgoSafePtr(pix, isCBuf)
	defer cgoFreePtr(cPix, isCBuf)

	rv := C.cgoWebpEncodeRGBA(
		(*C.uint8_t)(cPix), C.int(width), C.int(height),
		C.int(stride), C.float(quality),
	)
	if rv.ok != 1 {
		err = errors.New("webpEncodeRGBA: failed")
		return
	}

	output = newCBufferFrom(unsafe.Pointer(rv.ptr), int(rv.size))
	return
}

func webpEncodeLosslessGray(pix []byte, width, height, stride int, cbuf CBuffer) (output CBuffer, err error) {
	if len(pix) == 0 || width <= 0 || height <= 0 || stride <= 0 {
		err = errors.New("webpEncodeLosslessGray: bad arguments")
		return
	}
	if stride < width*1 && len(pix) < height*stride {
		err = errors.New("webpEncodeLosslessGray: bad arguments")
		return
	}
	isCBuf := cbuf.Own(pix)
	cPix := cgoSafePtr(pix, isCBuf)
	defer cgoFreePtr(cPix, isCBuf)

	rv := C.cgoWebpEncodeLosslessGray(
		(*C.uint8_t)(cPix), C.int(width), C.int(height),
		C.int(stride),
	)
	if rv.ok != 1 {
		err = errors.New("webpEncodeLosslessGray: failed")
		return
	}

	output = newCBufferFrom(unsafe.Pointer(rv.ptr), int(rv.size))
	return
}

func webpEncodeLosslessRGB(pix []byte, width, height, stride int, cbuf CBuffer) (output CBuffer, err error) {
	if len(pix) == 0 || width <= 0 || height <= 0 || stride <= 0 {
		err = errors.New("webpEncodeLosslessRGB: bad arguments")
		return
	}
	if stride < width*3 && len(pix) < height*stride {
		err = errors.New("webpEncodeLosslessRGB: bad arguments")
		return
	}
	isCBuf := cbuf.Own(pix)
	cPix := cgoSafePtr(pix, isCBuf)
	defer cgoFreePtr(cPix, isCBuf)

	rv := C.cgoWebpEncodeLosslessRGB(
		(*C.uint8_t)(cPix), C.int(width), C.int(height),
		C.int(stride),
	)
	if rv.ok != 1 {
		err = errors.New("webpEncodeLosslessRGB: failed")
		return
	}

	output = newCBufferFrom(unsafe.Pointer(rv.ptr), int(rv.size))
	return
}

func webpEncodeLosslessRGBA(pix []byte, width, height, stride int, cbuf CBuffer) (output CBuffer, err error) {
	if len(pix) == 0 || width <= 0 || height <= 0 || stride <= 0 {
		err = errors.New("webpEncodeLosslessRGBA: bad arguments")
		return
	}
	if stride < width*4 && len(pix) < height*stride {
		err = errors.New("webpEncodeLosslessRGBA: bad arguments")
		return
	}
	isCBuf := cbuf.Own(pix)
	cPix := cgoSafePtr(pix, isCBuf)
	defer cgoFreePtr(cPix, isCBuf)

	rv := C.cgoWebpEncodeLosslessRGBA(
		(*C.uint8_t)(cPix), C.int(width), C.int(height),
		C.int(stride),
	)
	if rv.ok != 1 {
		err = errors.New("webpEncodeLosslessRGBA: failed")
		return
	}

	output = newCBufferFrom(unsafe.Pointer(rv.ptr), int(rv.size))
	return
}

func webpGetEXIF(data []byte) (metadata []byte, err error) {
	if len(data) == 0 {
		err = errors.New("webpGetEXIF: bad arguments")
		return
	}
	isCBuf := false
	cData := cgoSafePtr(data, isCBuf)
	defer cgoFreePtr(cData, isCBuf)

	rv := C.cgoWebpGetEXIF((*C.uint8_t)(cData), C.size_t(len(data)))
	if rv.ok != 1 {
		err = errors.New("webpGetEXIF: failed")
		return
	}
	metadata = C.GoBytes(unsafe.Pointer(rv.ptr), C.int(rv.size))
	C.webpFree(unsafe.Pointer(rv.ptr))
	return
}
func webpGetICCP(data []byte) (metadata []byte, err error) {
	if len(data) == 0 {
		err = errors.New("webpGetICCP: bad arguments")
		return
	}
	isCBuf := false
	cData := cgoSafePtr(data, isCBuf)
	defer cgoFreePtr(cData, isCBuf)

	rv := C.cgoWebpGetICCP((*C.uint8_t)(cData), C.size_t(len(data)))
	if rv.ok != 1 {
		err = errors.New("webpGetICCP: failed")
		return
	}
	metadata = C.GoBytes(unsafe.Pointer(rv.ptr), C.int(rv.size))
	C.webpFree(unsafe.Pointer(rv.ptr))
	return
}
func webpGetXMP(data []byte) (metadata []byte, err error) {
	if len(data) == 0 {
		err = errors.New("webpGetXMP: bad arguments")
		return
	}
	isCBuf := false
	cData := cgoSafePtr(data, isCBuf)
	defer cgoFreePtr(cData, isCBuf)

	rv := C.cgoWebpGetXMP((*C.uint8_t)(cData), C.size_t(len(data)))
	if rv.ok != 1 {
		err = errors.New("webpGetXMP: failed")
		return
	}
	metadata = C.GoBytes(unsafe.Pointer(rv.ptr), C.int(rv.size))
	C.webpFree(unsafe.Pointer(rv.ptr))
	return
}
func webpGetMetadata(data []byte, format string) (metadata []byte, err error) {
	if len(data) == 0 {
		err = errors.New("webpGetMetadata: bad arguments")
		return
	}
	isCBuf := false
	cData := cgoSafePtr(data, isCBuf)
	defer cgoFreePtr(cData, isCBuf)

	switch format {
	case "EXIF":
		rv := C.cgoWebpGetEXIF((*C.uint8_t)(cData), C.size_t(len(data)))
		if rv.ok != 1 {
			err = errors.New("webpGetMetadata: not found")
			return
		}
		metadata = C.GoBytes(unsafe.Pointer(rv.ptr), C.int(rv.size))
		C.webpFree(unsafe.Pointer(rv.ptr))
		return
	case "ICCP":
		rv := C.cgoWebpGetICCP((*C.uint8_t)(cData), C.size_t(len(data)))
		if rv.ok != 1 {
			err = errors.New("webpGetMetadata: not found")
			return
		}
		metadata = C.GoBytes(unsafe.Pointer(rv.ptr), C.int(rv.size))
		C.webpFree(unsafe.Pointer(rv.ptr))
		return
	case "XMP":
		rv := C.cgoWebpGetXMP((*C.uint8_t)(cData), C.size_t(len(data)))
		if rv.ok != 1 {
			err = errors.New("webpGetMetadata: not found")
			return
		}
		metadata = C.GoBytes(unsafe.Pointer(rv.ptr), C.int(rv.size))
		C.webpFree(unsafe.Pointer(rv.ptr))
		return
	default:
		err = errors.New("webpGetMetadata: unknown format")
		return
	}
}

func webpSetEXIF(data, metadata []byte) (newData []byte, err error) {
	if len(data) == 0 || len(metadata) == 0 {
		err = errors.New("webpSetEXIF: bad arguments")
		return
	}
	isCBuf := false
	cData := cgoSafePtr(data, isCBuf)
	defer cgoFreePtr(cData, isCBuf)
	cMetadata := cgoSafePtr(metadata, isCBuf)
	defer cgoFreePtr(cMetadata, isCBuf)

	rv := C.cgoWebpSetEXIF(
		(*C.uint8_t)(cData), C.size_t(len(data)),
		(*C.char)(cMetadata), C.size_t(len(metadata)),
	)
	if rv.ok != 1 {
		err = errors.New("webpSetEXIF: failed")
		return
	}
	newData = C.GoBytes(unsafe.Pointer(rv.ptr), C.int(rv.size))
	C.webpFree(unsafe.Pointer(rv.ptr))
	return
}
func webpSetICCP(data, metadata []byte) (newData []byte, err error) {
	if len(data) == 0 || len(metadata) == 0 {
		err = errors.New("webpSetICCP: bad arguments")
		return
	}
	isCBuf := false
	cData := cgoSafePtr(data, isCBuf)
	defer cgoFreePtr(cData, isCBuf)
	cMetadata := cgoSafePtr(metadata, isCBuf)
	defer cgoFreePtr(cMetadata, isCBuf)

	rv := C.cgoWebpSetICCP(
		(*C.uint8_t)(cData), C.size_t(len(data)),
		(*C.char)(cMetadata), C.size_t(len(metadata)),
	)
	if rv.ok != 1 {
		err = errors.New("webpSetICCP: failed")
		return
	}
	newData = C.GoBytes(unsafe.Pointer(rv.ptr), C.int(rv.size))
	C.webpFree(unsafe.Pointer(rv.ptr))
	return
}
func webpSetXMP(data, metadata []byte) (newData []byte, err error) {
	if len(data) == 0 || len(metadata) == 0 {
		err = errors.New("webpSetXMP: bad arguments")
		return
	}
	isCBuf := false
	cData := cgoSafePtr(data, isCBuf)
	defer cgoFreePtr(cData, isCBuf)
	cMetadata := cgoSafePtr(metadata, isCBuf)
	defer cgoFreePtr(cMetadata, isCBuf)

	rv := C.cgoWebpSetXMP(
		(*C.uint8_t)(cData), C.size_t(len(data)),
		(*C.char)(cMetadata), C.size_t(len(metadata)),
	)
	if rv.ok != 1 {
		err = errors.New("webpSetXMP: failed")
		return
	}
	newData = C.GoBytes(unsafe.Pointer(rv.ptr), C.int(rv.size))
	C.webpFree(unsafe.Pointer(rv.ptr))
	return
}
func webpSetMetadata(data, metadata []byte, format string) (newData []byte, err error) {
	if len(data) == 0 || len(metadata) == 0 {
		err = errors.New("webpSetMetadata: bad arguments")
		return
	}
	isCBuf := false
	cData := cgoSafePtr(data, isCBuf)
	defer cgoFreePtr(cData, isCBuf)
	cMetadata := cgoSafePtr(metadata, isCBuf)
	defer cgoFreePtr(cMetadata, isCBuf)

	switch format {
	case "EXIF":
		rv := C.cgoWebpSetEXIF(
			(*C.uint8_t)(cData), C.size_t(len(data)),
			(*C.char)(cMetadata), C.size_t(len(metadata)),
		)
		if rv.ok != 1 {
			err = errors.New("webpSetMetadata: failed")
			return
		}
		newData = C.GoBytes(unsafe.Pointer(rv.ptr), C.int(rv.size))
		C.webpFree(unsafe.Pointer(rv.ptr))
		return
	case "ICCP":
		rv := C.cgoWebpSetICCP(
			(*C.uint8_t)(cData), C.size_t(len(data)),
			(*C.char)(cMetadata), C.size_t(len(metadata)),
		)
		if rv.ok != 1 {
			err = errors.New("webpSetMetadata: failed")
			return
		}
		newData = C.GoBytes(unsafe.Pointer(rv.ptr), C.int(rv.size))
		C.webpFree(unsafe.Pointer(rv.ptr))
		return
	case "XMP":
		rv := C.cgoWebpSetXMP(
			(*C.uint8_t)(cData), C.size_t(len(data)),
			(*C.char)(cMetadata), C.size_t(len(metadata)),
		)
		if rv.ok != 1 {
			err = errors.New("webpSetMetadata: failed")
			return
		}
		newData = C.GoBytes(unsafe.Pointer(rv.ptr), C.int(rv.size))
		C.webpFree(unsafe.Pointer(rv.ptr))
		return
	default:
		err = errors.New("webpSetMetadata: unknown format")
		return
	}
}

func webpDelEXIF(data []byte) (newData []byte, err error) {
	if len(data) == 0 {
		err = errors.New("webpDelEXIF: bad arguments")
		return
	}
	isCBuf := false
	cData := cgoSafePtr(data, isCBuf)
	defer cgoFreePtr(cData, isCBuf)

	rv := C.cgoWebpDelEXIF(
		(*C.uint8_t)(cData), C.size_t(len(data)),
	)
	if rv.ok != 1 {
		err = errors.New("webpDelEXIF: failed")
		return
	}
	newData = C.GoBytes(unsafe.Pointer(rv.ptr), C.int(rv.size))
	C.webpFree(unsafe.Pointer(rv.ptr))
	return
}
func webpDelICCP(data []byte) (newData []byte, err error) {
	if len(data) == 0 {
		err = errors.New("webpDelICCP: bad arguments")
		return
	}
	isCBuf := false
	cData := cgoSafePtr(data, isCBuf)
	defer cgoFreePtr(cData, isCBuf)

	rv := C.cgoWebpDelICCP(
		(*C.uint8_t)(cData), C.size_t(len(data)),
	)
	if rv.ok != 1 {
		err = errors.New("webpDelICCP: failed")
		return
	}
	newData = C.GoBytes(unsafe.Pointer(rv.ptr), C.int(rv.size))
	C.webpFree(unsafe.Pointer(rv.ptr))
	return
}
func webpDelXMP(data []byte) (newData []byte, err error) {
	if len(data) == 0 {
		err = errors.New("webpDelXMP: bad arguments")
		return
	}
	isCBuf := false
	cData := cgoSafePtr(data, isCBuf)
	defer cgoFreePtr(cData, isCBuf)

	rv := C.cgoWebpDelXMP(
		(*C.uint8_t)(cData), C.size_t(len(data)),
	)
	if rv.ok != 1 {
		err = errors.New("webpDelXMP: failed")
		return
	}
	newData = C.GoBytes(unsafe.Pointer(rv.ptr), C.int(rv.size))
	C.webpFree(unsafe.Pointer(rv.ptr))
	return
}
