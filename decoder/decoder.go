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

package decoder

/*
#cgo CFLAGS: -I../internal/libwebp-1.3.2/
#cgo CFLAGS: -I../internal/libwebp-1.3.2/src/
#cgo CFLAGS: -I../internal/include/
#cgo CFLAGS: -Wno-pointer-sign -DWEBP_USE_THREAD
#cgo !windows LDFLAGS: -lm

#include <stdlib.h>
#include <webp/decode.h>
*/
import "C"
import (
	"errors"
	"fmt"
	"image"
	"io"
	"unsafe"

	"github.com/cqzcqq/webp/utils"
)

// Decoder stores information to decode picture
type Decoder struct {
	data    []byte
	options *Options
	config  *C.WebPDecoderConfig
	dPtr    *C.uint8_t
	sPtr    C.size_t
}

// NewDecoder return new decoder instance
func NewDecoder(r io.Reader, options *Options) (d *Decoder, err error) {
	var data []byte

	if data, err = io.ReadAll(r); err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, errors.New("data is empty")
	}

	if options == nil {
		options = &Options{}
	}
	d = &Decoder{data: data, options: options}

	if d.config, err = d.options.GetConfig(); err != nil {
		return nil, err
	}

	d.dPtr = (*C.uint8_t)(&d.data[0])
	d.sPtr = (C.size_t)(len(d.data))

	// получаем WebPBitstreamFeatures
	if status := d.parseFeatures(d.dPtr, d.sPtr); status != utils.Vp8StatusOk {
		return nil, errors.New(fmt.Sprintf("cannot fetch features: %s", status.String()))
	}

	return
}

// Decode picture from reader
func (d *Decoder) Decode() (image.Image, error) {
	// вписываем размеры итоговой картинки
	d.config.output.width, d.config.output.height = d.getOutputDimensions()
	// указываем что декодируем в RGBA
	d.config.output.colorspace = C.MODE_RGBA
	d.config.output.is_external_memory = 1

	img := image.NewNRGBA(image.Rectangle{Max: image.Point{
		X: int(d.config.output.width),
		Y: int(d.config.output.height),
	}})

	buff := (*C.WebPRGBABuffer)(unsafe.Pointer(&d.config.output.u[0]))
	buff.stride = C.int(img.Stride)
	buff.rgba = (*C.uint8_t)(&img.Pix[0])
	buff.size = (C.size_t)(len(img.Pix))

	if status := utils.VP8StatusCode(C.WebPDecode(d.dPtr, d.sPtr, d.config)); status != utils.Vp8StatusOk {
		return nil, errors.New(fmt.Sprintf("cannot decode picture: %s", status.String()))
	}

	return img, nil
}

// GetFeatures return information about picture: width, height ...
func (d *Decoder) GetFeatures() utils.BitstreamFeatures {
	return utils.BitstreamFeatures{
		Width:        int(d.config.input.width),
		Height:       int(d.config.input.height),
		HasAlpha:     int(d.config.input.has_alpha) == 1,
		HasAnimation: int(d.config.input.has_animation) == 1,
		Format:       utils.FormatType(d.config.input.format),
	}
}

// parse features from picture
func (d *Decoder) parseFeatures(dataPtr *C.uint8_t, sizePtr C.size_t) utils.VP8StatusCode {
	return utils.VP8StatusCode(C.WebPGetFeatures(dataPtr, sizePtr, &d.config.input))
}

// return dimensions of result image
func (d *Decoder) getOutputDimensions() (width, height C.int) {
	width = d.config.input.width
	height = d.config.input.height

	if d.config.options.use_scaling > 0 {
		width = d.config.options.scaled_width
		height = d.config.options.scaled_height
	} else if d.config.options.use_cropping > 0 {
		width = d.config.options.crop_width
		height = d.config.options.crop_height
	}

	return
}
