// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package assert provides assert helper functions for testing package.

See failed test:

	go test -assert.failed

Example:

	package assert_test

	import (
		"bytes"
		"image"
		"math"
		"strings"
		"testing"

		. "github.com/chai2010/webp/internal/assert"
	)

	func TestAssert(t *testing.T) {
		Assert(t, 1 == 1)
		Assert(t, 1 == 1, "message1", "message2")
	}

	func TestAssertNil(t *testing.T) {
		AssertNil(t, nil)
	}

	func TestAssertNotNil(t *testing.T) {
		AssertNotNil(t, fmt.Errorf("error"))
	}

	func TestAssertTrue(t *testing.T) {
		AssertTrue(t, true)
	}

	func TestAssertFalse(t *testing.T) {
		AssertFalse(t, false)
	}

	func TestAssertEqual(t *testing.T) {
		AssertEqual(t, 2, 1+1)
		AssertEqual(t, "abc", strings.ToLower("ABC"))
		AssertEqual(t, []byte("abc"), bytes.ToLower([]byte("ABC")))
		AssertEqual(t, image.Pt(1, 2), image.Pt(1, 2))
	}

	func TestAssertNotEqual(t *testing.T) {
		AssertNotEqual(t, 2, 1)
		AssertNotEqual(t, "ABC", strings.ToLower("ABC"))
		AssertNotEqual(t, []byte("ABC"), bytes.ToLower([]byte("ABC")))
		AssertNotEqual(t, image.Pt(1, 2), image.Pt(2, 2))
		AssertNotEqual(t, image.Pt(1, 2), image.Rect(1, 2, 3, 4))
	}

	func TestAssertNear(t *testing.T) {
		AssertNear(t, 1.414, math.Sqrt(2), 0.1)
	}

	func TestAssertBetween(t *testing.T) {
		AssertBetween(t, 0, 255, 0)
		AssertBetween(t, 0, 255, 128)
		AssertBetween(t, 0, 255, 255)
	}

	func TestAssertNotBetween(t *testing.T) {
		AssertNotBetween(t, 0, 255, -1)
		AssertNotBetween(t, 0, 255, 256)
	}

	func TestAssertMatch(t *testing.T) {
		AssertMatch(t, `^\w+@\w+\.com$`, []byte("chaishushan@gmail.com"))
		AssertMatch(t, `^assert`, []byte("assert.go"))
		AssertMatch(t, `\.go$`, []byte("assert.go"))
	}

	func TestAssertMatchString(t *testing.T) {
		AssertMatchString(t, `^\w+@\w+\.com$`, "chaishushan@gmail.com")
		AssertMatchString(t, `^assert`, "assert.go")
		AssertMatchString(t, `\.go$`, "assert.go")
	}

	func TestAssertSliceContain(t *testing.T) {
		AssertSliceContain(t, []int{1, 1, 2, 3, 5, 8, 13}, 8)
		AssertSliceContain(t, []interface{}{1, 1, 2, 3, 5, "8", 13}, "8")
	}

	func TestAssertSliceNotContain(t *testing.T) {
		AssertSliceNotContain(t, []int{1, 1, 2, 3, 5, 8, 13}, 12)
		AssertSliceNotContain(t, []interface{}{1, 1, 2, 3, 5, "8", 13}, 8)
	}

	func TestAssertMapContain(t *testing.T) {
		AssertMapContain(t,
			map[string]int{
				"UTC": 0 * 60 * 60,
				"EST": -5 * 60 * 60,
				"CST": -6 * 60 * 60,
				"MST": -7 * 60 * 60,
				"PST": -8 * 60 * 60,
			},
			"MST", -7*60*60,
		)
	}

	func TestAssertMapContainKey(t *testing.T) {
		AssertMapContainKey(t,
			map[string]int{
				"UTC": 0 * 60 * 60,
				"EST": -5 * 60 * 60,
				"CST": -6 * 60 * 60,
				"MST": -7 * 60 * 60,
				"PST": -8 * 60 * 60,
			},
			"MST",
		)
	}

	func TestAssertMapContainVal(t *testing.T) {
		AssertMapContainVal(t,
			map[string]int{
				"UTC": 0 * 60 * 60,
				"EST": -5 * 60 * 60,
				"CST": -6 * 60 * 60,
				"MST": -7 * 60 * 60,
				"PST": -8 * 60 * 60,
			},
			-7*60*60,
		)
	}

	func TestAssertMapNotContain(t *testing.T) {
		AssertMapNotContain(t,
			map[string]int{
				"UTC": 0 * 60 * 60,
				"EST": -5 * 60 * 60,
				"CST": -6 * 60 * 60,
				"MST": -7 * 60 * 60,
				"PST": -8 * 60 * 60,
			},
			"ABC", -7*60*60,
		)
	}

	func TestAssertMapNotContainKey(t *testing.T) {
		AssertMapNotContainKey(t,
			map[string]int{
				"UTC": 0 * 60 * 60,
				"EST": -5 * 60 * 60,
				"CST": -6 * 60 * 60,
				"MST": -7 * 60 * 60,
				"PST": -8 * 60 * 60,
			},
			"ABC",
		)
	}

	func TestAssertMapNotContainVal(t *testing.T) {
		AssertMapNotContainVal(t,
			map[string]int{
				"UTC": 0 * 60 * 60,
				"EST": -5 * 60 * 60,
				"CST": -6 * 60 * 60,
				"MST": -7 * 60 * 60,
				"PST": -8 * 60 * 60,
			},
			1984,
		)
	}

	func TestAssertZero(t *testing.T) {
		AssertZero(t, struct {
			A bool
			B string
			C int
			d map[string]interface{}
		}{})
	}

	func TestAssertNotZero(t *testing.T) {
		AssertNotZero(t, struct {
			A bool
			B string
			C int
			d map[string]interface{}
		}{A: true})
	}

	func TestAssertFileExists(t *testing.T) {
		AssertFileExists(t, "assert.go")
	}

	func TestAssertFileNotExists(t *testing.T) {
		AssertFileNotExists(t, "assert.cc")
	}

	func TestAssertImplements(t *testing.T) {
		AssertImplements(t, (*error)(nil), fmt.Errorf("ErrorType"))
	}

	func TestAssertSameType(t *testing.T) {
		AssertSameType(t, string("abc"), string("ABC"))
	}

	func TestAssertPanic(t *testing.T) {
		AssertPanic(t, func() { panic("TestAssertPanic") })
	}

	func TestAssertNotPanic(t *testing.T) {
		AssertNotPanic(t, func() {})
	}

