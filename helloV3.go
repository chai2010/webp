//go:build ignore
// +build ignore

package main

import (
	"image/jpeg"
	"log"
	"os"

	"github.com/cqzcqq/webp/convertor"
	"github.com/cqzcqq/webp/decoder"
)

func main() {
	file, err := os.Open("testdata/m4_q75.webp")
	if err != nil {
		log.Fatalln(err)
	}

	output, err := os.Create("testdata/m4_q75.jpg")
	if err != nil {
		log.Fatal(err)
	}
	//noinspection GoUnhandledErrorResult
	defer output.Close()

	img, err := convertor.Decode(file, &decoder.Options{})
	if err != nil {
		log.Fatalln(err)
	}

	if err = jpeg.Encode(output, img, &jpeg.Options{Quality: 75}); err != nil {
		log.Fatalln(err)
	}
}
