# Copyright 2022 <chaishushan{AT}gmail.com>. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

project(WEBP)

#------------------------------------------------------------------------------

IF(WIN32)
  if(CMAKE_SIZEOF_VOID_P EQUAL 8)
    set(OS win64)
  else()
    set(OS win32)
  endif()
else()
  if(CMAKE_SIZEOF_VOID_P EQUAL 8)
    set(OS posix64)
  else()
    set(OS posix64)
  endif()
endif()

#------------------------------------------------------------------------------

# Extra flags to enable experimental features and code
add_definitions(
  #-DWEBP_EXPERIMENTAL_FEATURES
  #-DWEBP_HAVE_AVX2
)

include_directories(AFTER
  ../include
  ../libwebp-1.2.2
  ../libwebp-1.2.2/src
)

set(WEBP_DEC_SRC
  ../libwebp-1.2.2/src/dec/alpha_dec.c
  ../libwebp-1.2.2/src/dec/buffer_dec.c
  ../libwebp-1.2.2/src/dec/frame_dec.c
  ../libwebp-1.2.2/src/dec/idec_dec.c
  ../libwebp-1.2.2/src/dec/io_dec.c
  ../libwebp-1.2.2/src/dec/quant_dec.c
  ../libwebp-1.2.2/src/dec/tree_dec.c
  ../libwebp-1.2.2/src/dec/vp8_dec.c
  ../libwebp-1.2.2/src/dec/vp8l_dec.c
  ../libwebp-1.2.2/src/dec/webp_dec.c
)

set(WEBP_DEMUX_SRC
  ../libwebp-1.2.2/src/demux/demux.c
)

set(WEBP_DSP_DEC_SRC
  ../libwebp-1.2.2/src/dsp/alpha_processing.c
  ../libwebp-1.2.2/src/dsp/alpha_processing_mips_dsp_r2.c
  ../libwebp-1.2.2/src/dsp/alpha_processing_neon.c
  ../libwebp-1.2.2/src/dsp/alpha_processing_sse2.c
  ../libwebp-1.2.2/src/dsp/alpha_processing_sse41.c
  ../libwebp-1.2.2/src/dsp/cost.c
  ../libwebp-1.2.2/src/dsp/cost_mips_dsp_r2.c
  ../libwebp-1.2.2/src/dsp/cost_mips32.c
  ../libwebp-1.2.2/src/dsp/cost_neon.c
  ../libwebp-1.2.2/src/dsp/cost_sse2.c
  ../libwebp-1.2.2/src/dsp/cpu.c
  ../libwebp-1.2.2/src/dsp/dec.c
  ../libwebp-1.2.2/src/dsp/dec_clip_tables.c
  ../libwebp-1.2.2/src/dsp/dec_mips32.c
  ../libwebp-1.2.2/src/dsp/dec_neon.c
  ../libwebp-1.2.2/src/dsp/dec_sse2.c
  ../libwebp-1.2.2/src/dsp/dec_sse41.c
  ../libwebp-1.2.2/src/dsp/filters.c
  ../libwebp-1.2.2/src/dsp/filters_mips_dsp_r2.c
  ../libwebp-1.2.2/src/dsp/filters_msa.c
  ../libwebp-1.2.2/src/dsp/filters_neon.c
  ../libwebp-1.2.2/src/dsp/filters_sse2.c
  ../libwebp-1.2.2/src/dsp/lossless.c
  ../libwebp-1.2.2/src/dsp/lossless_enc.c
  ../libwebp-1.2.2/src/dsp/lossless_enc_mips_dsp_r2.c
  ../libwebp-1.2.2/src/dsp/lossless_enc_mips32.c
  ../libwebp-1.2.2/src/dsp/lossless_enc_msa.c
  ../libwebp-1.2.2/src/dsp/lossless_enc_neon.c
  ../libwebp-1.2.2/src/dsp/lossless_enc_sse2.c
  ../libwebp-1.2.2/src/dsp/lossless_enc_sse41.c
  ../libwebp-1.2.2/src/dsp/lossless_neon.c
  ../libwebp-1.2.2/src/dsp/lossless_sse2.c
  ../libwebp-1.2.2/src/dsp/lossless_sse41.c
  ../libwebp-1.2.2/src/dsp/rescaler.c
  ../libwebp-1.2.2/src/dsp/rescaler_mips_dsp_r2.c
  ../libwebp-1.2.2/src/dsp/rescaler_msa.c
  ../libwebp-1.2.2/src/dsp/rescaler_neon.c
  ../libwebp-1.2.2/src/dsp/rescaler_sse2.c
  ../libwebp-1.2.2/src/dsp/ssim.c
  ../libwebp-1.2.2/src/dsp/ssim_sse2.c
  ../libwebp-1.2.2/src/dsp/upsampling.c
  ../libwebp-1.2.2/src/dsp/upsampling_mips_dsp_r2.c
  ../libwebp-1.2.2/src/dsp/upsampling_msa.c
  ../libwebp-1.2.2/src/dsp/upsampling_neon.c
  ../libwebp-1.2.2/src/dsp/upsampling_sse2.c
  ../libwebp-1.2.2/src/dsp/upsampling_sse41.c
  ../libwebp-1.2.2/src/dsp/yuv.c
  ../libwebp-1.2.2/src/dsp/yuv_mips_dsp_r2.c
  ../libwebp-1.2.2/src/dsp/yuv_mips32.c
  ../libwebp-1.2.2/src/dsp/yuv_neon.c
  ../libwebp-1.2.2/src/dsp/yuv_sse2.c
  ../libwebp-1.2.2/src/dsp/yuv_sse41.c
)

