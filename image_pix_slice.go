// Copyright 2015 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package webp

import (
	"reflect"
	"unsafe"
)

type PixSlice []byte

// AsPixSilce convert a normal slice to byte slice.
//
// Convert []X to []byte:
//
//	x := make([]X, xLen)
//	y := AsPixSilce(x)
//
func AsPixSilce(slice interface{}) (d PixSlice) {
	sv := reflect.ValueOf(slice)
	h := (*reflect.SliceHeader)((unsafe.Pointer(&d)))
	h.Cap = sv.Cap() * int(sv.Type().Elem().Size())
	h.Len = sv.Len() * int(sv.Type().Elem().Size())
	h.Data = sv.Pointer()
	return
}

// Slice convert a normal slice to new type slice.
//
// Convert []byte to []Y:
//	x := make([]byte, xLen)
//	y := PixSlice(x).Slice(reflect.TypeOf([]Y(nil))).([]Y)
//
func (d PixSlice) Slice(newSliceType reflect.Type) interface{} {
	sv := reflect.ValueOf(d)
	newSlice := reflect.New(newSliceType)
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(newSlice.Pointer()))
	hdr.Cap = sv.Cap() * int(sv.Type().Elem().Size()) / int(newSliceType.Elem().Size())
	hdr.Len = sv.Len() * int(sv.Type().Elem().Size()) / int(newSliceType.Elem().Size())
	hdr.Data = uintptr(sv.Pointer())
	return newSlice.Elem().Interface()
}

func (d PixSlice) Bytes() (v []byte) {
	return d
}

func (d PixSlice) Int8s() (v []int8) {
	h0 := (*reflect.SliceHeader)(unsafe.Pointer(&d))
	h1 := (*reflect.SliceHeader)(unsafe.Pointer(&v))

	h1.Cap = h0.Cap
	h1.Len = h0.Len
	h1.Data = h0.Data
	return
}

func (d PixSlice) Int16s() (v []int16) {
	h0 := (*reflect.SliceHeader)(unsafe.Pointer(&d))
	h1 := (*reflect.SliceHeader)(unsafe.Pointer(&v))

	h1.Cap = h0.Cap / 2
	h1.Len = h0.Len / 2
	h1.Data = h0.Data
	return
}

func (d PixSlice) Int32s() (v []int32) {
	h0 := (*reflect.SliceHeader)(unsafe.Pointer(&d))
	h1 := (*reflect.SliceHeader)(unsafe.Pointer(&v))

	h1.Cap = h0.Cap / 4
	h1.Len = h0.Len / 4
	h1.Data = h0.Data
	return
}

func (d PixSlice) Int64s() (v []int64) {
	h0 := (*reflect.SliceHeader)(unsafe.Pointer(&d))
	h1 := (*reflect.SliceHeader)(unsafe.Pointer(&v))

	h1.Cap = h0.Cap / 8
	h1.Len = h0.Len / 8
	h1.Data = h0.Data
	return
}

func (d PixSlice) Uint8s() []uint8 {
	return d
}

func (d PixSlice) Uint16s() (v []uint16) {
	h0 := (*reflect.SliceHeader)(unsafe.Pointer(&d))
	h1 := (*reflect.SliceHeader)(unsafe.Pointer(&v))

	h1.Cap = h0.Cap / 2
	h1.Len = h0.Len / 2
	h1.Data = h0.Data
	return
}

func (d PixSlice) Uint32s() (v []uint32) {
	h0 := (*reflect.SliceHeader)(unsafe.Pointer(&d))
	h1 := (*reflect.SliceHeader)(unsafe.Pointer(&v))

	h1.Cap = h0.Cap / 4
	h1.Len = h0.Len / 4
	h1.Data = h0.Data
	return
}

func (d PixSlice) Uint64s() (v []uint64) {
	h0 := (*reflect.SliceHeader)(unsafe.Pointer(&d))
	h1 := (*reflect.SliceHeader)(unsafe.Pointer(&v))

	h1.Cap = h0.Cap / 8
	h1.Len = h0.Len / 8
	h1.Data = h0.Data
	return
}

func (d PixSlice) Float32s() (v []float32) {
	h0 := (*reflect.SliceHeader)(unsafe.Pointer(&d))
	h1 := (*reflect.SliceHeader)(unsafe.Pointer(&v))

	h1.Cap = h0.Cap / 4
	h1.Len = h0.Len / 4
	h1.Data = h0.Data
	return
}

