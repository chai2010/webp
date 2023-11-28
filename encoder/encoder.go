// The MIT License (MIT)
//
// Copyright (c) 2019 Amangeldy Kadyl
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
// the Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
// FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
// COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
// IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package encoder

/*

#cgo CFLAGS: -I../internal/libwebp-1.3.2/
#cgo CFLAGS: -I../internal/libwebp-1.3.2/src/
#cgo CFLAGS: -I../internal/include/
#cgo CFLAGS: -Wno-pointer-sign -DWEBP_USE_THREAD
#cgo !windows LDFLAGS: -lm

#include <stdlib.h>
#include <webp/encode.h>
static uint8_t* encodeNRBBA(WebPConfig* config, const uint8_t* rgba, int width, int height, int stride, size_t* output_size) {
	WebPPicture pic;
	WebPMemoryWriter wrt;
	int ok;

	if (!WebPPictureInit(&pic)) {
		return NULL;
	}

	pic.use_argb = 1;
	pic.width = width;
	pic.height = height;
	pic.writer = WebPMemoryWrite;
	pic.custom_ptr = &wrt;
	WebPMemoryWriterInit(&wrt);

	ok = WebPPictureImportRGBA(&pic, rgba, stride) && WebPEncode(config, &pic);
	WebPPictureFree(&pic);

	if (!ok) {
		WebPMemoryWriterClear(&wrt);
		return NULL;
	}

	*output_size = wrt.size;
	return wrt.mem;
}
*/
import "C"
import (
	"errors"
	"image"
	"image/draw"
	"io"
	"unsafe"
)

// Encoder stores information to encode image
type Encoder struct {
	options *Options
	config  *C.WebPConfig
	img     *image.NRGBA
}

// NewEncoder return new encoder instance
func NewEncoder(src image.Image, options *Options) (e *Encoder, err error) {
	var config *C.WebPConfig

	if options == nil {
		options, _ = NewLossyEncoderOptions(PresetDefault, 75)
	}
	if config, err = options.GetConfig(); err != nil {
		return nil, err
	}

	e = &Encoder{options: options, config: config}

	switch v := src.(type) {
	case *image.NRGBA:
		e.img = v
	default:
		e.img = e.convertToNRGBA(src)
	}

	return
}

// Encode picture and flush to io.Writer
func (e *Encoder) Encode(w io.Writer) error {
	var size C.size_t

	output := C.encodeNRBBA(
		e.config,
		(*C.uint8_t)(&e.img.Pix[0]),
		C.int(e.img.Rect.Dx()),
		C.int(e.img.Rect.Dy()),
		C.int(e.img.Stride),
		&size,
	)

	if output == nil || size == 0 {
		return errors.New("cannot encode webppicture")
	}
	defer C.free(unsafe.Pointer(output))

	_, err := w.Write(((*[1 << 30]byte)(unsafe.Pointer(output)))[0:int(size):int(size)])

	return err
}

// Convert picture from any image.Image type to *image.NRGBA
// @todo optimization needed
func (e *Encoder) convertToNRGBA(src image.Image) (dst *image.NRGBA) {
	dst = image.NewNRGBA(src.Bounds())
	draw.Draw(dst, dst.Bounds(), src, src.Bounds().Min, draw.Src)

	return
}
