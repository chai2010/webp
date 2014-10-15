// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package assert_test

import (
	"bytes"
	"fmt"
	"image"
	"math"
	"strings"
	"testing"

	. "github.com/chai2010/assert"
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
