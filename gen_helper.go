// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore
// +build ignore

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

var (
	oldGenFiles = make(map[string]bool)
)

func main() {
	clearOldGenFiles()
	genIncludeFiles()
	printOldGenFiles()
}

func clearOldGenFiles() {
	ss, err := filepath.Glob("z_*.c")
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(ss); i++ {
		ioutil.WriteFile(ss[i], []byte("#error file removed!!!\n"), 0666)
		oldGenFiles[ss[i]] = true
	}
}

func genIncludeFiles() {
	ss := parseCMakeListsTxt("internal/libwebp-1.3.2/CMakeLists.txt", "WEBP_SRC_DIR", "*.c")
	muxSS, err := findFiles("internal/libwebp-1.3.2/src/mux", "*.c")
	if err != nil {
		log.Fatal(err)
	}
	ss = append(ss, muxSS...)
	for i := 0; i < len(ss); i++ {
		relpath := ss[i][23:] // drop `./`
		newname := "z_libwebp_" + strings.Replace(relpath, "/", "_", -1)

		ioutil.WriteFile(newname, []byte(fmt.Sprintf(
			`// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Auto Generated By 'go generate', DONOT EDIT!!!

// +build cgo

#include "%s"
`, relpath,
		)), 0666)

		delete(oldGenFiles, newname)
	}
}

func printOldGenFiles() {
	if len(oldGenFiles) == 0 {
		return
	}
	fmt.Printf("Removed Files:\n")
	for k, _ := range oldGenFiles {
		fmt.Printf("%s\n", k)
	}
	fmt.Printf("Total %d\n", len(oldGenFiles))
}

func parseCMakeListsTxt(filename, varname, ext string) (ss []string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	br := bufio.NewReader(bytes.NewReader(data))

	// find set($varname
	for {
		line, _, err := br.ReadLine()
		if err != nil {
			log.Fatal(err)
		}
		if strings.HasPrefix(string(line), "set("+varname) {
			break
		}
	}

	// read parse_makefile_am(${varname}, end with `)`
	prefix := fmt.Sprintf("parse_makefile_am(${%s}/", varname)
	baseDir := filepath.Join(filepath.Dir(filename), "src")
	for {
		line, _, err := br.ReadLine()
		if err != nil {
			log.Fatal(err)
		}
		s := string(line)
		if !strings.HasPrefix(s, prefix) {
			break
		}
		subdir := strings.Split(s, " ")[0][len(prefix):]
		dir := filepath.Join(baseDir, subdir)
		sf, err := findFiles(dir, ext)
		if err != nil {
			log.Fatal(err)
		}
		ss = append(ss, sf...)
	}
	return
}

func findFiles(dir, ext string) ([]string, error) {
	return filepath.Glob(fmt.Sprintf("%s/%s", dir, ext))
}