func (d PixSlice) Float64s() (v []float64) {
	h0 := (*reflect.SliceHeader)(unsafe.Pointer(&d))
	h1 := (*reflect.SliceHeader)(unsafe.Pointer(&v))

	h1.Cap = h0.Cap / 8
	h1.Len = h0.Len / 8
	h1.Data = h0.Data
	return
}

func (d PixSlice) Complex64s() (v []complex64) {
	h0 := (*reflect.SliceHeader)(unsafe.Pointer(&d))
	h1 := (*reflect.SliceHeader)(unsafe.Pointer(&v))

	h1.Cap = h0.Cap / 16
	h1.Len = h0.Len / 16
	h1.Data = h0.Data
	return
}

func (d PixSlice) Complex128s() (v []complex128) {
	h0 := (*reflect.SliceHeader)(unsafe.Pointer(&d))
	h1 := (*reflect.SliceHeader)(unsafe.Pointer(&v))

	h1.Cap = h0.Cap / 32
	h1.Len = h0.Len / 32
	h1.Data = h0.Data
	return
}

func (d PixSlice) Value(i int, dataType reflect.Kind) float64 {
	switch dataType {
	case reflect.Int8:
		return float64(d.Int8s()[i])
	case reflect.Int16:
		return float64(d.Int16s()[i])
	case reflect.Int32:
		return float64(d.Int32s()[i])
	case reflect.Int64:
		return float64(d.Int64s()[i])
	case reflect.Uint8:
		return float64(d[i])
	case reflect.Uint16:
		return float64(d.Uint16s()[i])
	case reflect.Uint32:
		return float64(d.Uint32s()[i])
	case reflect.Uint64:
		return float64(d.Uint64s()[i])
	case reflect.Float32:
		return float64(d.Float32s()[i])
	case reflect.Float64:
		return float64(d.Float64s()[i])
	case reflect.Complex64:
		return float64(real(d.Complex64s()[i]))
	case reflect.Complex128:
		return float64(real(d.Complex128s()[i]))
	}
	return 0
}

func (d PixSlice) SetValue(i int, dataType reflect.Kind, v float64) {
	switch dataType {
	case reflect.Int8:
		d.Int8s()[i] = int8(v)
	case reflect.Int16:
		d.Int16s()[i] = int16(v)
	case reflect.Int32:
		d.Int32s()[i] = int32(v)
	case reflect.Int64:
		d.Int64s()[i] = int64(v)
	case reflect.Uint8:
		d[i] = byte(v)
	case reflect.Uint16:
		d.Uint16s()[i] = uint16(v)
	case reflect.Uint32:
		d.Uint32s()[i] = uint32(v)
	case reflect.Uint64:
		d.Uint64s()[i] = uint64(v)
	case reflect.Float32:
		d.Float32s()[i] = float32(v)
	case reflect.Float64:
		d.Float64s()[i] = float64(v)
	case reflect.Complex64:
		d.Complex64s()[i] = complex(float32(v), 0)
	case reflect.Complex128:
		d.Complex128s()[i] = complex(float64(v), 0)
	}
}

func (d PixSlice) SwapEndian(dataType reflect.Kind) {
	switch dataType {
	case reflect.Int16, reflect.Uint16:
		for i := 0; i+2-1 < len(d); i = i + 2 {
			d[i+0], d[i+1] = d[i+1], d[i+0]
		}
	case reflect.Int32, reflect.Uint32, reflect.Float32, reflect.Complex64:
		for i := 0; i+4-1 < len(d); i = i + 4 {
			d[i+0], d[i+3] = d[i+3], d[i+0]
			d[i+1], d[i+2] = d[i+2], d[i+1]
		}
	case reflect.Int64, reflect.Uint64, reflect.Float64, reflect.Complex128:
		for i := 0; i+8-1 < len(d); i = i + 8 {
			d[i+0], d[i+7] = d[i+7], d[i+0]
			d[i+1], d[i+6] = d[i+6], d[i+1]
			d[i+2], d[i+5] = d[i+5], d[i+2]
			d[i+3], d[i+4] = d[i+4], d[i+3]
		}
	}
}
