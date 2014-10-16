webp
=====

PkgDoc: [http://godoc.org/github.com/chai2010/webp](http://godoc.org/github.com/chai2010/webp)

Install
=======

Install `GCC` or `MinGW` ([download here](http://tdm-gcc.tdragon.net/download)) at first,
and then run these commands:

1. `go get github.com/chai2010/webp`
2. `go run hello.go`

Example
=======

This is a simple example:

```Go
package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/chai2010/webp"
)

func main() {
	var buf bytes.Buffer
	var width, height int
	var data []byte
	var err error

	// Load file data
	if data, err = ioutil.ReadFile("./testdata/1_webp_ll.webp"); err != nil {
		log.Fatal(err)
	}

	// GetInfo
	if width, height, _, err = webp.GetInfo(data); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("width = %d, height = %d\n", width, height)

	// Decode webp
	m, err := webp.Decode(bytes.NewReader(data))
	if err != nil {
		log.Fatal(err)
	}

	// Encode lossless webp
	if err = webp.Encode(&buf, m, &webp.Options{Lossless: true}); err != nil {
		log.Fatal(err)
	}
	if err = ioutil.WriteFile("output.webp", buf.Bytes(), 0666); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Save output.webp ok\n")
}
```

Decode and Encode as RGB format:

```Go
m, err := webp.DecodeRGB(data)
if err != nil {
	log.Fatal(err)
}

data, err := webp.EncodeRGB(m)
if err != nil {
	log.Fatal(err)
}
```

Use `m.Pix` form `webp.Image` interface:

```Go
m, _ := webp.DecodeRGB(data)
rgb := m.(webp.Image)
b := rgb.Bounds()

for y := b.Min.Y; y < b.Max.Y; y++ {
	off := (y-b.Min.Y)*rgb.Stride()
	pix := rgb.Pix()[off:][:rgb.Stride()]
	for i := 0; i < len(pix); i += 3 {
		R := pix[i+0]
		G := pix[i+1]
		B := pix[i+2]
		...
	}
}
```

BUGS
====

Report bugs to <chaishushan@gmail.com>.

Thanks!
