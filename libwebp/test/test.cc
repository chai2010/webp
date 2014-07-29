// Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "test.h"

#include <stdarg.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <math.h>
#include <time.h>

#include <string>
#include <algorithm>

static std::vector<std::string> args;
static std::string flag_list_regexp = "";
static std::string flag_test_regexp = ".*";
static std::string flag_test_bench_regexp = "";
static std::string flag_test_bench_benchtime_second = "1";

static struct { int N; double benchtime, timer_start, timer_duration; bool timer_on; } bench;
static struct { void (*fn)(void); const char *name, *type; } tests[10000];
static int ntests = 0;

static bool strHasPrefix(const std::string& str, const std::string& prefix) {
	return str.size() >= prefix.size()
		&& str.compare(0, prefix.size(), prefix) == 0;
}
static bool strHasSuffix(const std::string& str, const std::string& suffix) {
	return str.size() >= suffix.size()
		&& str.compare(str.size() - suffix.size(), suffix.size(), suffix) == 0;
}

static const char* getBaseName(const char* fname) {
	int len = strlen(fname);
	const char* s = fname + len;
	while(s > fname) {
		if(s[-1] == '/' || s[-1] == '\\') return s;
		s--;
	}
	return s;
}

static int matchhere(const char* regexp, const char* text);
static int match(const char *regexp, const char *text) {
	if (regexp[0] == '^') {
		return matchhere(regexp+1, text);
	}
	do {
		if (matchhere(regexp, text)) return 1;
	} while (*text++ != '\0');
	return 0;
}
static int matchstar(int c, const char *regexp, const char *text) {
	do {
		if (matchhere(regexp, text)) return 1;
	} while (*text != '\0' && (*text++ == c || c == '.'));
	return 0;
}
static int matchhere(const char *regexp, const char *text) {
	if (regexp[0] == '\0') return 1;
	if (regexp[1] == '*') return matchstar(regexp[0], regexp+2, text);
	if (regexp[0] == '$' && regexp[1] == '\0') return *text == '\0';
	if (*text!='\0' && (regexp[0]=='.' || regexp[0]==*text)) return matchhere(regexp+1, text+1);
	return 0;
}

const std::vector<std::string>& TestArgs() {
	return args;
}

void RegisterTest(void (*fn)(void), const char* name, const char* type) {
	if(ntests >= sizeof(tests)/sizeof(tests[0])) {
		printf("%s %s, line %d: RegisterTest failed\n", name, getBaseName(__FILE__), __LINE__);
		exit(-1);
	}
	tests[ntests].fn = fn;
	tests[ntests].name = name;
	tests[ntests].type = type;
	ntests++;
}

void TestAssertTrue(bool condition, const char* fname, int lineno, const char* fmt, ...) {
	if(!condition) {
		fname = getBaseName(fname);
		if(fmt != NULL && fmt[0] != '\0') {
			va_list ap;
			va_start(ap, fmt);
			printf("fail, %s, line %d: ASSERT_TRUE(false), ", fname, lineno);
			vprintf(fmt, ap);
			printf("\n");
			va_end(ap);
		} else {
			printf("fail, %s, line %d: ASSERT_TRUE(false)\n", fname, lineno);
		}
		exit(-1);
	}
}

