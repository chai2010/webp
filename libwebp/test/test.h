// Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#ifndef TEST_H_
#define TEST_H_

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
