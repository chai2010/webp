// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// go test -assert.failed

package assert

import (
	"bytes"
	"flag"
	"fmt"
	"image"
	"math"
	"strings"
	"testing"
)

var (
	flagAssertFailedTest = flag.Bool("assert.failed", false, "run assert failed test")
)

func TestAssert_failed_01(t *testing.T) {
	if !*flagAssertFailedTest {
		t.SkipNow()
	}
	Assert(t, 1 == 2)
}

func TestAssert_failed_02(t *testing.T) {
	if !*flagAssertFailedTest {
		t.SkipNow()
	}
	Assert(t, 1 == 2, "message1", "message2")
}

func TestAssertNil_failed(t *testing.T) {
	if !*flagAssertFailedTest {
		t.SkipNow()
	}
	AssertNil(t, fmt.Errorf("error"))
}

func TestAssertNotNil_failed(t *testing.T) {
	if !*flagAssertFailedTest {
		t.SkipNow()
	}
	AssertNotNil(t, nil)
}

func TestAssertTrue_failed(t *testing.T) {
	if !*flagAssertFailedTest {
		t.SkipNow()
	}
	AssertTrue(t, false)
}

func TestAssertFalse_failed(t *testing.T) {
	if !*flagAssertFailedTest {
		t.SkipNow()
	}
	AssertFalse(t, true)
}

func TestAssertEqual_failed_01(t *testing.T) {
	if !*flagAssertFailedTest {
		t.SkipNow()
	}
	AssertEqual(t, 1, 1+1)
}

func TestAssertEqual_failed_02(t *testing.T) {
	if !*flagAssertFailedTest {
		t.SkipNow()
	}
	AssertEqual(t, "abC", strings.ToLower("ABC"))
}

func TestAssertEqual_failed_03(t *testing.T) {
	if !*flagAssertFailedTest {
		t.SkipNow()
	}
	AssertEqual(t, []byte("abC"), bytes.ToLower([]byte("ABC")))
}

func TestAssertEqual_failed_04(t *testing.T) {
	if !*flagAssertFailedTest {
		t.SkipNow()
	}
	AssertEqual(t, image.Pt(1, 2), image.Pt(0, 0))
}

func TestAssertNotEqual_failed_01(t *testing.T) {
	if !*flagAssertFailedTest {
		t.SkipNow()
	}
	AssertNotEqual(t, 1, 1)
}

func TestAssertNotEqual_failed_02(t *testing.T) {
	if !*flagAssertFailedTest {
		t.SkipNow()
	}
	AssertNotEqual(t, "abc", strings.ToLower("ABC"))
}

func TestAssertNotEqual_failed_03(t *testing.T) {
	if !*flagAssertFailedTest {
		t.SkipNow()
	}
	AssertNotEqual(t, []byte("abc"), bytes.ToLower([]byte("ABC")))
}

func TestAssertNotEqual_failed_04(t *testing.T) {
	if !*flagAssertFailedTest {
		t.SkipNow()
	}
	AssertNotEqual(t, image.Pt(1, 2), image.Pt(1, 2))
}

func TestAssertNear_failed(t *testing.T) {
	if !*flagAssertFailedTest {
		t.SkipNow()
	}
	AssertNear(t, 1.414, math.Sqrt(2), 0.00001)
}

func TestAssertBetween_failed_01(t *testing.T) {
	if !*flagAssertFailedTest {
		t.SkipNow()
	}
	AssertBetween(t, 0, 255, -1)
}

func TestAssertBetween_failed_02(t *testing.T) {
	if !*flagAssertFailedTest {
		t.SkipNow()
	}
	AssertBetween(t, 0, 255, 256)
}

func TestAssertNotBetween_failed_01(t *testing.T) {
	if !*flagAssertFailedTest {
		t.SkipNow()
	}
	AssertNotBetween(t, 0, 255, 0)
}

func TestAssertNotBetween_failed_02(t *testing.T) {
	if !*flagAssertFailedTest {
		t.SkipNow()
	}
	AssertNotBetween(t, 0, 255, 128)
}

func TestAssertNotBetween_failed_03(t *testing.T) {
	if !*flagAssertFailedTest {
		t.SkipNow()
	}
	AssertNotBetween(t, 0, 255, 255)
}

func TestAssertMatch_failed(t *testing.T) {
	if !*flagAssertFailedTest {
		t.SkipNow()
	}
	AssertMatch(t, `\.go$`, []byte("assert.cc"))
}

func TestAssertMatchString_failed(t *testing.T) {
	if !*flagAssertFailedTest {
		t.SkipNow()
	}
	AssertMatchString(t, `\.go$`, "assert.cc")
}

