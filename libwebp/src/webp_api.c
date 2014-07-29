// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "webp.h"
#include "webp/encode.h"
#include "webp/decode.h"

#include <stdlib.h>
#include <string.h>

int webpGetInfo(
	const uint8_t* data, size_t data_size,
	int* width, int* height,
	int* has_alpha
) {
	WebPBitstreamFeatures features;
	if (WebPGetFeatures(data, data_size, &features) != VP8_STATUS_OK) {
		return 0;
	}
	if (width != NULL) {
		*width  = features.width;
	}
	if (height != NULL) {
		*height = features.height;
	}
	if (has_alpha != NULL) {
		*has_alpha = features.has_alpha;
	}
	return 1;
}

uint8_t* webpDecodeGray(
	const uint8_t* data, size_t data_size,
	int* width, int* height
) {
	int w, h;
	uint8_t *y, *u, *v;
	uint8_t *gray, *dst, *src;
	int stride, uv_stride;
	int i;

	if((y = WebPDecodeYUV(data, data_size, &w, &h, &u, &v, &stride, &uv_stride)) == NULL) {
		return NULL;
	}
	if (width != NULL) {
		*width  = w;
	}
	if (height != NULL) {
		*height = h;
	}

	if(stride == w) {
		return y;
	}

	if((gray = (uint8_t*)malloc(w*h)) == NULL) {
		free(y);
		return NULL;
	}

	src = y;
	dst = gray;
	for(i = 0; i < h; ++i) {
		memmove(dst, src, w);
		src += stride;
		dst += w;
	}

	free(y);
	return gray;
}

uint8_t* webpDecodeRGB(
	const uint8_t* data, size_t data_size,
	int* width, int* height
) {
	return WebPDecodeRGB(data, data_size, width, height);
}

uint8_t* webpDecodeRGBA(
	const uint8_t* data, size_t data_size,
	int* width, int* height
) {
	return WebPDecodeRGBA(data, data_size, width, height);
}

size_t webpEncodeGray(
	const uint8_t* gray, int width, int height, int stride, float quality_factor,
	uint8_t** output
) {
	size_t output_size;
	uint8_t* rgb;
	int x, y;

	if((rgb = (uint8_t*)malloc(width*height*3)) == NULL) {
		return 0;
	}
	for(y = 0; y < height; ++y) {
		const uint8_t* src = gray + y*stride;
		uint8_t* dst = rgb + y*width*3;
		for(x = 0; x < width; ++x) {
			uint8_t v = *src++;
			*dst++ = v;
			*dst++ = v;
			*dst++ = v;
		}
	}

	output_size = WebPEncodeRGB(rgb, width, height, width*3, quality_factor, output);
	free(rgb);
	return output_size;
}

size_t webpEncodeRGB(
	const uint8_t* rgb, int width, int height, int stride, float quality_factor,
	uint8_t** output
) {
	return WebPEncodeRGB(rgb, width, height, stride, quality_factor, output);
}

size_t webpEncodeRGBA(
	const uint8_t* rgba, int width, int height, int stride, float quality_factor,
	uint8_t** output
) {
	return WebPEncodeRGBA(rgba, width, height, stride, quality_factor, output);
}

size_t webpEncodeLosslessGray(
	const uint8_t* gray, int width, int height, int stride,
	uint8_t** output
) {
	size_t output_size;
	uint8_t* rgb;
	int x, y;

	if((rgb = (uint8_t*)malloc(width*height*3)) == NULL) {
		return 0;
	}
	for(y = 0; y < height; ++y) {
		const uint8_t* src = gray + y*stride;
		uint8_t* dst = rgb + y*width*3;
		for(x = 0; x < width; ++x) {
			uint8_t v = *src++;
			*dst++ = v;
			*dst++ = v;
			*dst++ = v;
		}
	}

	output_size = WebPEncodeLosslessRGB(rgb, width, height, width*3, output);
	free(rgb);
	return output_size;
}

size_t webpEncodeLosslessRGB(
	const uint8_t* rgb, int width, int height, int stride,
	uint8_t** output
) {
	return WebPEncodeLosslessRGB(rgb, width, height, stride, output);
}

size_t webpEncodeLosslessRGBA(
	const uint8_t* rgba, int width, int height, int stride,
	uint8_t** output
) {
	return WebPEncodeLosslessRGBA(rgba, width, height, stride, output);
}

void webpFree(void* p) {
	free(p);
}