Report bugs to <chaishushan@gmail.com>.

Thanks!
*/
package assert

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

func callerFileLine() (file string, line int) {
	_, file, line, ok := runtime.Caller(2)
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

func Assert(t testing.TB, condition bool, args ...interface{}) {
	if !condition {
		file, line := callerFileLine()
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("%s:%d: Assert failed, %s", file, line, msg)
		} else {
			t.Fatalf("%s:%d: Assert failed", file, line)
		}
	}
}

func AssertNil(t testing.TB, p interface{}, args ...interface{}) {
	if p != nil {
		file, line := callerFileLine()
		if msg := fmt.Sprint(args...); msg != "" {
			if err, ok := p.(error); ok && err != nil {
				t.Fatalf("%s:%d: AssertNil failed, err = %v, %s", file, line, err, msg)
			} else {
				t.Fatalf("%s:%d: AssertNil failed, %s", file, line, msg)
			}
		} else {
			if err, ok := p.(error); ok && err != nil {
				t.Fatalf("%s:%d: AssertNil failed, err = %v", file, line, err)
			} else {
				t.Fatalf("%s:%d: AssertNil failed", file, line)
			}
		}
	}
}

func AssertNotNil(t testing.TB, p interface{}, args ...interface{}) {
	if p == nil {
		file, line := callerFileLine()
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("%s:%d: AssertNotNil failed, %s", file, line, msg)
		} else {
			t.Fatalf("%s:%d: AssertNotNil failed", file, line)
		}
	}
}

func AssertTrue(t testing.TB, condition bool, args ...interface{}) {
	if condition != true {
		file, line := callerFileLine()
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("%s:%d: AssertTrue failed, %s", file, line, msg)
		} else {
			t.Fatalf("%s:%d: AssertTrue failed", file, line)
		}
	}
}

func AssertFalse(t testing.TB, condition bool, args ...interface{}) {
	if condition != false {
		file, line := callerFileLine()
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("%s:%d: AssertFalse failed, %s", file, line, msg)
		} else {
			t.Fatalf("%s:%d: AssertFalse failed", file, line)
		}
	}
}

func AssertEqual(t testing.TB, expected, got interface{}, args ...interface{}) {
	if !reflect.DeepEqual(expected, got) {
		file, line := callerFileLine()
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("%s:%d: AssertEqual failed, expected = %v, got = %v, %s", file, line, expected, got, msg)
		} else {
			t.Fatalf("%s:%d: AssertEqual failed, expected = %v, got = %v", file, line, expected, got)
		}
	}
}

