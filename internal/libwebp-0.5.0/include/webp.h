// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#ifndef _WEBP_H_
#define _WEBP_H_

#include <stddef.h>
#include <stdint.h>
#include <webp/decode.h>

#ifdef __cplusplus
extern "C" {
#endif

int webpGetInfo(
	const uint8_t* data, size_t data_size,
	int* width, int* height,
	int* has_alpha
);

uint8_t* webpDecodeGray(
	const uint8_t* data, size_t data_size,
	int* width, int* height
);
uint8_t* webpDecodeRGB(
	const uint8_t* data, size_t data_size,
	int* width, int* height
);
uint8_t* webpDecodeRGBA(
	const uint8_t* data, size_t data_size,
	int* width, int* height
);

uint8_t* webpEncodeGray(
	const uint8_t* gray, int width, int height, int stride, float quality_factor,
	size_t* output_size
);
uint8_t* webpEncodeRGB(
	const uint8_t* rgb, int width, int height, int stride, float quality_factor,
	size_t* output_size
);
uint8_t* webpEncodeRGBA(
	const uint8_t* rgba, int width, int height, int stride, float quality_factor,
	size_t* output_size
);

uint8_t* webpEncodeLosslessGray(
	const uint8_t* gray, int width, int height, int stride,
	size_t* output_size
);
uint8_t* webpEncodeLosslessRGB(
	const uint8_t* rgb, int width, int height, int stride,
	size_t* output_size
);
uint8_t* webpEncodeLosslessRGBA(
	const uint8_t* rgba, int width, int height, int stride,
	size_t* output_size
);

char* webpGetEXIF(const uint8_t* data, size_t data_size, size_t* metadata_size);
char* webpGetICCP(const uint8_t* data, size_t data_size, size_t* metadata_size);
char* webpGetXMP(const uint8_t* data, size_t data_size, size_t* metadata_size);

uint8_t* webpSetEXIF(const uint8_t* data, size_t data_size, const char* metadata, size_t metadata_size, size_t* new_data_size);
uint8_t* webpSetICCP(const uint8_t* data, size_t data_size, const char* metadata, size_t metadata_size, size_t* new_data_size);
uint8_t* webpSetXMP(const uint8_t* data, size_t data_size, const char* metadata, size_t metadata_size, size_t* new_data_size);

uint8_t* webpDelEXIF(const uint8_t* data, size_t data_size, size_t* new_data_size);
uint8_t* webpDelICCP(const uint8_t* data, size_t data_size, size_t* new_data_size);
uint8_t* webpDelXMP(const uint8_t* data, size_t data_size, size_t* new_data_size);

void* webpMalloc(size_t size);
void webpFree(void* p);

#ifdef __cplusplus
}
#endif
#endif // _WEBP_H_
