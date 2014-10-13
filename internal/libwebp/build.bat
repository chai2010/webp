:: Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
:: Use of this source code is governed by a BSD-style
:: license that can be found in the LICENSE file.

setlocal

cd %~dp0

mkdir zz_build_tmp
cd    zz_build_tmp

cmake ..^
  -G "NMake Makefiles"^
  -DCMAKE_BUILD_TYPE=release^
  ^
  -DCMAKE_C_FLAGS_DEBUG="/MTd /Zi /Od /Ob0 /RTC1"^
  -DCMAKE_CXX_FLAGS_DEBUG="/MTd /Zi /Od /Ob0 /RTC1"^
  ^
  -DCMAKE_C_FLAGS_RELEASE="/MT /O2 /Ob2 /DNDEBUG"^
  -DCMAKE_CXX_FLAGS_RELEASE="/MT /O2 /Ob2 /DNDEBUG"^
  ^
  -DCMAKE_EXE_LINKER_FLAGS="/MANIFEST:NO"

nmake VERBOSE=1
nmake install

cd ..