func AssertNotEqual(t testing.TB, expected, got interface{}, args ...interface{}) {
	if reflect.DeepEqual(expected, got) {
		file, line := callerFileLine()
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("%s:%d: AssertNotEqual failed, expected = %v, got = %v, %s", file, line, expected, got, msg)
		} else {
			t.Fatalf("%s:%d: AssertNotEqual failed, expected = %v, got = %v", file, line, expected, got)
		}
	}
}

func AssertNear(t testing.TB, expected, got, abs float64, args ...interface{}) {
	if math.Abs(expected-got) > abs {
		file, line := callerFileLine()
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("%s:%d: AssertNear failed, expected = %v, got = %v, abs = %v, %s", file, line, expected, got, abs, msg)
		} else {
			t.Fatalf("%s:%d: AssertNear failed, expected = %v, got = %v, abs = %v", file, line, expected, got, abs)
		}
	}
}

func AssertBetween(t testing.TB, min, max, val float64, args ...interface{}) {
	if val < min || max < val {
		file, line := callerFileLine()
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("%s:%d: AssertBetween failed, min = %v, max = %v, val = %v, %s", file, line, min, max, val, msg)
		} else {
			t.Fatalf("%s:%d: AssertBetween failed, min = %v, max = %v, val = %v", file, line, min, max, val)
		}
	}
}

func AssertNotBetween(t testing.TB, min, max, val float64, args ...interface{}) {
	if min <= val && val <= max {
		file, line := callerFileLine()
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("%s:%d: AssertNotBetween failed, min = %v, max = %v, val = %v, %s", file, line, min, max, val, msg)
		} else {
			t.Fatalf("%s:%d: AssertNotBetween failed, min = %v, max = %v, val = %v", file, line, min, max, val)
		}
	}
}

func AssertMatch(t testing.TB, expectedPattern string, got []byte, args ...interface{}) {
	if matched, err := regexp.Match(expectedPattern, got); err != nil || !matched {
		file, line := callerFileLine()
		if err != nil {
			if msg := fmt.Sprint(args...); msg != "" {
				t.Fatalf("%s:%d: AssertMatch failed, expected = %q, got = %v, err = %v, %s", file, line, expectedPattern, got, err, msg)
			} else {
				t.Fatalf("%s:%d: AssertMatch failed, expected = %q, got = %v, err = %v", file, line, expectedPattern, got, err)
			}
		} else {
			if msg := fmt.Sprint(args...); msg != "" {
				t.Fatalf("%s:%d: AssertMatch failed, expected = %q, got = %v, %s", file, line, expectedPattern, got, msg)
			} else {
				t.Fatalf("%s:%d: AssertMatch failed, expected = %q, got = %v", file, line, expectedPattern, got)
			}
		}
	}
}

func AssertMatchString(t testing.TB, expectedPattern, got string, args ...interface{}) {
	if matched, err := regexp.MatchString(expectedPattern, got); err != nil || !matched {
		file, line := callerFileLine()
		if err != nil {
			if msg := fmt.Sprint(args...); msg != "" {
				t.Fatalf("%s:%d: AssertMatchString failed, expected = %q, got = %v, err = %v, %s", file, line, expectedPattern, got, err, msg)
			} else {
				t.Fatalf("%s:%d: AssertMatchString failed, expected = %q, got = %v, err = %v", file, line, expectedPattern, got, err)
			}
		} else {
			if msg := fmt.Sprint(args...); msg != "" {
				t.Fatalf("%s:%d: AssertMatchString failed, expected = %q, got = %v, %s", file, line, expectedPattern, got, msg)
			} else {
				t.Fatalf("%s:%d: AssertMatchString failed, expected = %q, got = %v", file, line, expectedPattern, got)
			}
		}
	}
}

func AssertSliceContain(t testing.TB, slice, val interface{}, args ...interface{}) {
	sliceVal := reflect.ValueOf(slice)
	if sliceVal.Kind() != reflect.Slice {
		panic(fmt.Sprintf("AssertSliceContain called with non-slice value of type %T", slice))
	}
	var contained bool
	for i := 0; i < sliceVal.Len(); i++ {
		if reflect.DeepEqual(sliceVal.Index(i).Interface(), val) {
			contained = true
			break
		}
	}
	if !contained {
		file, line := callerFileLine()
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("%s:%d: AssertSliceContain failed, slice = %v, val = %v, %s", file, line, slice, val, msg)
		} else {
			t.Fatalf("%s:%d: AssertSliceContain failed, slice = %v, val = %v", file, line, slice, val)
		}
	}
}

