// Copyright 2015 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// testing helper, please fork this file, and fix the package name.
//
// See http://github.com/chai2010/assert

package webp

import (
	"fmt"
	"math"
	"os"
	"reflect"
	"regexp"
	"runtime"
	"strings"
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

func tCallerFileLine(skip int) (file string, line int) {
	_, file, line, ok := runtime.Caller(skip + 1)
	if ok {
		// Truncate file name at last file name separator.
		if index := strings.LastIndex(file, "/"); index >= 0 {
			file = file[index+1:]
		} else if index = strings.LastIndex(file, "\\"); index >= 0 {
			file = file[index+1:]
		}
	} else {
		file = "???"
		line = 1
	}
	return
}

func tAssert(t testing.TB, condition bool, args ...interface{}) {
	if !condition {
		file, line := tCallerFileLine(1)
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("%s:%d: tAssert failed, %s", file, line, msg)
		} else {
			t.Fatalf("%s:%d: tAssert failed", file, line)
		}
	}
}

func tAssertf(t *testing.T, condition bool, format string, a ...interface{}) {
	if !condition {
		file, line := tCallerFileLine(1)
		if msg := fmt.Sprintf(format, a...); msg != "" {
			t.Fatalf("%s:%d: tAssert failed, %s", file, line, msg)
		} else {
			t.Fatalf("%s:%d: tAssert failed", file, line)
		}
	}
}

func tAssertNil(t testing.TB, p interface{}, args ...interface{}) {
	if p != nil {
		file, line := tCallerFileLine(1)
		if msg := fmt.Sprint(args...); msg != "" {
			if err, ok := p.(error); ok && err != nil {
				t.Fatalf("%s:%d: tAssertNil failed, err = %v, %s", file, line, err, msg)
			} else {
				t.Fatalf("%s:%d: tAssertNil failed, %s", file, line, msg)
			}
		} else {
			if err, ok := p.(error); ok && err != nil {
				t.Fatalf("%s:%d: tAssertNil failed, err = %v", file, line, err)
			} else {
				t.Fatalf("%s:%d: tAssertNil failed", file, line)
			}
		}
	}
}

func tAssertNotNil(t testing.TB, p interface{}, args ...interface{}) {
	if p == nil {
		file, line := tCallerFileLine(1)
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("%s:%d: tAssertNotNil failed, %s", file, line, msg)
		} else {
			t.Fatalf("%s:%d: tAssertNotNil failed", file, line)
		}
	}
}

func tAssertTrue(t testing.TB, condition bool, args ...interface{}) {
	if condition != true {
		file, line := tCallerFileLine(1)
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("%s:%d: tAssertTrue failed, %s", file, line, msg)
		} else {
			t.Fatalf("%s:%d: tAssertTrue failed", file, line)
		}
	}
}

func tAssertFalse(t testing.TB, condition bool, args ...interface{}) {
	if condition != false {
		file, line := tCallerFileLine(1)
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("%s:%d: tAssertFalse failed, %s", file, line, msg)
		} else {
			t.Fatalf("%s:%d: tAssertFalse failed", file, line)
		}
	}
}

func tAssertEqual(t testing.TB, expected, got interface{}, args ...interface{}) {
	if !reflect.DeepEqual(expected, got) {
		file, line := tCallerFileLine(1)
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("%s:%d: tAssertEqual failed, expected = %v, got = %v, %s", file, line, expected, got, msg)
		} else {
			t.Fatalf("%s:%d: tAssertEqual failed, expected = %v, got = %v", file, line, expected, got)
		}
	}
}

func tAssertNotEqual(t testing.TB, expected, got interface{}, args ...interface{}) {
	if reflect.DeepEqual(expected, got) {
		file, line := tCallerFileLine(1)
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("%s:%d: tAssertNotEqual failed, expected = %v, got = %v, %s", file, line, expected, got, msg)
		} else {
			t.Fatalf("%s:%d: tAssertNotEqual failed, expected = %v, got = %v", file, line, expected, got)
		}
	}
}

