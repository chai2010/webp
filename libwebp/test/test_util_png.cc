// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "test_util.h"

#include "./png/lodepng.h"
#include "./png/lodepng.cpp"

#include <string>

bool DecodePng32(
	std::string* dst, const char* data, int size,
	int* width, int* height
) {
	if(dst == NULL || data == NULL || size <= 0) {
		return false;
	}
	if(width == NULL || height == NULL) {
		return false;
	}

	unsigned char* img;
	unsigned w, h;

	auto err = lodepng_decode32(&img, &w, &h, (const unsigned char*)data, size);
	if(err != 0) return false;

	dst->assign((const char*)img, w*h*4);
	free(img);

	*width = w;
	*height = h;
	return true;
}

bool DecodePng24(
	std::string* dst, const char* data, int size,
	int* width, int* height
) {
	if(dst == NULL || data == NULL || size <= 0) {
		return false;
	}
	if(width == NULL || height == NULL) {
		return false;
	}

	unsigned char* img;
	unsigned w, h;

	auto err = lodepng_decode32(&img, &w, &h, (const unsigned char*)data, size);
	if(err != 0) return false;

	dst->assign((const char*)img, w*h*3);
	free(img);

	*width = w;
	*height = h;
	return true;
}

bool EncodePng32(
	std::string* dst, const char* data, int size,
	int width, int height, int width_step /*=0*/
) {
	if(dst == NULL || data == NULL || size <= 0) {
		return false;
	}
	if(width <= 0 || height <= 0) {
		return false;
	}

	if(width_step < width*4) {
		width_step = width*4;
	}

	std::string tmp;
	auto pSrcData = data;

	if(width_step > width*4) {
		tmp.resize(width*height*4);
		for(int i = 0; i < height; ++i) {
			auto ppTmp = (char*)tmp.data() + i*width*4;
			auto ppSrc = (char*)data + i*width_step;
			memcpy(ppTmp, ppSrc, width*4);
		}
		pSrcData = tmp.data();
	}

	unsigned char* png;
	size_t pngsize;

	auto err = lodepng_encode32(&png, &pngsize, (const unsigned char*)pSrcData, width, height);
	if(err != 0) return false;

	dst->assign((const char*)png, pngsize);
	free(png);

	return true;
}

bool EncodePng24(
	std::string* dst, const char* data, int size,
	int width, int height, int width_step /*=0*/
) {
	if(dst == NULL || data == NULL || size <= 0) {
		return false;
	}
	if(width <= 0 || height <= 0) {
		return false;
	}

	if(width_step < width*3) {
		width_step = width*3;
	}

	std::string tmp;
	auto pSrcData = data;

	if(width_step > width*3) {
		tmp.resize(width*height*3);
		for(int i = 0; i < height; ++i) {
			auto ppTmp = (char*)tmp.data() + i*width*3;
			auto ppSrc = (char*)data + i*width_step;
			memcpy(ppTmp, ppSrc, width*3);
		}
		pSrcData = tmp.data();
	}

	unsigned char* png;
	size_t pngsize;

	auto err = lodepng_encode24(&png, &pngsize, (const unsigned char*)pSrcData, width, height);
	if(err != 0) return false;

	dst->assign((const char*)png, pngsize);
	free(png);

	return true;
}