void TestAssertEQ(int a, int b, const char* fname, int lineno, const char* fmt, ...) {
	if(a != b) {
		fname = getBaseName(fname);
		if(fmt != NULL && fmt[0] != '\0') {
			va_list ap;
			va_start(ap, fmt);
			printf("fail, %s, line %d: ASSERT_EQ(%d, %d), ", fname, lineno, a, b);
			vprintf(fmt, ap);
			printf("\n");
			va_end(ap);
		} else {
			printf("fail, %s, line %d: ASSERT_EQ(%d, %d)\n", fname, lineno, a, b);
		}
		exit(-1);
	}
}
void TestAssertStrEQ(const char* a, const char* b, const char* fname, int lineno, const char* fmt, ...) {
	if(strcmp(a, b) != 0) {
		fname = getBaseName(fname);
		if(fmt != NULL && fmt[0] != '\0') {
			va_list ap;
			va_start(ap, fmt);
			printf("fail, %s, line %d: ASSERT_STREQ(\"%s\", \"%s\"), ", fname, lineno, a, b);
			vprintf(fmt, ap);
			printf("\n");
			va_end(ap);
		} else {
			printf("fail, %s, line %d: ASSERT_STREQ(\"%s\", \"%s\")\n", fname, lineno, a, b);
		}
		exit(-1);
	}
}
void TestAssertNear(float a, float b, float abs_error, const char* fname, int lineno, const char* fmt, ...) {
	if(abs(a-b) > abs(abs_error)) {
		fname = getBaseName(fname);
		if(fmt != NULL && fmt[0] != '\0') {
			va_list ap;
			va_start(ap, fmt);
			printf("fail, %s, line %d: ASSERT_NEAR(%f, %f, %f), ", fname, lineno, a, b, abs_error);
			vprintf(fmt, ap);
			printf("\n");
			va_end(ap);
		} else {
			printf("fail, %s, line %d: ASSERT_NEAR(%f, %f, %f)\n", fname, lineno, a, b, abs_error);
		}
		exit(-1);
	}
}

// roundDown10 rounds a number down to the nearest power of 10.
static int roundDown10(int n) {
	int tens = 0;
	// tens = floor(log_10(n))
	while(n >= 10) {
		n = n / 10;
		tens++;
	}
	// result = 10^tens
	int result = 1;
	for(int i = 0; i < tens; ++i) {
		result *= 10;
	}
	return result;
}

// roundUp rounds x up to a number of the form [1eX, 2eX, 5eX].
static int roundUp(int n) {
	int base = roundDown10(n);
	if(n <= base) return base;
	if(n <= base*2) return base*2;
	if(n <= base*5) return base*5;
	return base*10;
}

static double timeNowSec() {
	return 1.0 * clock() / CLOCKS_PER_SEC;
}

static void benchRunN(int id, int n) {
	bench.N = n;
	BenchResetTimer();
	BenchStartTimer();
	tests[id].fn();
	BenchStopTimer();
}

static void benchRun(int id) {
	int n = 1;
	benchRunN(id, n);
	while(bench.timer_duration < bench.benchtime && n < 1e9) {
		int last = n;
		// Run more iterations than we think we'll need for a second (1.5x).
		// Don't grow too fast in case we had timing errors previously.
		// Be sure to run at least one more than last time.
		n = std::max(std::min(n+n/2, 100*last), last+1);
		// Round up to something easy to read.
		n = roundUp(n);
		benchRunN(id, n);
	}

	double nsop = 1e9*bench.timer_duration/bench.N;
	if(nsop < 10) {
		printf("[bench] %s %d %.2f ns/op\n", tests[id].name, bench.N, float(nsop));
	} else if(nsop < 100) {
		printf("[bench] %s %d %.1f ns/op\n", tests[id].name, bench.N, float(nsop));
	} else {
		printf("[bench] %s %d %d ns/op\n", tests[id].name, bench.N, int(nsop));
	}
}

int BenchN() {
	return bench.N;
}
void BenchResetTimer() {
	bench.timer_start = timeNowSec();
	bench.timer_duration = 0.0;
}
void BenchStartTimer() {
	if(!bench.timer_on) {
		bench.timer_start = timeNowSec();
		bench.timer_on = true;
	}
}
void BenchStopTimer() {
	if(bench.timer_on) {
		bench.timer_duration += timeNowSec() - bench.timer_start;
		bench.timer_on = false;
	}
}

static void usage(int argc, char* argv[]) {
	printf("C++ Mini UnitTest and Benchmark Library.\n");
	printf("https://github.com/chai2010/cc-mini-test\n");
	printf("\n");

	printf("Usage: %s\n", getBaseName(argv[0]));
	printf("  [-list=.*]\n");
	printf("  [-test=.*]\n");
	printf("  [-test.bench=]\n");
	printf("  [-test.benchtime=1second]\n");
	printf("  [-help]\n");
	printf("  [-h]\n");
	printf("\n");

	printf("Report bugs to <chaishushan{AT}gmail.com>.\n");
}

