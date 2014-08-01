// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// https://docs.google.com/document/d/1KzP6nWeVsobU1AtTHbchdrmA-_UoNLJGEPgZikHAzVg/pub

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <webp/types.h>
#include <webp/encode.h>
#include <webp/decode.h>

uint8_t interpolate(float v0, float v1, float x) {
	return (uint8_t)(v0 + x * (v1 - v0));
}

int webpWriter(const uint8_t* data, size_t data_size, const WebPPicture* const pic) {
	FILE* const out = (FILE*)pic->custom_ptr;
	return data_size ? (fwrite(data, data_size, 1, out) == 1) : 1;
}

int WebPImportGray(const uint8_t* gray_data, WebPPicture* pic) {
	int y, width, uv_width;
	if (pic == NULL || gray_data == NULL) return 0;
	pic->colorspace = WEBP_YUV420;
	if (!WebPPictureAlloc(pic)) {
		return 0;
	}
	width = pic->width;
	uv_width = (width + 1) >> 1;
	for(y = 0; y < pic->height; ++y) {
		memcpy(pic->y + y * pic->y_stride, gray_data, width);
		gray_data += width;
		if((y & 1) == 0) {
			memset(pic->u + (y >> 1) * pic->uv_stride, 128, uv_width);
			memset(pic->v + (y >> 1) * pic->uv_stride, 128, uv_width);
		}
	}
	return 1;
}

int main() {
	int ret = 0;
	unsigned int width = 640;
	unsigned int height = 480;
	unsigned int data_size = width*height;
	uint8_t *gray_data = (uint8_t*)malloc(data_size*sizeof(uint8_t));
	int i = 0;
	int h = 0;
	int w = 0;

	for(h = 0; h < height; h++) {
		for(w = 0; w < width; w++) {
			gray_data[i] = interpolate(0, 255, ((float)(w))/((float)(width)));
			i++;
		}
	}

	WebPPicture pic;
	WebPPictureInit(&pic);
	pic.writer = webpWriter;
	pic.width = width;
	pic.height = height;

	pic.custom_ptr = fopen("output.webp", "wb");
	if(!WebPImportGray(gray_data, &pic)) {
		printf("ERR: WebPImportGray failed!\n");
		return -1;
	}
	free(gray_data);

	WebPConfig config;
	WebPConfigInit(&config);
	WebPEncode(&config, &pic);

	WebPPictureFree(&pic);

	uint8_t* rgb = (uint8_t*)malloc(width*height*3);
	uint8_t* output;
	size_t output_size;
	output_size = WebPEncodeLosslessRGB(rgb, width, height, width*3, &output);

	int new_w, new_h;
	if(!WebPGetInfo(output, output_size, &new_w, &new_h)) {
		printf("ERR: WebPImportGray failed!\n");
		return -1;
	}
	printf("WebPGetInfo: width = %d, height = %d\n", new_w, new_h);

	return 0;
}
