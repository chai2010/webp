// Copyright 2015 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package webp

import (
	"fmt"
	"image"
	"image/color"
	"reflect"
)

func ExampleAsPixSilce() {
	a := []int32{101, 102, 103}
	b := AsPixSilce(a)

	b.Int32s()[0] = 12345
	b.SetValue(1, reflect.Int32, 1002)
	b.SetValue(2, reflect.Int32, 1003.5)
	fmt.Printf("len(b) = %d\n", len(b))
	fmt.Printf("b.Int32s() = %v\n", b.Int32s())
	// Output:
	// len(b) = 12
	// b.Int32s() = [12345 1002 1003]
}

func ExampleAsPixSilce_swapEndian() {
	rgba64 := image.NewRGBA64(image.Rect(0, 0, 1, 1))
	rgba64.SetRGBA64(0, 0, color.RGBA64{
		R: 0x0102,
		G: 0x0304,
		B: 0x0506,
		A: 0x0708,
	})

	// pix is big-endian format
	fmt.Printf("big-endian: %v\n", rgba64.Pix)

	AsPixSilce(rgba64.Pix).SwapEndian(reflect.Uint16)
	fmt.Printf("little-endian: %v\n", rgba64.Pix)

	// Output:
	// big-endian: [1 2 3 4 5 6 7 8]
	// little-endian: [2 1 4 3 6 5 8 7]
}

func ExampleColor() {
	c := MemPColor{
		Channels: 4,
		DataType: reflect.Uint8,
		Pix:      []byte{101, 102, 103, 104},
	}
	rgba := color.RGBAModel.Convert(c).(color.RGBA)
	fmt.Printf("rgba = %v\n", rgba)
	// Output:
	// rgba = {101 102 103 104}
}

func ExampleColor_uint16() {
	c := MemPColor{
		Channels: 4,
		DataType: reflect.Uint16,
		Pix:      AsPixSilce([]uint16{11101, 11102, 11103, 11104}),
	}
	rgba64 := color.RGBA64Model.Convert(c).(color.RGBA64)
	fmt.Printf("rgba64 = %v\n", rgba64)
	// Output:
	// rgba64 = {11101 11102 11103 11104}
}

func ExampleColorModel() {
	rgba := color.RGBA{R: 101, G: 102, B: 103, A: 104}
	c := ColorModel(4, reflect.Uint8).Convert(rgba).(MemPColor)
	fmt.Printf("c = %v\n", c)
	// Output:
	// c = {4 uint8 [101 102 103 104]}
}

func ExampleSizeofKind() {
	fmt.Printf("%v = %v\n", reflect.Uint8, SizeofKind(reflect.Uint8))
	fmt.Printf("%v = %v\n", reflect.Uint16, SizeofKind(reflect.Uint16))
	fmt.Printf("%v = %v\n", reflect.Uint32, SizeofKind(reflect.Uint32))
	fmt.Printf("%v = %v\n", reflect.Float32, SizeofKind(reflect.Float32))
	fmt.Printf("%v = %v\n", reflect.Float64, SizeofKind(reflect.Float64))
	// Output:
	// uint8 = 1
	// uint16 = 2
	// uint32 = 4
	// float32 = 4
	// float64 = 8
}

func ExampleSizeofPixel() {
	fmt.Printf("sizeof(gray) = %d\n", SizeofPixel(1, reflect.Uint8))
	fmt.Printf("sizeof(gray16) = %d\n", SizeofPixel(1, reflect.Uint16))
	fmt.Printf("sizeof(rgb) = %d\n", SizeofPixel(3, reflect.Uint8))
	fmt.Printf("sizeof(rgb48) = %d\n", SizeofPixel(3, reflect.Uint16))
	fmt.Printf("sizeof(rgba) = %d\n", SizeofPixel(4, reflect.Uint8))
	fmt.Printf("sizeof(rgba64) = %d\n", SizeofPixel(4, reflect.Uint16))
	fmt.Printf("sizeof(float32) = %d\n", SizeofPixel(1, reflect.Float32))
	// Output:
	// sizeof(gray) = 1
	// sizeof(gray16) = 2
	// sizeof(rgb) = 3
	// sizeof(rgb48) = 6
	// sizeof(rgba) = 4
	// sizeof(rgba64) = 8
	// sizeof(float32) = 4
}

func ExampleImage_rgb() {
	type RGB struct {
		R, G, B uint8
	}

	b := image.Rect(0, 0, 300, 400)
	rgbImage := NewMemPImage(b, 3, reflect.Uint8)

	for y := b.Min.Y; y < b.Max.Y; y++ {
		var (
			line     []byte = rgbImage.XPix[rgbImage.PixOffset(b.Min.X, y):][:rgbImage.XStride]
			rgbSlice []RGB  = PixSlice(line).Slice(reflect.TypeOf([]RGB(nil))).([]RGB)
		)

		for i, _ := range rgbSlice {
			rgbSlice[i] = RGB{
				R: uint8(i + 1),
				G: uint8(i + 2),
				B: uint8(i + 3),
			}
		}
	}
}

func ExampleImage_rgb48() {
	type RGB struct {
		R, G, B uint16
	}

	b := image.Rect(0, 0, 300, 400)
	rgbImage := NewMemPImage(b, 3, reflect.Uint16)

	for y := b.Min.Y; y < b.Max.Y; y++ {
		var (
			line     []byte = rgbImage.XPix[rgbImage.PixOffset(b.Min.X, y):][:rgbImage.XStride]
			rgbSlice []RGB  = PixSlice(line).Slice(reflect.TypeOf([]RGB(nil))).([]RGB)
		)

		for i, _ := range rgbSlice {
			rgbSlice[i] = RGB{
				R: uint16(i + 1),
				G: uint16(i + 2),
				B: uint16(i + 3),
			}
		}
	}
}

func ExampleImage_unsafe() {
	// struct must same as Image
	type MyImage struct {
		MemPMagic string // MemP
		Rect      image.Rectangle
		Channels  int
		DataType  reflect.Kind
		Pix       []byte

		// Stride is the Pix stride (in bytes, must align with PixelSize)
		// between vertically adjacent pixels.
		Stride int
	}

	p := &MyImage{
		MemPMagic: MemPMagic,
		Rect:      image.Rect(0, 0, 300, 400),
		Channels:  3,
		DataType:  reflect.Uint16,
		Pix:       make([]byte, 300*400*3*SizeofKind(reflect.Uint16)),
		Stride:    300 * 3 * SizeofKind(reflect.Uint16),
	}

	q, _ := AsMemPImage(p)
	_ = q
}