func AssertSliceNotContain(t testing.TB, slice, val interface{}, args ...interface{}) {
	sliceVal := reflect.ValueOf(slice)
	if sliceVal.Kind() != reflect.Slice {
		panic(fmt.Sprintf("AssertSliceNotContain called with non-slice value of type %T", slice))
	}
	var contained bool
	for i := 0; i < sliceVal.Len(); i++ {
		if reflect.DeepEqual(sliceVal.Index(i).Interface(), val) {
			contained = true
			break
		}
	}
	if contained {
		file, line := callerFileLine()
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("%s:%d: AssertSliceNotContain failed, slice = %v, val = %v, %s", file, line, slice, val, msg)
		} else {
			t.Fatalf("%s:%d: AssertSliceNotContain failed, slice = %v, val = %v", file, line, slice, val)
		}
	}
}

func AssertMapContain(t testing.TB, m, key, val interface{}, args ...interface{}) {
	mapVal := reflect.ValueOf(m)
	if mapVal.Kind() != reflect.Map {
		panic(fmt.Sprintf("AssertMapContain called with non-map value of type %T", m))
	}
	elemVal := mapVal.MapIndex(reflect.ValueOf(key))
	if !elemVal.IsValid() || !reflect.DeepEqual(elemVal.Interface(), val) {
		file, line := callerFileLine()
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("%s:%d: AssertMapContain failed, map = %v, key = %v, val = %v, %s", file, line, m, key, val, msg)
		} else {
			t.Fatalf("%s:%d: AssertMapContain failed, map = %v, key = %v, val = %v", file, line, m, key, val)
		}
	}
}

func AssertMapContainKey(t testing.TB, m, key interface{}, args ...interface{}) {
	mapVal := reflect.ValueOf(m)
	if mapVal.Kind() != reflect.Map {
		panic(fmt.Sprintf("AssertMapContainKey called with non-map value of type %T", m))
	}
	elemVal := mapVal.MapIndex(reflect.ValueOf(key))
	if !elemVal.IsValid() {
		file, line := callerFileLine()
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("%s:%d: AssertMapContainKey failed, map = %v, key = %v, %s", file, line, m, key, msg)
		} else {
			t.Fatalf("%s:%d: AssertMapContainKey failed, map = %v, key = %v", file, line, m, key)
		}
	}
}

func AssertMapContainVal(t testing.TB, m, val interface{}, args ...interface{}) {
	mapVal := reflect.ValueOf(m)
	if mapVal.Kind() != reflect.Map {
		panic(fmt.Sprintf("AssertMapContainVal called with non-map value of type %T", m))
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
		file, line := callerFileLine()
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("%s:%d: AssertMapContainVal failed, map = %v, val = %v, %s", file, line, m, val, msg)
		} else {
			t.Fatalf("%s:%d: AssertMapContainVal failed, map = %v, val = %v", file, line, m, val)
		}
	}
}

func AssertMapNotContain(t testing.TB, m, key, val interface{}, args ...interface{}) {
	mapVal := reflect.ValueOf(m)
	if mapVal.Kind() != reflect.Map {
		panic(fmt.Sprintf("AssertMapNotContain called with non-map value of type %T", m))
	}
	elemVal := mapVal.MapIndex(reflect.ValueOf(key))
	if elemVal.IsValid() && reflect.DeepEqual(elemVal.Interface(), val) {
		file, line := callerFileLine()
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("%s:%d: AssertMapNotContain failed, map = %v, key = %v, val = %v, %s", file, line, m, key, val, msg)
		} else {
			t.Fatalf("%s:%d: AssertMapNotContain failed, map = %v, key = %v, val = %v", file, line, m, key, val)
		}
	}
}

func AssertMapNotContainKey(t testing.TB, m, key interface{}, args ...interface{}) {
	mapVal := reflect.ValueOf(m)
	if mapVal.Kind() != reflect.Map {
		panic(fmt.Sprintf("AssertMapNotContainKey called with non-map value of type %T", m))
	}
	elemVal := mapVal.MapIndex(reflect.ValueOf(key))
	if elemVal.IsValid() {
		file, line := callerFileLine()
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("%s:%d: AssertMapNotContainKey failed, map = %v, key = %v, %s", file, line, m, key, msg)
		} else {
			t.Fatalf("%s:%d: AssertMapNotContainKey failed, map = %v, key = %v", file, line, m, key)
		}
	}
}

func AssertMapNotContainVal(t testing.TB, m, val interface{}, args ...interface{}) {
	mapVal := reflect.ValueOf(m)
	if mapVal.Kind() != reflect.Map {
		panic(fmt.Sprintf("AssertMapNotContainVal called with non-map value of type %T", m))
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
		file, line := callerFileLine()
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("%s:%d: AssertMapNotContainVal failed, map = %v, val = %v, %s", file, line, m, val, msg)
		} else {
			t.Fatalf("%s:%d: AssertMapNotContainVal failed, map = %v, val = %v", file, line, m, val)
		}
	}
}

