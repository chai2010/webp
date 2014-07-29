// Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#ifndef TEST_UTIL_H_
#define TEST_UTIL_H_

#include <math.h>
#include <string>

// int arr[5][3];
// assert(TEST_DIM(arr) == 5);
// assert(TEST_DIM(arr[0]) == 3);
#ifndef TEST_DIM
#define TEST_DIM(x) ((sizeof(x))/(sizeof((x)[0])))
#endif

bool DecodeJpeg(
	std::string* dst, const char* data, int size,
	int* width, int* height, int* channels
);

bool EncodeJpeg(
	std::string* dst, const char* data, int size,
	int width, int height, int channels, int quality /* =90 */,
	int width_step /* =0 */
);

bool DecodePng32(
	std::string* dst, const char* data, int size,
	int* width, int* height
);

bool DecodePng24(
	std::string* dst, const char* data, int size,
	int* width, int* height
);

bool EncodePng32(
	std::string* dst, const char* data, int size,
	int width, int height, int width_step /*=0*/
);

bool EncodePng24(
	std::string* dst, const char* data, int size,
	int width, int height, int width_step /*=0*/
);

bool LoadFileData(const char* name, std::string* data);
bool SaveFileData(const char* name, const char* data, int size);

template<typename T> double DiffImageData(
	const T* b0, const T* b1,
	int width, int height, int channels
) {
	double sum = 0;
	int n = width*height*channels;
	for(int i = 0; i < n; ++i) {
		sum += abs(int(b0[i])-int(b1[i]));
	}
	return sum/n;
}

#endif // TEST_UTIL_H_

