// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "webp.h"
#include "webp/encode.h"
#include "webp/decode.h"
#include "webp/demux.h"
#include "webp/mux.h"

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
	if(width != NULL) {
		*width  = features.width;
	}
	if(height != NULL) {
		*height = features.height;
	}
	if(has_alpha != NULL) {
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

int webpDecodeGrayToSize(const uint8_t* data, size_t data_size,
	int width, int height, int outStride, uint8_t* out
) {
	WebPDecoderConfig config;
	if(!WebPInitDecoderConfig(&config)) {
		return -1;
	}

	config.options.bypass_filtering = 1;
	config.options.no_fancy_upsampling = 1;
	config.options.use_scaling = 1;
	config.options.scaled_width = width;
	config.options.scaled_height = height;
	config.output.colorspace = MODE_YUV;

	int status = WebPDecode(data, data_size, &config);
	if(status != VP8_STATUS_OK) {
		return status;
	}

	int yStride = config.output.u.YUVA.y_stride;
	uint8_t* src = config.output.u.YUVA.y;
	uint8_t* dst = out;
	int i;

	for(i = 0; i < height; ++i) {
		memmove(dst, src, width);
		src += yStride;
		dst += outStride;
	}

	WebPFreeDecBuffer(&config.output);
	return status;
}

int webpDecodeRGBToSize(const uint8_t* data, size_t data_size,
	int width, int height, int outStride, uint8_t* out
) {
	WebPDecoderConfig config;
	if(!WebPInitDecoderConfig(&config)) {
		return -1;
	}

	config.options.bypass_filtering = 1;
	config.options.no_fancy_upsampling = 1;
	config.options.use_scaling = 1;
	config.options.scaled_width = width;
	config.options.scaled_height = height;
	config.output.colorspace = MODE_RGB;
	config.output.u.RGBA.rgba = out;
	config.output.u.RGBA.stride = outStride;
	config.output.u.RGBA.size = outStride * height;
	config.output.is_external_memory = 1;

	return WebPDecode(data, data_size, &config);
}

int webpDecodeRGBAToSize(const uint8_t* data, size_t data_size,
	int width, int height, int outStride, uint8_t* out
) {
	WebPDecoderConfig config;
	if(!WebPInitDecoderConfig(&config)) {
		return -1;
	}

	config.options.bypass_filtering = 1;
	config.options.no_fancy_upsampling = 1;
	config.options.use_scaling = 1;
	config.options.scaled_width = width;
	config.options.scaled_height = height;
	config.output.colorspace = MODE_RGBA;
	config.output.u.RGBA.rgba = out;
	config.output.u.RGBA.stride = outStride;
	config.output.u.RGBA.size = outStride * height;
	config.output.is_external_memory = 1;

	return WebPDecode(data, data_size, &config);
}

uint8_t* webpEncodeGray(
	const uint8_t* gray, int width, int height, int stride, float quality_factor,
	size_t* output_size
) {
	uint8_t* output;
	uint8_t* rgb;
	int x, y;

	if((rgb = (uint8_t*)malloc(width*height*3)) == NULL) {
		return NULL;
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

	*output_size = WebPEncodeRGB(rgb, width, height, width*3, quality_factor, &output);
	free(rgb);
	return output;
}

uint8_t* webpEncodeRGB(
	const uint8_t* rgb, int width, int height, int stride, float quality_factor,
	size_t* output_size
) {
	uint8_t* output = NULL;
	*output_size = WebPEncodeRGB(rgb, width, height, stride, quality_factor, &output);
	return output;
}

uint8_t* webpEncodeRGBA(
	const uint8_t* rgba, int width, int height, int stride, float quality_factor,
	size_t* output_size
) {
	uint8_t* output = NULL;
	*output_size = WebPEncodeRGBA(rgba, width, height, stride, quality_factor, &output);
	return output;
}

uint8_t* webpEncodeLosslessGray(
	const uint8_t* gray, int width, int height, int stride,
	size_t* output_size
) {
	uint8_t* output;
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

	*output_size = WebPEncodeLosslessRGB(rgb, width, height, width*3, &output);
	free(rgb);
	return output;
}

uint8_t* webpEncodeLosslessRGB(
	const uint8_t* rgb, int width, int height, int stride,
	size_t* output_size
) {
	uint8_t* output = NULL;
	*output_size = WebPEncodeLosslessRGB(rgb, width, height, stride, &output);
	return output;
}

uint8_t* webpEncodeLosslessRGBA(
	const uint8_t* rgba, int width, int height, int stride,
	size_t* output_size
) {
	uint8_t* output = NULL;
	*output_size = WebPEncodeLosslessRGBA(rgba, width, height, stride, &output);
	return output;
}

char* webpGetEXIF(const uint8_t* data, size_t data_size, size_t* metadata_size) {
	char* metadata = NULL;
	WebPData webp_data = {data, data_size};
	WebPDemuxer* demux = WebPDemux(&webp_data);
	uint32_t flags = WebPDemuxGetI(demux, WEBP_FF_FORMAT_FLAGS);
	*metadata_size = 0;
	if(flags & EXIF_FLAG) {
		WebPChunkIterator it;
		memset(&it, 0, sizeof(it));
		if(WebPDemuxGetChunk(demux, "EXIF", 1, &it)) {
			if(it.chunk.bytes != NULL && it.chunk.size > 0) {
				metadata = (char*)malloc(it.chunk.size);
				memcpy(metadata, it.chunk.bytes, it.chunk.size);
				*metadata_size = it.chunk.size;
			}
		}
		WebPDemuxReleaseChunkIterator(&it);
	}
	WebPDemuxDelete(demux);
	return metadata;
}
char* webpGetICCP(const uint8_t* data, size_t data_size, size_t* metadata_size) {
	char* metadata = NULL;
	WebPData webp_data = {data, data_size};
	WebPDemuxer* demux = WebPDemux(&webp_data);
	uint32_t flags = WebPDemuxGetI(demux, WEBP_FF_FORMAT_FLAGS);
	*metadata_size = 0;
	if(flags & ICCP_FLAG) {
		WebPChunkIterator it;
		memset(&it, 0, sizeof(it));
		if(WebPDemuxGetChunk(demux, "ICCP", 1, &it)) {
			if(it.chunk.bytes != NULL && it.chunk.size > 0) {
				metadata = (char*)malloc(it.chunk.size);
				memcpy(metadata, it.chunk.bytes, it.chunk.size);
				*metadata_size = it.chunk.size;
			}
		}
		WebPDemuxReleaseChunkIterator(&it);
	}
	WebPDemuxDelete(demux);
	return metadata;
}
char* webpGetXMP(const uint8_t* data, size_t data_size, size_t* metadata_size) {
	char* metadata = NULL;
	WebPData webp_data = {data, data_size};
	WebPDemuxer* demux = WebPDemux(&webp_data);
	uint32_t flags = WebPDemuxGetI(demux, WEBP_FF_FORMAT_FLAGS);
	*metadata_size = 0;
	if(flags & XMP_FLAG) {
		WebPChunkIterator it;
		memset(&it, 0, sizeof(it));
		if(WebPDemuxGetChunk(demux, "XMP ", 1, &it)) {
			if(it.chunk.bytes != NULL && it.chunk.size > 0) {
				metadata = (char*)malloc(it.chunk.size);
				memcpy(metadata, it.chunk.bytes, it.chunk.size);
				*metadata_size = it.chunk.size;
			}
		}
		WebPDemuxReleaseChunkIterator(&it);
	}
	WebPDemuxDelete(demux);
	return metadata;
}

uint8_t* webpSetEXIF(
	const uint8_t* data, size_t data_size,
	const char* metadata, size_t metadata_size,
	size_t* new_data_size
) {
	WebPData image = {data, data_size};
	WebPData profile = {metadata, metadata_size};
	WebPData output_data = {NULL, 0};
	WebPMux* mux = WebPMuxCreate(&image, 0);
	if(WebPMuxSetChunk(mux, "EXIF", &profile, 0) == WEBP_MUX_OK) {
		WebPMuxAssemble(mux, &output_data);
	}
	WebPMuxDelete(mux);
	*new_data_size = output_data.size;
	return (uint8_t*)(output_data.bytes);
}
uint8_t* webpSetICCP(
	const uint8_t* data, size_t data_size,
	const char* metadata, size_t metadata_size,
	size_t* new_data_size
) {
	WebPData image = {data, data_size};
	WebPData profile = {metadata, metadata_size};
	WebPData output_data = {NULL, 0};
	WebPMux* mux = WebPMuxCreate(&image, 0);
	if(WebPMuxSetChunk(mux, "ICCP", &profile, 0) == WEBP_MUX_OK) {
		WebPMuxAssemble(mux, &output_data);
	}
	WebPMuxDelete(mux);
	*new_data_size = output_data.size;
	return (uint8_t*)(output_data.bytes);
}
uint8_t* webpSetXMP(
	const uint8_t* data, size_t data_size,
	const char* metadata, size_t metadata_size,
	size_t* new_data_size
) {
	WebPData image = {data, data_size};
	WebPData profile = {metadata, metadata_size};
	WebPData output_data = {NULL, 0};
	WebPMux* mux = WebPMuxCreate(&image, 0);
	if(WebPMuxSetChunk(mux, "XMP ", &profile, 0) == WEBP_MUX_OK) {
		WebPMuxAssemble(mux, &output_data);
	}
	WebPMuxDelete(mux);
	*new_data_size = output_data.size;
	return (uint8_t*)(output_data.bytes);
}

uint8_t* webpDelEXIF(const uint8_t* data, size_t data_size, size_t* new_data_size) {
	WebPData image = {data, data_size};
	WebPData output_data = {NULL, 0};
	WebPMux* mux = WebPMuxCreate(&image, 0);
	if(WebPMuxDeleteChunk(mux, "EXIF") == WEBP_MUX_OK) {
		WebPMuxAssemble(mux, &output_data);
	}
	WebPMuxDelete(mux);
	*new_data_size = output_data.size;
	return (uint8_t*)(output_data.bytes);
}
uint8_t* webpDelICCP(const uint8_t* data, size_t data_size, size_t* new_data_size) {
	WebPData image = {data, data_size};
	WebPData output_data = {NULL, 0};
	WebPMux* mux = WebPMuxCreate(&image, 0);
	if(WebPMuxDeleteChunk(mux, "ICCP") == WEBP_MUX_OK) {
		WebPMuxAssemble(mux, &output_data);
	}
	WebPMuxDelete(mux);
	*new_data_size = output_data.size;
	return (uint8_t*)(output_data.bytes);
}
uint8_t* webpDelXMP(const uint8_t* data, size_t data_size, size_t* new_data_size) {
	WebPData image = {data, data_size};
	WebPData output_data = {NULL, 0};
	WebPMux* mux = WebPMuxCreate(&image, 0);
	if(WebPMuxDeleteChunk(mux, "XMP ") == WEBP_MUX_OK) {
		WebPMuxAssemble(mux, &output_data);
	}
	WebPMuxDelete(mux);
	*new_data_size = output_data.size;
	return (uint8_t*)(output_data.bytes);
}

void* webpMalloc(size_t size) {
	return malloc(size);
}

void webpFree(void* p) {
	free(p);
}