set(WEBP_DSP_ENC_SRC
  ../libwebp-1.2.2/src/dsp/enc.c
  ../libwebp-1.2.2/src/dsp/enc_mips32.c
  ../libwebp-1.2.2/src/dsp/enc_neon.c
  ../libwebp-1.2.2/src/dsp/enc_sse2.c
  ../libwebp-1.2.2/src/dsp/enc_sse41.c
)

set(WEBP_EX_FORMAT_DEC_SRC
  ../libwebp-1.2.2/examples/jpegdec.c
  ../libwebp-1.2.2/examples/metadata.c
  ../libwebp-1.2.2/examples/pngdec.c
  ../libwebp-1.2.2/examples/tiffdec.c
  ../libwebp-1.2.2/examples/webpdec.c
  ../libwebp-1.2.2/examples/wicdec.c
)

set(WEBP_EX_UTIL_SRC
  ../libwebp-1.2.2/examples/example_util.c
)

set(WEBP_ENC_SRC
  ../libwebp-1.2.2/src/enc/alpha_enc.c
  ../libwebp-1.2.2/src/enc/analysis_enc.c
  ../libwebp-1.2.2/src/enc/backward_references_enc.c
  ../libwebp-1.2.2/src/enc/backward_references_cost_enc.c
  ../libwebp-1.2.2/src/enc/config_enc.c
  ../libwebp-1.2.2/src/enc/cost_enc.c
  ../libwebp-1.2.2/src/enc/filter_enc.c
  ../libwebp-1.2.2/src/enc/frame_enc.c
  ../libwebp-1.2.2/src/enc/histogram_enc.c
  ../libwebp-1.2.2/src/enc/iterator_enc.c
  ../libwebp-1.2.2/src/enc/near_lossless_enc.c
  ../libwebp-1.2.2/src/enc/picture_enc.c
  ../libwebp-1.2.2/src/enc/picture_csp_enc.c
  ../libwebp-1.2.2/src/enc/picture_psnr_enc.c
  ../libwebp-1.2.2/src/enc/picture_rescale_enc.c
  ../libwebp-1.2.2/src/enc/picture_tools_enc.c
  ../libwebp-1.2.2/src/enc/predictor_enc.c
  ../libwebp-1.2.2/src/enc/quant_enc.c
  ../libwebp-1.2.2/src/enc/syntax_enc.c
  ../libwebp-1.2.2/src/enc/token_enc.c
  ../libwebp-1.2.2/src/enc/tree_enc.c
  ../libwebp-1.2.2/src/enc/vp8l_enc.c
  ../libwebp-1.2.2/src/enc/webp_enc.c
)