int main(int argc, char* argv[]) {
	args.assign(argv, argv + argc);
	for(int i = 1; i < argc; ++i) {
		if(argv[i] == std::string("-help") || argv[i] == std::string("-h")) {
			usage(argc, argv);
			return 0;
		}

		if(argv[i] == std::string("-list")) {
			flag_list_regexp = ".*";
			break;
		}
		if(argv[i] == std::string("-test")) {
			flag_test_regexp = ".*";
			continue;
		}
		if(argv[i] == std::string("-test.bench")) {
			flag_test_bench_regexp = ".*";
			continue;
		}

		if(strHasPrefix(argv[i], "-list=")) {
			flag_list_regexp = argv[i]+sizeof("-list=")-1;
			continue;
		}
		if(strHasPrefix(argv[i], "-test=")) {
			flag_test_regexp = argv[i]+sizeof("-test=")-1;
			continue;
		}
		if(strHasPrefix(argv[i], "-test.bench=")) {
			flag_test_bench_regexp = argv[i]+sizeof("-test.bench=")-1;
			continue;
		}

		if(strHasPrefix(argv[i], "-test.benchtime=")) {
			flag_test_bench_benchtime_second = argv[i]+sizeof("-test.benchtime=")-1;
			bench.benchtime = atof(flag_test_bench_benchtime_second.c_str());
			if(bench.benchtime <= 0.1) bench.benchtime = 1.0;
			continue;
		}

		// ingore user defined flag
	}

	if(!flag_list_regexp.empty()) {
		int total = 0;
		for(int id = 0; id < ntests; ++id) {
			if(std::string(tests[id].type) == "init") {
				if(match(flag_list_regexp.c_str(), tests[id].name) != 0) {
					printf("[init] %s\n", tests[id].name);
					total++;
				}
			}
		}
		for(int id = 0; id < ntests; ++id) {
			if(std::string(tests[id].type) == "exit") {
				if(match(flag_list_regexp.c_str(), tests[id].name) != 0) {
					printf("[exit] %s\n", tests[id].name);
					total++;
				}
			}
		}
		for(int id = 0; id < ntests; ++id) {
			if(std::string(tests[id].type) == "test") {
				if(match(flag_list_regexp.c_str(), tests[id].name) != 0) {
					printf("[test] %s\n", tests[id].name);
					total++;
				}
			}
		}
		for(int id = 0; id < ntests; ++id) {
			if(std::string(tests[id].type) == "bench") {
				if(match(flag_list_regexp.c_str(), tests[id].name) != 0) {
					printf("[bench] %s\n", tests[id].name);
					total++;
				}
			}
		}
		printf("total %d\n", total);
		return 0;
	}

	// run init func
	for(int id = 0; id < ntests; ++id) {
		if(std::string(tests[id].type) == "init") {
			printf("[init] %s ", tests[id].name);
			tests[id].fn();
			printf("\n");
		}
	}

	// run test func
	if(!flag_test_regexp.empty()) {
		for(int id = 0; id < ntests; ++id) {
			if(std::string(tests[id].type) == "test") {
				if(match(flag_test_regexp.c_str(), tests[id].name) != 0) {
					printf("[test] %s ", tests[id].name);
					tests[id].fn();
					printf("ok\n");
				}
			}
		}
	}

	// run bench func
	if(!flag_test_bench_regexp.empty()) {
		if(bench.benchtime <= 0.1) bench.benchtime = 1.0;
		for(int id = 0; id < ntests; ++id) {
			if(std::string(tests[id].type) == "bench") {
				if(match(flag_test_bench_regexp.c_str(), tests[id].name) != 0) {
					benchRun(id);
				}
			}
		}
	}

	// run exit func
	for(int id = 0; id < ntests; ++id) {
		if(std::string(tests[id].type) == "exit") {
			printf("[exit] %s ", tests[id].name);
			tests[id].fn();
			printf("\n");
		}
	}

	printf("PASS\n");
	return 0;
}
