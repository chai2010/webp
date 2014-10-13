// Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "test_util.h"

#include "./jpeg/jpge.cpp"
#include "./jpeg/jpgd.cpp"

#include <string>

bool DecodeJpeg(std::string* dst, const char* data, int size, int* width, int* height, int* channels) {
	if(dst == NULL || data == NULL || size <= 0) {
		return false;
	}
	if(width == NULL || height == NULL || channels == NULL) {
		return false;
	}

	auto p = jpgd::decompress_jpeg_image_from_memory(
		(const unsigned char *)data, size,
		width, height, channels,
		3
	);
	if(p == NULL) {
		return false;
	}
	if(*width <= 0 || *height <= 0 || (*channels != 1 && *channels != 3)) {
		free(p);
		return false;
	}

	// if Gray: convert to RGB;
	if(*channels == 1) {
		dst->resize((*width)*(*height));
		auto pDst = (unsigned char*)dst->data();
		for(int i = 0; i < (*width)*(*height); ++i) {
			pDst[i] = p[i*3];
		}
	} else {
		dst->assign((const char*)p, (*width)*(*height)*(*channels));
	}
	free(p);

	return true;
}

bool EncodeJpeg(
	std::string* dst, const char* data, int size,
	int width, int height, int channels, int quality /* =90 */,
	int width_step /* =0 */
) {
	if(dst == NULL || data == NULL || size <= 0) {
		return false;
	}
	if(width <= 0 || height <= 0) {
		return false;
	}
	if(channels != 1 && channels != 3) {
		return false;
	}
	if(quality <= 0 || quality > 100) {
		return false;
	}

	if(width_step < width*channels) {
		width_step = width*channels;
	}

	std::string tmp;
	auto pSrcData = data;

	jpge::params comp_params;
	if(channels == 3) {
		comp_params.m_subsampling = jpge::H2V2;   // RGB
		comp_params.m_quality = quality;

		if(width_step > width*channels) {
			tmp.resize(width*height*3);
			for(int i = 0; i < height; ++i) {
				auto ppTmp = (char*)tmp.data() + i*width*channels;
				auto ppSrc = (char*)data + i*width_step;
				memcpy(ppTmp, ppSrc, width*channels);
			}
			pSrcData = tmp.data();
		}
	} else {
		comp_params.m_subsampling = jpge::Y_ONLY; // Gray
		comp_params.m_quality = quality;

		// if Gray: convert to RGB;
		tmp.resize(width*height*3);
		for(int i = 0; i < height; ++i) {
			auto ppTmp = (char*)tmp.data() + i*width*3;
			auto ppSrc = (char*)data + i*width_step;
			for(int j = 0; j < width; ++j) {
				ppTmp[i*3+0] = ppSrc[i];
				ppTmp[i*3+1] = ppSrc[i];
				ppTmp[i*3+2] = ppSrc[i];
			}
		}
		channels = 3;
		pSrcData = tmp.data();
	}

	int buf_size = ((width*height*3)>1024)? (width*height*3): 1024;
	dst->resize(buf_size);
	bool rv = compress_image_to_jpeg_file_in_memory(
		(void*)dst->data(), buf_size, width, height, channels,
		(const jpge::uint8*)pSrcData,
		comp_params
	);
	if(!rv) {
		dst->clear();
		return false;
	}

	dst->resize(buf_size);
	return true;
}

