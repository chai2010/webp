# Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
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
  ./include
  ./src/webp
  ./src
)

set(WEBP_DEC_SRC
  ./src/dec/alpha.c
  ./src/dec/buffer.c
  ./src/dec/frame.c
  ./src/dec/idec.c
  ./src/dec/io.c
  ./src/dec/quant.c
  ./src/dec/tree.c
  ./src/dec/vp8.c
  ./src/dec/vp8l.c
  ./src/dec/webp.c
)

set(WEBP_DEMUX_SRC
  ./src/demux/demux.c
)

set(WEBP_DSP_DEC_SRC
  ./src/dsp/alpha_processing.c
  ./src/dsp/cpu.c
  ./src/dsp/dec.c
  ./src/dsp/dec_clip_tables.c
  ./src/dsp/dec_mips32.c
  ./src/dsp/dec_neon.c
  ./src/dsp/dec_sse2.c
  ./src/dsp/lossless.c
  ./src/dsp/lossless_mips32.c
  ./src/dsp/lossless_neon.c
  ./src/dsp/lossless_sse2.c
  ./src/dsp/upsampling.c
  ./src/dsp/upsampling_neon.c
  ./src/dsp/upsampling_sse2.c
  ./src/dsp/yuv.c
  ./src/dsp/yuv_mips32.c
  ./src/dsp/yuv_sse2.c
)

set(WEBP_DSP_ENC_SRC
  ./src/dsp/enc.c
  ./src/dsp/enc_avx2.c
  ./src/dsp/enc_mips32.c
  ./src/dsp/enc_neon.c
  ./src/dsp/enc_sse2.c
)

set(WEBP_EX_FORMAT_DEC_SRC
  ./examples/jpegdec.c
  ./examples/metadata.c
  ./examples/pngdec.c
  ./examples/tiffdec.c
  ./examples/webpdec.c
  ./examples/wicdec.c
)

set(WEBP_EX_UTIL_SRC
  ./examples/example_util.c
)

set(WEBP_ENC_SRC
  ./src/enc/alpha.c
  ./src/enc/analysis.c
  ./src/enc/backward_references.c
  ./src/enc/config.c
  ./src/enc/cost.c
  ./src/enc/filter.c
  ./src/enc/frame.c
  ./src/enc/histogram.c
  ./src/enc/iterator.c
  ./src/enc/picture.c
  ./src/enc/picture_csp.c
  ./src/enc/picture_psnr.c
  ./src/enc/picture_rescale.c
  ./src/enc/picture_tools.c
  ./src/enc/quant.c
  ./src/enc/syntax.c
  ./src/enc/token.c
  ./src/enc/tree.c
  ./src/enc/vp8l.c
  ./src/enc/webpenc.c
)

set(WEBP_MUX_SRC
  ./src/mux/muxedit.c
  ./src/mux/muxinternal.c
  ./src/mux/muxread.c
)

set(WEBP_UTILS_DEC_SRC
  ./src/utils/bit_reader.c
  ./src/utils/color_cache.c
  ./src/utils/filters.c
  ./src/utils/huffman.c
  ./src/utils/quant_levels_dec.c
  ./src/utils/rescaler.c
  ./src/utils/random.c
  ./src/utils/thread.c
  ./src/utils/utils.c
)

set(WEBP_UTILS_ENC_SRC
  ./src/utils/bit_writer.c
  ./src/utils/huffman_encode.c
  ./src/utils/quant_levels.c
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
  ./include/webp.h
  ./src/webp.c

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

include_directories(AFTER
  ./webplib
)

# webpapp
add_executable(webpapp
  ./demo/demo.cc

  ./include/webp.h
  ./src/webp.c

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
  ./test
)

add_executable(webptest
  ./test/test.cc
  ./test/test_util.cc
  ./test/test_util_jpg.cc

  #./test/webp_test.cc
  #./test/webp_bench_test.cc

  ./include/webp.h
  ./src/webp.c

  ${WEBP_SRC}
)

install(TARGETS webptest
  RUNTIME DESTINATION ${CMAKE_CURRENT_SOURCE_DIR}
  LIBRARY DESTINATION ${CMAKE_CURRENT_SOURCE_DIR}
  ARCHIVE DESTINATION ${CMAKE_CURRENT_SOURCE_DIR}
)

#------------------------------------------------------------------------------