func tAssertNear(t testing.TB, expected, got, abs float64, args ...interface{}) {
	if math.Abs(expected-got) > abs {
		file, line := tCallerFileLine(1)
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("%s:%d: tAssertNear failed, expected = %v, got = %v, abs = %v, %s", file, line, expected, got, abs, msg)
		} else {
			t.Fatalf("%s:%d: tAssertNear failed, expected = %v, got = %v, abs = %v", file, line, expected, got, abs)
		}
	}
}

func tAssertBetween(t testing.TB, min, max, val float64, args ...interface{}) {
	if val < min || max < val {
		file, line := tCallerFileLine(1)
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("%s:%d: tAssertBetween failed, min = %v, max = %v, val = %v, %s", file, line, min, max, val, msg)
		} else {
			t.Fatalf("%s:%d: tAssertBetween failed, min = %v, max = %v, val = %v", file, line, min, max, val)
		}
	}
}

func tAssertNotBetween(t testing.TB, min, max, val float64, args ...interface{}) {
	if min <= val && val <= max {
		file, line := tCallerFileLine(1)
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("%s:%d: tAssertNotBetween failed, min = %v, max = %v, val = %v, %s", file, line, min, max, val, msg)
		} else {
			t.Fatalf("%s:%d: tAssertNotBetween failed, min = %v, max = %v, val = %v", file, line, min, max, val)
		}
	}
}

func tAssertMatch(t testing.TB, expectedPattern string, got []byte, args ...interface{}) {
	if matched, err := regexp.Match(expectedPattern, got); err != nil || !matched {
		file, line := tCallerFileLine(1)
		if err != nil {
			if msg := fmt.Sprint(args...); msg != "" {
				t.Fatalf("%s:%d: tAssertMatch failed, expected = %q, got = %v, err = %v, %s", file, line, expectedPattern, got, err, msg)
			} else {
				t.Fatalf("%s:%d: tAssertMatch failed, expected = %q, got = %v, err = %v", file, line, expectedPattern, got, err)
			}
		} else {
			if msg := fmt.Sprint(args...); msg != "" {
				t.Fatalf("%s:%d: tAssertMatch failed, expected = %q, got = %v, %s", file, line, expectedPattern, got, msg)
			} else {
				t.Fatalf("%s:%d: tAssertMatch failed, expected = %q, got = %v", file, line, expectedPattern, got)
			}
		}
	}
}

func tAssertMatchString(t testing.TB, expectedPattern, got string, args ...interface{}) {
	if matched, err := regexp.MatchString(expectedPattern, got); err != nil || !matched {
		file, line := tCallerFileLine(1)
		if err != nil {
			if msg := fmt.Sprint(args...); msg != "" {
				t.Fatalf("%s:%d: tAssertMatchString failed, expected = %q, got = %v, err = %v, %s", file, line, expectedPattern, got, err, msg)
			} else {
				t.Fatalf("%s:%d: tAssertMatchString failed, expected = %q, got = %v, err = %v", file, line, expectedPattern, got, err)
			}
		} else {
			if msg := fmt.Sprint(args...); msg != "" {
				t.Fatalf("%s:%d: tAssertMatchString failed, expected = %q, got = %v, %s", file, line, expectedPattern, got, msg)
			} else {
				t.Fatalf("%s:%d: tAssertMatchString failed, expected = %q, got = %v", file, line, expectedPattern, got)
			}
		}
	}
}

func tAssertSliceContain(t testing.TB, slice, val interface{}, args ...interface{}) {
	sliceVal := reflect.ValueOf(slice)
	if sliceVal.Kind() != reflect.Slice {
		panic(fmt.Sprintf("tAssertSliceContain called with non-slice value of type %T", slice))
	}
	var contained bool
	for i := 0; i < sliceVal.Len(); i++ {
		if reflect.DeepEqual(sliceVal.Index(i).Interface(), val) {
			contained = true
			break
		}
	}
	if !contained {
		file, line := tCallerFileLine(1)
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("%s:%d: tAssertSliceContain failed, slice = %v, val = %v, %s", file, line, slice, val, msg)
		} else {
			t.Fatalf("%s:%d: tAssertSliceContain failed, slice = %v, val = %v", file, line, slice, val)
		}
	}
}

