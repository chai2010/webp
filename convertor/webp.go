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

package convertor

import (
	"image"
	"image/color"
	"io"

	"github.com/cqzcqq/webp/decoder"
	"github.com/cqzcqq/webp/encoder"
)

func init() {
	image.RegisterFormat("webp", "RIFF????WEBPVP8", quickDecode, quickDecodeConfig)
}

func quickDecode(r io.Reader) (image.Image, error) {
	return Decode(r, &decoder.Options{})
}

func quickDecodeConfig(r io.Reader) (image.Config, error) {
	return DecodeConfig(r, &decoder.Options{})
}

// Decode picture from reader
func Decode(r io.Reader, options *decoder.Options) (image.Image, error) {
	if dec, err := decoder.NewDecoder(r, options); err != nil {
		return nil, err
	} else {
		return dec.Decode()
	}
}

func DecodeConfig(r io.Reader, options *decoder.Options) (image.Config, error) {
	if dec, err := decoder.NewDecoder(r, options); err != nil {
		return image.Config{}, err
	} else {
		return image.Config{
			ColorModel: color.RGBAModel,
			Width:      dec.GetFeatures().Width,
			Height:     dec.GetFeatures().Height,
		}, nil
	}
}

// Encode picture and write to io.Writer
func Encode(w io.Writer, src image.Image, options *encoder.Options) error {
	if enc, err := encoder.NewEncoder(src, options); err != nil {
		return err
	} else {
		return enc.Encode(w)
	}
}