set(WEBP_MUX_SRC
  ../libwebp-1.2.2/src/mux/muxedit.c
  ../libwebp-1.2.2/src/mux/muxinternal.c
  ../libwebp-1.2.2/src/mux/muxread.c
)

set(WEBP_UTILS_DEC_SRC
  ../libwebp-1.2.2/src/utils/bit_reader_utils.c
  ../libwebp-1.2.2/src/utils/color_cache_utils.c
  ../libwebp-1.2.2/src/utils/filters_utils.c
  ../libwebp-1.2.2/src/utils/huffman_utils.c
  ../libwebp-1.2.2/src/utils/quant_levels_dec_utils.c
  ../libwebp-1.2.2/src/utils/rescaler_utils.c
  ../libwebp-1.2.2/src/utils/random_utils.c
  ../libwebp-1.2.2/src/utils/thread_utils.c
  ../libwebp-1.2.2/src/utils/utils.c
)

set(WEBP_UTILS_ENC_SRC
  ../libwebp-1.2.2/src/utils/bit_writer_utils.c
  ../libwebp-1.2.2/src/utils/huffman_encode_utils.c
  ../libwebp-1.2.2/src/utils/quant_levels_utils.c
)

set(WEBP_SRC
  ${WEBP_DEC_SRC}
  ${WEBP_DEMUX_SRC}
  ${WEBP_DSP_DEC_SRC}
  ${WEBP_DSP_ENC_SRC}
  ${WEBP_ENC_SRC}
  ${WEBP_MUX_SRC}
  ${WEBP_UTILS_DEC_SRC}
  ${WEBP_UTILS_ENC_SRC}
)

#------------------------------------------------------------------------------

add_library(webplib STATIC
  ../include/webp.h
  ../src/webp.c

  ${WEBP_SRC}
)
if(CMAKE_BUILD_TYPE STREQUAL "debug")
  set_target_properties(webplib
    PROPERTIES OUTPUT_NAME "webp-${OS}-debug"
  )
else()
  set_target_properties(webplib
    PROPERTIES OUTPUT_NAME "webp-${OS}"
  )
endif()

install(TARGETS webplib
  RUNTIME DESTINATION ${CMAKE_CURRENT_SOURCE_DIR}
  LIBRARY DESTINATION ${CMAKE_CURRENT_SOURCE_DIR}
  ARCHIVE DESTINATION ${CMAKE_CURRENT_SOURCE_DIR}
)

#------------------------------------------------------------------------------
# WebP Demo

# webpapp
add_executable(webpapp
  ../libwebp/demo/demo.cc

  ../include/webp.h
  ../src/webp.c

  ${WEBP_SRC}
)
install(TARGETS webpapp
  RUNTIME DESTINATION ${CMAKE_CURRENT_SOURCE_DIR}
  LIBRARY DESTINATION ${CMAKE_CURRENT_SOURCE_DIR}
  ARCHIVE DESTINATION ${CMAKE_CURRENT_SOURCE_DIR}
)

#------------------------------------------------------------------------------
# test

include_directories(AFTER
  ../libwebp/test
)

add_executable(webptest
  ../libwebp/test/test.cc
  ../libwebp/test/test_util.cc
  ../libwebp/test/test_util_jpg.cc

  #./test/webp_test.cc
  #./test/webp_bench_test.cc

  ../include/webp.h
  ../src/webp.c

  ${WEBP_SRC}
)

install(TARGETS webptest
  RUNTIME DESTINATION ${CMAKE_CURRENT_SOURCE_DIR}
  LIBRARY DESTINATION ${CMAKE_CURRENT_SOURCE_DIR}
  ARCHIVE DESTINATION ${CMAKE_CURRENT_SOURCE_DIR}
)

#------------------------------------------------------------------------------
