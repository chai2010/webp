// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "test.h"
#include "test_util.h"

#include "webp.h"

struct tImgInfo {
	int width;
	int height;
	int channels;
	int depth;
	const char* name;
};

static tImgInfo testCaseJpg[] = {
	{ 512 , 512 , 3, 8, "testdata/video-001.jpg" },
};

static tImgInfo testCasePng[] = {
	{ 512 , 512 , 3, 8, "testdata/video-001.png" },
};

static tImgInfo testCaseWebP[] = {
	{ 512 , 512 , 3, 8, "testdata/video-001.webp" },
};

TEST(webp, JpgHelper) {
	auto buf = new std::string;
	auto src = new std::string;
	auto dst = new std::string;

	for(int i = 0; i < TEST_DIM(testCaseJpg); ++i) {
		bool rv = loadImageData(testCaseJpg[i].name, buf);
		ASSERT_TRUE(rv);

		// decode raw file data
		int width, height, channels;
		rv = jpegDecode(src, buf->data(), buf->size(), &width, &height, &channels);
		ASSERT_TRUE(rv);
		ASSERT_TRUE(width == testCaseJpg[i].width);
		ASSERT_TRUE(height == testCaseJpg[i].height);
		ASSERT_TRUE(channels == testCaseJpg[i].channels);

		// encode as jpg
		buf->clear();
		rv = jpegEncode(buf, src->data(), src->size(), width, height, channels, 90, 0);
		ASSERT_TRUE(rv);

		// decode again
		rv = jpegDecode(dst, buf->data(), buf->size(), &width, &height, &channels);
		ASSERT_TRUE(rv);
		ASSERT_TRUE(width == testCaseJpg[i].width);
		ASSERT_TRUE(height == testCaseJpg[i].height);
		ASSERT_TRUE(channels == testCaseJpg[i].channels);

		// compare
		double diff = diffImageData(
			(const unsigned char*)src->data(), (const unsigned char*)dst->data(),
			width, height, channels
		);
		ASSERT_TRUE(diff < 20);
	}

	delete buf;
	delete src;
	delete dst;
}

TEST(webp, DecodeConfig) {
	auto buf = new std::string;
	for(int i = 0; i < TEST_DIM(testCaseJxr); ++i) {
		bool rv = loadImageData(testCaseJxr[i].name, buf);
		ASSERT_TRUE(rv);

		// decode webp data
		int width, height, channels, depth;
		jxr_data_type_t type;
		jxr_bool_t ret = jxr_decode_config(buf->data(), buf->size(),
			&width, &height, &channels, &depth, &type
		);
		ASSERT_TRUE(ret == jxr_true);
		ASSERT_TRUE(width == testCaseJxr[i].width);
		ASSERT_TRUE(height == testCaseJxr[i].height);
		ASSERT_TRUE(channels == testCaseJxr[i].channels);
		ASSERT_TRUE(depth == testCaseJxr[i].depth);
	}
	delete buf;
}

TEST(webp, Decode) {
	auto buf = new std::string;
	auto src = new std::string;
	auto dst = new std::string;

	for(int i = 0; i < TEST_DIM(testCaseJxr); ++i) {
		int width, height, channels, depth;
		jxr_data_type_t type;

		// decode webp data
		bool rv = loadImageData(testCaseJxr[i].name, buf);
		ASSERT_TRUE(rv);
		src->resize(testCaseJxr[i].width*testCaseJxr[i].height*testCaseJxr[i].channels);
		int n = jxr_decode(
			(char*)src->data(), src->size(), 0, buf->data(), buf->size(),
			&width, &height, &channels, &depth, &type
		);
		ASSERT_TRUE(n == jxr_true);
		ASSERT_TRUE(width == testCaseJxr[i].width);
		ASSERT_TRUE(height == testCaseJxr[i].height);
		ASSERT_TRUE(channels == testCaseJxr[i].channels);
		ASSERT_TRUE(depth == testCaseJxr[i].depth);

		// decode jpg data
		rv = loadImageData(testCaseJpg[i].name, buf);
		ASSERT_TRUE(rv);
		rv = jpegDecode(dst, buf->data(), buf->size(), &width, &height, &channels);
		ASSERT_TRUE(rv);
		ASSERT_TRUE(width == testCaseJpg[i].width);
		ASSERT_TRUE(height == testCaseJpg[i].height);
		ASSERT_TRUE(channels == testCaseJpg[i].channels);

		// compare
		double diff = diffImageData(
			(const unsigned char*)src->data(), (const unsigned char*)dst->data(),
			width, height, channels
		);
		ASSERT_TRUE(diff < 20);
	}

	delete buf;
	delete src;
	delete dst;
}

TEST(webp, Encode) {
	//
}

TEST(webp, DecodeAndEncode) {
	return; // skip

	auto buf = new std::string;
	auto src = new std::string;
	auto dst = new std::string;

	for(int i = 0; i < TEST_DIM(testCaseJxr); ++i) {
		int width, height, channels, depth;
		jxr_data_type_t type;
		int newSize;

		// decode raw file data
		bool rv = loadImageData(testCaseJxr[i].name, buf);
		ASSERT_TRUE(rv);
		src->resize(testCaseJxr[i].width*testCaseJxr[i].height*testCaseJxr[i].channels);
		int n = jxr_decode(
			(char*)src->data(), src->size(), 0, buf->data(), buf->size(),
			&width, &height, &channels, &depth, &type
		);
		ASSERT_TRUE(n == jxr_true);
		ASSERT_TRUE(width == testCaseJxr[i].width);
		ASSERT_TRUE(height == testCaseJxr[i].height);
		ASSERT_TRUE(channels == testCaseJxr[i].channels);
		ASSERT_TRUE(depth == testCaseJxr[i].depth);

		// encode as webp
		buf->clear();
		buf->resize(src->size());
		n = jxr_encode(
			(char*)buf->data(), buf->size(), src->data(), src->size(), 0,
			width, height, channels, depth,
			90, jxr_unsigned,
			&newSize
		);
		ASSERT_TRUE(n == jxr_true);
		ASSERT_TRUE(newSize > 0);

		// decode again
		dst->resize(testCaseJxr[i].width*testCaseJxr[i].height*testCaseJxr[i].channels);
		n = jxr_decode(
			(char*)dst->data(), dst->size(), 0, buf->data(), newSize,
			&width, &height, &channels, &depth, &type
		);
		ASSERT_TRUE(n == jxr_true);
		ASSERT_TRUE(width == testCaseJxr[i].width);
		ASSERT_TRUE(height == testCaseJxr[i].height);
		ASSERT_TRUE(channels == testCaseJxr[i].channels);
		ASSERT_TRUE(depth == testCaseJxr[i].depth);

		// compare
		if(depth == 8) {
			double diff = diffImageData(
				(const unsigned char*)src->data(), (const unsigned char*)dst->data(),
				width, height, channels
			);
			ASSERT_TRUE(diff < 20);
		} else if(depth == 16) {
			double diff = diffImageData(
				(const unsigned short*)src->data(), (const unsigned short*)dst->data(),
				width, height, channels
			);
			ASSERT_TRUE(diff < 20);
		} else {
			ASSERT_TRUE(false);
		}
	}

	delete buf;
	delete src;
	delete dst;
}

TEST(webp, CompareJxrJpg) {
	// diff(webp, jpg) < 20
}

// benchmark
