// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ingore

package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	flagRevert = flag.Bool("revert", false, "revert all changes")
)

var convertMap = [][2]string{
	[2]string{
		`"github.com/chai2010/assert"`,
		`"github.com/chai2010/webp/internal/assert"`,
	},
}

func main() {
	flag.Parse()
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal("filepath.Walk: ", err)
			return err
		}
		if info.IsDir() {
			return nil
		}
		if strings.HasSuffix(path, "gen_helper.go") {
			return nil
		}
		if strings.HasSuffix(path, "hello.go") {
			return nil
		}
		if strings.HasSuffix(path, ".go") {
			fixImportPath(path)
		}
		return nil
	})
}

func fixImportPath(filename string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal("ioutil.ReadFile: ", err)
	}

	for _, v := range convertMap {
		oldPath, newPath := v[0], v[1]
		if !*flagRevert {
			data = bytes.Replace(data, []byte(oldPath), []byte(newPath), -1)
		} else {
			data = bytes.Replace(data, []byte(newPath), []byte(oldPath), -1)
		}
	}

	if err = ioutil.WriteFile(filename, data, 0666); err != nil {
		log.Fatal("ioutil.WriteFile: ", err)
	}

	if !*flagRevert {
		fmt.Printf("convert %s ok\n", filename)
	} else {
		fmt.Printf("revert %s ok\n", filename)
	}
}
