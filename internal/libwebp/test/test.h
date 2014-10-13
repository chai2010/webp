// Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#ifndef TEST_H_
#define TEST_H_

/*
# C++ Mini UnitTest and Benchmark Library

https://github.com/chai2010/cc-mini-test

This package implements a C++ mini unittest and benchmark library.

Talk: http://go-talks.appspot.com/github.com/chai2010/talks/chai2010-cc-mini-test-intro.slide

## Unittest

Use `TEST` define a unit test:

	#include "test.h"

	// 1, 1, 2, 3, 5, ...
	int Fibonacci(int i) {
		return (i < 2)? 1: Fibonacci(i-1) + Fibonacci(i-2);
	}

	TEST(Fibonacci, Simple) {
		ASSERT_TRUE(Fibonacci(0) == 1);
		ASSERT_TRUE(Fibonacci(1) == 1);
		ASSERT_TRUE(Fibonacci(2) == 2);
		ASSERT_TRUE(Fibonacci(3) == 3);
		ASSERT_TRUE(Fibonacci(4) == 5);
		ASSERT_TRUE(Fibonacci(5) == 8);
	}

	TEST(Fibonacci, All) {
		static const int fib[] = { 1, 1, 2, 3, 5, ... };
		for(int i = 0; i < sizeof(fib)/sizeof(fib[0]); ++i) {
			ASSERT_TRUE_MSG(Fibonacci(i) == fib[i],
				"failed on fib[%d], expected = %d, got = %d",
				i, fib[i], Fibonacci(i)
			);
		}
	}

Run test: `./a.out` (or `./a.out -test=regexp`):

	[test] Fibonacci.Simple ok
	[test] Fibonacci.All ok
	PASS


## Benchmark

Use `BENCH` define a bench test:

	BENCH(Fibonacci, 5) {
		for(int i = 0; i < BenchN(); ++i) {
			Fibonacci(5);
		}
	}
	BENCH(Fibonacci, 10) {
		for(int i = 0; i < BenchN(); ++i) {
			Fibonacci(10);
		}
	}
	BENCH(Fibonacci, 15) {
		for(int i = 0; i < BenchN(); ++i) {
			Fibonacci(15);
		}
	}

Run benchmark: `./a.out -test.bench`:

	[bench] Fibonacci.5 20000000 65.5 ns/op
	[bench] Fibonacci.10 2000000 763 ns/op
	[bench] Fibonacci.15 200000 8740 ns/op

The output means that the loop ran 20000000 times at a speed of 65.5 ns per loop.

If a benchmark needs some expensive setup before running, the timer may be reset:

	BENCH(Name, case1) {
		auto big = NewBig();
		BenchResetTimer();

		for(int i = 0; i < BenchN(); ++i) {
			big.Len();
		}
	}

	BENCH(Name, case2) {
		BenchStopTimer();
		auto big = NewBig();
		BenchStartTimer();

		for(int i = 0; i < BenchN(); ++i) {
			big.Len();
		}
		BenchStopTimer();
		delete big;
	}


## Init and Exit

We can use `INIT` define a init func, and use `EXIT` define a exit func:

	INIT(Fibonacci, init) {
		// do some init work
	}
	EXIT(Fibonacci, exit) {
		// do some clean work
	}

The init funcs run before the tests, the exit funcs run after the tests.

## Usage

	./a.out -help
	usage: a.out
	  [-list=.*]
	  [-test=.*]
	  [-test.bench=]
	  [-test.benchtime=1second]
	  [-help]
	  [-h]

## BUGS

Please report bugs to <chaishushan@gmail.com>.

Thanks!
*/

#include <string>
#include <vector>

#define INIT(x, y) \
	static void _init_##x##y(void); \
	static TestRegisterer _r_init_##x##y(_init_##x##y, # x "." # y , "init"); \
	static void _init_##x##y(void)

#define EXIT(x, y) \
	static void _exit_##x##y(void); \
	static TestRegisterer _r_exit_##x##y(_exit_##x##y, # x "." # y , "exit"); \
	static void _exit_##x##y(void)

#define TEST(x, y) \
	static void _test_##x##y(void); \
	static TestRegisterer _r_test_##x##y(_test_##x##y, # x "." # y , "test"); \
	static void _test_##x##y(void)

#define BENCH(x, y) \
	static void _bench_##x##y(void); \
	static TestRegisterer _r_bench_##x##y(_bench_##x##y, # x "." # y , "bench"); \
	static void _bench_##x##y(void)

#define ASSERT_TRUE(x) TestAssertTrue((x), __FILE__, __LINE__, "")
#define ASSERT_EQ(x, y) TestAssertEQ((x), (y), __FILE__, __LINE__, "")
#define ASSERT_STREQ(x, y) TestAssertStrEQ((x), (y), __FILE__, __LINE__, "")
#define ASSERT_NEAR(x, y, abs_error) TestAssertNear((x), (y), (abs_error), __FILE__, __LINE__, "")

#if !defined(_MSC_VER) || (_MSC_VER >= 1600)
#	define ASSERT_TRUE_MSG(x, fmt, ...) TestAssertTrue((x), __FILE__, __LINE__, (fmt), __VA_ARGS__)
#	define ASSERT_EQ_MSG(x, y, fmt, ...) TestAssertEQ((x), (y), __FILE__, __LINE__, (fmt), __VA_ARGS__)
#	define ASSERT_STREQ_MSG(x, y, fmt, ...) TestAssertStrEQ((x), (y), __FILE__, __LINE__, (fmt), __VA_ARGS__)
#	define ASSERT_NEAR_MSG(x, y, abs_error, fmt, ...) TestAssertNear((x), (y), (abs_error), __FILE__, __LINE__, (fmt), __VA_ARGS__)
#endif

const std::vector<std::string>& TestArgs();

void RegisterTest(void (*fn)(void), const char *name, const char *type);

void TestAssertTrue(bool condition, const char* fname, int lineno, const char* fmt, ...);
void TestAssertEQ(int a, int b, const char* fname, int lineno, const char* fmt, ...);
void TestAssertStrEQ(const char* a, const char* b, const char* fname, int lineno, const char* fmt, ...);
void TestAssertNear(float a, float b, float abs_error, const char* fname, int lineno, const char* fmt, ...);

int  BenchN();
void BenchResetTimer();
void BenchStartTimer();
void BenchStopTimer();

struct TestRegisterer {
	TestRegisterer(void (*fn)(void), const char *name, const char* type) {
		RegisterTest(fn, name, type);
	}
};

#endif  // TEST_H_