func tAssertSliceNotContain(t testing.TB, slice, val interface{}, args ...interface{}) {
	sliceVal := reflect.ValueOf(slice)
	if sliceVal.Kind() != reflect.Slice {
		panic(fmt.Sprintf("tAssertSliceNotContain called with non-slice value of type %T", slice))
	}
	var contained bool
	for i := 0; i < sliceVal.Len(); i++ {
		if reflect.DeepEqual(sliceVal.Index(i).Interface(), val) {
			contained = true
			break
		}
	}
	if contained {
		file, line := tCallerFileLine(1)
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("%s:%d: tAssertSliceNotContain failed, slice = %v, val = %v, %s", file, line, slice, val, msg)
		} else {
			t.Fatalf("%s:%d: tAssertSliceNotContain failed, slice = %v, val = %v", file, line, slice, val)
		}
	}
}

func tAssertMapContain(t testing.TB, m, key, val interface{}, args ...interface{}) {
	mapVal := reflect.ValueOf(m)
	if mapVal.Kind() != reflect.Map {
		panic(fmt.Sprintf("tAssertMapContain called with non-map value of type %T", m))
	}
	elemVal := mapVal.MapIndex(reflect.ValueOf(key))
	if !elemVal.IsValid() || !reflect.DeepEqual(elemVal.Interface(), val) {
		file, line := tCallerFileLine(1)
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("%s:%d: tAssertMapContain failed, map = %v, key = %v, val = %v, %s", file, line, m, key, val, msg)
		} else {
			t.Fatalf("%s:%d: tAssertMapContain failed, map = %v, key = %v, val = %v", file, line, m, key, val)
		}
	}
}

func tAssertMapContainKey(t testing.TB, m, key interface{}, args ...interface{}) {
	mapVal := reflect.ValueOf(m)
	if mapVal.Kind() != reflect.Map {
		panic(fmt.Sprintf("tAssertMapContainKey called with non-map value of type %T", m))
	}
	elemVal := mapVal.MapIndex(reflect.ValueOf(key))
	if !elemVal.IsValid() {
		file, line := tCallerFileLine(1)
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("%s:%d: tAssertMapContainKey failed, map = %v, key = %v, %s", file, line, m, key, msg)
		} else {
			t.Fatalf("%s:%d: tAssertMapContainKey failed, map = %v, key = %v", file, line, m, key)
		}
	}
}

func tAssertMapContainVal(t testing.TB, m, val interface{}, args ...interface{}) {
	mapVal := reflect.ValueOf(m)
	if mapVal.Kind() != reflect.Map {
		panic(fmt.Sprintf("tAssertMapContainVal called with non-map value of type %T", m))
	}
	var contained bool
	for _, key := range mapVal.MapKeys() {
		elemVal := mapVal.MapIndex(key)
		if elemVal.IsValid() && reflect.DeepEqual(elemVal.Interface(), val) {
			contained = true
			break
		}
	}
	if !contained {
		file, line := tCallerFileLine(1)
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("%s:%d: tAssertMapContainVal failed, map = %v, val = %v, %s", file, line, m, val, msg)
		} else {
			t.Fatalf("%s:%d: tAssertMapContainVal failed, map = %v, val = %v", file, line, m, val)
		}
	}
}

