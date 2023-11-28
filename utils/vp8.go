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

package utils

//noinspection GoUnusedConst
const (
	Vp8StatusOk VP8StatusCode = iota
	Vp8StatusOutOfMemory
	Vp8StatusInvalidParam
	Vp8StatusBitstreamError
	Vp8StatusUnsupportedFeature
	Vp8StatusSuspended
	Vp8StatusUserAbort
	Vp8StatusNotEnoughData
)

type VP8StatusCode int

func (c VP8StatusCode) String() (label string) {
	switch c {
	case Vp8StatusOk:
		label = "VP8_STATUS_OK"
	case Vp8StatusOutOfMemory:
		label = "VP8_STATUS_OUT_OF_MEMORY"
	case Vp8StatusInvalidParam:
		label = "VP8_STATUS_INVALID_PARAM"
	case Vp8StatusBitstreamError:
		label = "VP8_STATUS_BITSTREAM_ERROR"
	case Vp8StatusUnsupportedFeature:
		label = "VP8_STATUS_UNSUPPORTED_FEATURE"
	case Vp8StatusSuspended:
		label = "VP8_STATUS_SUSPENDED"
	case Vp8StatusUserAbort:
		label = "VP8_STATUS_USER_ABORT"
	case Vp8StatusNotEnoughData:
		label = "VP8_STATUS_NOT_ENOUGH_DATA"
	default:
		label = "VP8 undefined status code"
	}

	return
}

const (
	Vp8EncOk Vp8EncStatus = iota
	Vp8EncErrorOutOfMemory
	Vp8EncErrorBitstreamOutOfMemory
	Vp8EncErrorNullParameter
	Vp8EncErrorInvalidConfiguration
	Vp8EncErrorBadDimension
	Vp8EncErrorPartition0Overflow
	Vp8EncErrorPartitionOverflow
	Vp8EncErrorBadWrite
	Vp8EncErrorFileTooBig
	Vp8EncErrorUserAbort
	Vp8EncErrorLast
)

type Vp8EncStatus int

func (c Vp8EncStatus) String() (label string) {
	switch c {
	case Vp8EncOk:
		label = "VP8_ENC_OK"
	case Vp8EncErrorOutOfMemory:
		label = "VP8_ENC_ERROR_OUT_OF_MEMORY"
	case Vp8EncErrorBitstreamOutOfMemory:
		label = "VP8_ENC_ERROR_BITSTREAM_OUT_OF_MEMORY"
	case Vp8EncErrorNullParameter:
		label = "VP8_ENC_ERROR_NULL_PARAMETER"
	case Vp8EncErrorInvalidConfiguration:
		label = "VP8_ENC_ERROR_INVALID_CONFIGURATION"
	case Vp8EncErrorBadDimension:
		label = "VP8_ENC_ERROR_BAD_DIMENSION"
	case Vp8EncErrorPartition0Overflow:
		label = "VP8_ENC_ERROR_PARTITION0_OVERFLOW"
	case Vp8EncErrorPartitionOverflow:
		label = "VP8_ENC_ERROR_PARTITION_OVERFLOW"
	case Vp8EncErrorBadWrite:
		label = "VP8_ENC_ERROR_BAD_WRITE"
	case Vp8EncErrorFileTooBig:
		label = "VP8_ENC_ERROR_FILE_TOO_BIG"
	case Vp8EncErrorUserAbort:
		label = "VP8_ENC_ERROR_USER_ABORT"
	case Vp8EncErrorLast:
		label = "VP8_ENC_ERROR_LAST"
	default:
		label = "VP8 undefined status code"
	}

	return
}