func AssertZero(t testing.TB, val interface{}, args ...interface{}) {
	if !reflect.DeepEqual(reflect.Zero(reflect.TypeOf(val)).Interface(), val) {
		file, line := callerFileLine()
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("%s:%d: AssertZero failed, val = %v, %s", file, line, val, msg)
		} else {
			t.Fatalf("%s:%d: AssertZero failed, val = %v", file, line, val)
		}
	}
}

func AssertNotZero(t testing.TB, val interface{}, args ...interface{}) {
	if reflect.DeepEqual(reflect.Zero(reflect.TypeOf(val)).Interface(), val) {
		file, line := callerFileLine()
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("%s:%d: AssertNotZero failed, val = %v, %s", file, line, val, msg)
		} else {
			t.Fatalf("%s:%d: AssertNotZero failed, val = %v", file, line, val)
		}
	}
}

func AssertFileExists(t testing.TB, path string, args ...interface{}) {
	if _, err := os.Stat(path); err != nil {
		file, line := callerFileLine()
		if msg := fmt.Sprint(args...); msg != "" {
			if err != nil {
				t.Fatalf("%s:%d: AssertFileExists failed, path = %v, err = %v, %s", file, line, path, err, msg)
			} else {
				t.Fatalf("%s:%d: AssertFileExists failed, path = %v, %s", file, line, path, msg)
			}
		} else {
			if err != nil {
				t.Fatalf("%s:%d: AssertFileExists failed, path = %v, err = %v", file, line, path, err)
			} else {
				t.Fatalf("%s:%d: AssertFileExists failed, path = %v", file, line, path)
			}
		}
	}
}

func AssertFileNotExists(t testing.TB, path string, args ...interface{}) {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		file, line := callerFileLine()
		if msg := fmt.Sprint(args...); msg != "" {
			if err != nil {
				t.Fatalf("%s:%d: AssertFileNotExists failed, path = %v, err = %v, %s", file, line, path, err, msg)
			} else {
				t.Fatalf("%s:%d: AssertFileNotExists failed, path = %v, %s", file, line, path, msg)
			}
		} else {
			if err != nil {
				t.Fatalf("%s:%d: AssertFileNotExists failed, path = %v, err = %v", file, line, path, err)
			} else {
				t.Fatalf("%s:%d: AssertFileNotExists failed, path = %v", file, line, path)
			}
		}
	}
}

func AssertImplements(t testing.TB, interfaceObj, obj interface{}, args ...interface{}) {
	if !reflect.TypeOf(obj).Implements(reflect.TypeOf(interfaceObj).Elem()) {
		file, line := callerFileLine()
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("%s:%d: AssertImplements failed, interface = %T, obj = %T, %s", file, line, interfaceObj, obj, msg)
		} else {
			t.Fatalf("%s:%d: AssertImplements failed, interface = %T, obj = %T", file, line, interfaceObj, obj)
		}
	}
}

func AssertSameType(t testing.TB, expectedType interface{}, obj interface{}, args ...interface{}) {
	if !reflect.DeepEqual(reflect.TypeOf(obj), reflect.TypeOf(expectedType)) {
		file, line := callerFileLine()
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("%s:%d: AssertSameType failed, expected = %T, obj = %T, %s", file, line, expectedType, obj, msg)
		} else {
			t.Fatalf("%s:%d: AssertSameType failed, expected = %T, obj = %T", file, line, expectedType, obj)
		}
	}
}

func AssertPanic(t testing.TB, f func(), args ...interface{}) {
	panicVal := func() (panicVal interface{}) {
		defer func() {
			panicVal = recover()
		}()
		f()
		return
	}()

	if panicVal == nil {
		file, line := callerFileLine()
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("%s:%d: AssertPanic failed, %s", file, line, msg)
		} else {
			t.Fatalf("%s:%d: AssertPanic failed", file, line)
		}
	}
}

func AssertNotPanic(t testing.TB, f func(), args ...interface{}) {
	panicVal := func() (panicVal interface{}) {
		defer func() {
			panicVal = recover()
		}()
		f()
		return
	}()

	if panicVal != nil {
		file, line := callerFileLine()
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("%s:%d: AssertNotPanic failed, panic = %v, %s", file, line, panicVal, msg)
		} else {
			t.Fatalf("%s:%d: AssertNotPanic failed, panic = %v", file, line, panicVal)
		}
	}
}