func tAssertMapNotContain(t testing.TB, m, key, val interface{}, args ...interface{}) {
	mapVal := reflect.ValueOf(m)
	if mapVal.Kind() != reflect.Map {
		panic(fmt.Sprintf("tAssertMapNotContain called with non-map value of type %T", m))
	}
	elemVal := mapVal.MapIndex(reflect.ValueOf(key))
	if elemVal.IsValid() && reflect.DeepEqual(elemVal.Interface(), val) {
		file, line := tCallerFileLine(1)
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("%s:%d: tAssertMapNotContain failed, map = %v, key = %v, val = %v, %s", file, line, m, key, val, msg)
		} else {
			t.Fatalf("%s:%d: tAssertMapNotContain failed, map = %v, key = %v, val = %v", file, line, m, key, val)
		}
	}
}

func tAssertMapNotContainKey(t testing.TB, m, key interface{}, args ...interface{}) {
	mapVal := reflect.ValueOf(m)
	if mapVal.Kind() != reflect.Map {
		panic(fmt.Sprintf("tAssertMapNotContainKey called with non-map value of type %T", m))
	}
	elemVal := mapVal.MapIndex(reflect.ValueOf(key))
	if elemVal.IsValid() {
		file, line := tCallerFileLine(1)
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("%s:%d: tAssertMapNotContainKey failed, map = %v, key = %v, %s", file, line, m, key, msg)
		} else {
			t.Fatalf("%s:%d: tAssertMapNotContainKey failed, map = %v, key = %v", file, line, m, key)
		}
	}
}

func tAssertMapNotContainVal(t testing.TB, m, val interface{}, args ...interface{}) {
	mapVal := reflect.ValueOf(m)
	if mapVal.Kind() != reflect.Map {
		panic(fmt.Sprintf("tAssertMapNotContainVal called with non-map value of type %T", m))
	}
	var contained bool
	for _, key := range mapVal.MapKeys() {
		elemVal := mapVal.MapIndex(key)
		if elemVal.IsValid() && reflect.DeepEqual(elemVal.Interface(), val) {
			contained = true
			break
		}
	}
	if contained {
		file, line := tCallerFileLine(1)
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("%s:%d: tAssertMapNotContainVal failed, map = %v, val = %v, %s", file, line, m, val, msg)
		} else {
			t.Fatalf("%s:%d: tAssertMapNotContainVal failed, map = %v, val = %v", file, line, m, val)
		}
	}
}

func tAssertZero(t testing.TB, val interface{}, args ...interface{}) {
	if !reflect.DeepEqual(reflect.Zero(reflect.TypeOf(val)).Interface(), val) {
		file, line := tCallerFileLine(1)
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("%s:%d: tAssertZero failed, val = %v, %s", file, line, val, msg)
		} else {
			t.Fatalf("%s:%d: tAssertZero failed, val = %v", file, line, val)
		}
	}
}

func tAssertNotZero(t testing.TB, val interface{}, args ...interface{}) {
	if reflect.DeepEqual(reflect.Zero(reflect.TypeOf(val)).Interface(), val) {
		file, line := tCallerFileLine(1)
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("%s:%d: tAssertNotZero failed, val = %v, %s", file, line, val, msg)
		} else {
			t.Fatalf("%s:%d: tAssertNotZero failed, val = %v", file, line, val)
		}
	}
}

func tAssertFileExists(t testing.TB, path string, args ...interface{}) {
	if _, err := os.Stat(path); err != nil {
		file, line := tCallerFileLine(1)
		if msg := fmt.Sprint(args...); msg != "" {
			if err != nil {
				t.Fatalf("%s:%d: tAssertFileExists failed, path = %v, err = %v, %s", file, line, path, err, msg)
			} else {
				t.Fatalf("%s:%d: tAssertFileExists failed, path = %v, %s", file, line, path, msg)
			}
		} else {
			if err != nil {
				t.Fatalf("%s:%d: tAssertFileExists failed, path = %v, err = %v", file, line, path, err)
			} else {
				t.Fatalf("%s:%d: tAssertFileExists failed, path = %v", file, line, path)
			}
		}
	}
}

