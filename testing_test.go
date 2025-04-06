// Copyright 2015 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// testing helper, please fork this file, and fix the package name.
//
// See http://github.com/chai2010/assert

package webp

import (
	"fmt"
	"reflect"
	"testing"
)

func tIsIntType(v interface{}) bool {
	switch reflect.ValueOf(v).Kind() {
	case reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64:
		return true
	}
	return false
}

func tIsUintType(v interface{}) bool {
	switch reflect.ValueOf(v).Kind() {
	case reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64,
		reflect.Uintptr:
		return true
	}
	return false
}

func tIsFloatType(v interface{}) bool {
	switch reflect.ValueOf(v).Kind() {
	case reflect.Float32,
		reflect.Float64:
		return true
	}
	return false
}

func tIsNumberType(v interface{}) bool {
	switch reflect.ValueOf(v).Kind() {
	case reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64,
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64,
		reflect.Uintptr,
		reflect.Float32,
		reflect.Float64,
		reflect.Complex64,
		reflect.Complex128:
		return true
	}
	return false
}

func tIsNumberEqual(a, b interface{}) bool {
	if tIsNumberType(a) && tIsNumberType(b) {
		return fmt.Sprintf("%v", a) == fmt.Sprintf("%v", b)
	}
	return false
}

func tAssert(t testing.TB, condition bool, args ...interface{}) {
	t.Helper()
	if !condition {
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("tAssert failed, %s", msg)
		} else {
			t.Fatalf("tAssert failed")
		}
	}
}

func tAssertNil(t testing.TB, p interface{}, args ...interface{}) {
	if p != nil {
		t.Helper()
		if msg := fmt.Sprint(args...); msg != "" {
			if err, ok := p.(error); ok && err != nil {
				t.Fatalf("tAssertNil failed, err = %v, %s", err, msg)
			} else {
				t.Fatalf("tAssertNil failed, %s", msg)
			}
		} else {
			if err, ok := p.(error); ok && err != nil {
				t.Fatalf("tAssertNil failed, err = %v", err)
			} else {
				t.Fatalf("tAssertNil failed")
			}
		}
	}
}

func tAssertEQ(t testing.TB, expected, got interface{}, args ...interface{}) {
	t.Helper()
	if !reflect.DeepEqual(expected, got) && !tIsNumberEqual(expected, got) {
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("tAssertEQ failed, expected = %v, got = %v, %s", expected, got, msg)
		} else {
			t.Fatalf("tAssertEQ failed, expected = %v, got = %v", expected, got)
		}
	}
}