func TestAssertSliceContain_failed_01(t *testing.T) {
	if !*flagAssertFailedTest {
		t.SkipNow()
	}
	AssertSliceContain(t, []int{1, 1, 2, 3, 5, 8, 13}, "8")
}

func TestAssertSliceContain_failed_02(t *testing.T) {
	if !*flagAssertFailedTest {
		t.SkipNow()
	}
	AssertSliceContain(t, []interface{}{1, 1, 2, 3, 5, "8", 13}, 8)
}

func TestAssertSliceNotContain_failed_01(t *testing.T) {
	if !*flagAssertFailedTest {
		t.SkipNow()
	}
	AssertSliceNotContain(t, []int{1, 1, 2, 3, 5, 8, 13}, 8)
}

func TestAssertSliceNotContain_failed_02(t *testing.T) {
	if !*flagAssertFailedTest {
		t.SkipNow()
	}
	AssertSliceNotContain(t, []interface{}{1, 1, 2, 3, 5, "8", 13}, "8")
}

func TestAssertMapContain_failed_01(t *testing.T) {
	if !*flagAssertFailedTest {
		t.SkipNow()
	}
	AssertMapContain(t,
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

func TestAssertMapContain_failed_02(t *testing.T) {
	if !*flagAssertFailedTest {
		t.SkipNow()
	}
	AssertMapContain(t,
		map[string]int{
			"UTC": 0 * 60 * 60,
			"EST": -5 * 60 * 60,
			"CST": -6 * 60 * 60,
			"MST": -7 * 60 * 60,
			"PST": -8 * 60 * 60,
		},
		"MST", 1984,
	)
}

func TestAssertMapContainKey_failed(t *testing.T) {
	if !*flagAssertFailedTest {
		t.SkipNow()
	}
	AssertMapContainKey(t,
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

func TestAssertMapContainVal_failed(t *testing.T) {
	if !*flagAssertFailedTest {
		t.SkipNow()
	}
	AssertMapContainVal(t,
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

func TestAssertMapNotContain_failed(t *testing.T) {
	if !*flagAssertFailedTest {
		t.SkipNow()
	}
	AssertMapNotContain(t,
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

func TestAssertMapNotContainKey_failed(t *testing.T) {
	if !*flagAssertFailedTest {
		t.SkipNow()
	}
	AssertMapNotContainKey(t,
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

func TestAssertMapNotContainVal_failed(t *testing.T) {
	if !*flagAssertFailedTest {
		t.SkipNow()
	}
	AssertMapNotContainVal(t,
		map[string]int{
			"UTC": 0 * 60 * 60,
			"EST": -5 * 60 * 60,
			"CST": -6 * 60 * 60,
			"MST": -7 * 60 * 60,
			"PST": -8 * 60 * 60,
		},
		-8*60*60,
	)
}

func TestAssertZero_failed(t *testing.T) {
	if !*flagAssertFailedTest {
		t.SkipNow()
	}
	AssertZero(t, struct {
		A bool
		B string
		C int
		d map[string]interface{}
	}{A: true})
}

func TestAssertNotZero_failed(t *testing.T) {
	if !*flagAssertFailedTest {
		t.SkipNow()
	}
	AssertNotZero(t, struct {
		A bool
		B string
		C int
		d map[string]interface{}
	}{})
}

func TestAssertFileExists_failed(t *testing.T) {
	if !*flagAssertFailedTest {
		t.SkipNow()
	}
	AssertFileExists(t, "assert.cc")
}

func TestAssertFileNotExists_failed(t *testing.T) {
	if !*flagAssertFailedTest {
		t.SkipNow()
	}
	AssertFileNotExists(t, "assert.go")
}

func TestAssertImplements_failed(t *testing.T) {
	if !*flagAssertFailedTest {
		t.SkipNow()
	}
	AssertImplements(t, (*error)(nil), "NotErrorType")
}

func TestAssertSameType_failed(t *testing.T) {
	if !*flagAssertFailedTest {
		t.SkipNow()
	}
	AssertSameType(t, string("abc"), []byte("ABC"))
}

func TestAssertPanic_failed(t *testing.T) {
	if !*flagAssertFailedTest {
		t.SkipNow()
	}
	AssertPanic(t, func() {})
}

func TestAssertNotPanic_failed(t *testing.T) {
	if !*flagAssertFailedTest {
		t.SkipNow()
	}
	AssertNotPanic(t, func() { panic("TestAssertNotPanic_failed") })
}