func tAssertFileNotExists(t testing.TB, path string, args ...interface{}) {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		file, line := tCallerFileLine(1)
		if msg := fmt.Sprint(args...); msg != "" {
			if err != nil {
				t.Fatalf("%s:%d: tAssertFileNotExists failed, path = %v, err = %v, %s", file, line, path, err, msg)
			} else {
				t.Fatalf("%s:%d: tAssertFileNotExists failed, path = %v, %s", file, line, path, msg)
			}
		} else {
			if err != nil {
				t.Fatalf("%s:%d: tAssertFileNotExists failed, path = %v, err = %v", file, line, path, err)
			} else {
				t.Fatalf("%s:%d: tAssertFileNotExists failed, path = %v", file, line, path)
			}
		}
	}
}

func tAssertImplements(t testing.TB, interfaceObj, obj interface{}, args ...interface{}) {
	if !reflect.TypeOf(obj).Implements(reflect.TypeOf(interfaceObj).Elem()) {
		file, line := tCallerFileLine(1)
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("%s:%d: tAssertImplements failed, interface = %T, obj = %T, %s", file, line, interfaceObj, obj, msg)
		} else {
			t.Fatalf("%s:%d: tAssertImplements failed, interface = %T, obj = %T", file, line, interfaceObj, obj)
		}
	}
}

func tAssertSameType(t testing.TB, expectedType interface{}, obj interface{}, args ...interface{}) {
	if !reflect.DeepEqual(reflect.TypeOf(obj), reflect.TypeOf(expectedType)) {
		file, line := tCallerFileLine(1)
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("%s:%d: tAssertSameType failed, expected = %T, obj = %T, %s", file, line, expectedType, obj, msg)
		} else {
			t.Fatalf("%s:%d: tAssertSameType failed, expected = %T, obj = %T", file, line, expectedType, obj)
		}
	}
}

func tAssertPanic(t testing.TB, f func(), args ...interface{}) {
	panicVal := func() (panicVal interface{}) {
		defer func() {
			panicVal = recover()
		}()
		f()
		return
	}()

	if panicVal == nil {
		file, line := tCallerFileLine(1)
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("%s:%d: tAssertPanic failed, %s", file, line, msg)
		} else {
			t.Fatalf("%s:%d: tAssertPanic failed", file, line)
		}
	}
}

func tAssertNotPanic(t testing.TB, f func(), args ...interface{}) {
	panicVal := func() (panicVal interface{}) {
		defer func() {
			panicVal = recover()
		}()
		f()
		return
	}()

	if panicVal != nil {
		file, line := tCallerFileLine(1)
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("%s:%d: tAssertNotPanic failed, panic = %v, %s", file, line, panicVal, msg)
		} else {
			t.Fatalf("%s:%d: tAssertNotPanic failed, panic = %v", file, line, panicVal)
		}
	}
}

func tAssertEQ(t testing.TB, expected, got interface{}, args ...interface{}) {
	if !reflect.DeepEqual(expected, got) && !tIsNumberEqual(expected, got) {
		file, line := tCallerFileLine(1)
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("%s:%d: tAssertEQ failed, expected = %v, got = %v, %s", file, line, expected, got, msg)
		} else {
			t.Fatalf("%s:%d: tAssertEQ failed, expected = %v, got = %v", file, line, expected, got)
		}
	}
}

func tAssertNE(t testing.TB, expected, got interface{}, args ...interface{}) {
	if reflect.DeepEqual(expected, got) || tIsNumberEqual(expected, got) {
		file, line := tCallerFileLine(1)
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("%s:%d: tAssertNE failed, expected = %v, got = %v, %s", file, line, expected, got, msg)
		} else {
			t.Fatalf("%s:%d: tAssertNE failed, expected = %v, got = %v", file, line, expected, got)
		}
	}
}

func tAssertLE(t testing.TB, a, b int, args ...interface{}) {
	if !(a <= b) {
		file, line := tCallerFileLine(1)
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("%s:%d: tAssertLE failed, expected %v <= %v, %s", file, line, a, b, msg)
		} else {
			t.Fatalf("%s:%d: tAssertLE failed, expected %v <= %v", file, line, a, b)
		}
	}
}
