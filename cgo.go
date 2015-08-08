// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build cgo

package webp

// #include <stdlib.h>
import "C"
import "unsafe"

// Go1.3: Changes to the garbage collector
// http://golang.org/doc/go1.3#garbage_collector

func cgoSafePtr(data []byte, isCBuf bool) unsafe.Pointer {
	if len(data) == 0 {
		return nil
	}
	if !isCBuf && cgoIsUnsafePtr {
		p := C.malloc(C.size_t(len(data)))
		copy(((*[1 << 30]byte)(p))[0:len(data):len(data)], data)
		return p
	} else {
		p := unsafe.Pointer(&data[0])
		return p
	}
}

func cgoFreePtr(p unsafe.Pointer, isCBuf bool) {
	if !isCBuf && cgoIsUnsafePtr && p != nil {
		C.free(p)
	}
}
