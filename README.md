webp
=====

```
██╗    ██╗███████╗██████╗ ██████╗ 
██║    ██║██╔════╝██╔══██╗██╔══██╗
██║ █╗ ██║█████╗  ██████╔╝██████╔╝
██║███╗██║██╔══╝  ██╔══██╗██╔═══╝ 
╚███╔███╔╝███████╗██████╔╝██║     
 ╚══╝╚══╝ ╚══════╝╚═════╝ ╚═╝     
```

PkgDoc: [http://godoc.org/github.com/chai2010/webp](http://godoc.org/github.com/chai2010/webp)

Install CGO Version
===================

Install `GCC` or `MinGW` ([download here](http://tdm-gcc.tdragon.net/download)) at first,
and then run these commands:

1. Assure set the `CGO_ENABLED` environment variable to `1` to enable `CGO` (Default is enabled).
2. `go get github.com/chai2010/webp`
3. `go run hello.go`

Install Pure Go Version
=======================

1. Assure set the `CGO_ENABLED` environment variable to `0` to disable `CGO` (Default is enabled).
2. `go get github.com/chai2010/webp`
3. `go run hello.go`

Note: Only support `Decode` and `DecodeConfig`, `go test` will failed on some other api test.


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
		log.Println(err)
	}

	// GetInfo
	if width, height, _, err = webp.GetInfo(data); err != nil {
		log.Println(err)
	}
	fmt.Printf("width = %d, height = %d\n", width, height)

	// GetMetadata
	if metadata, err := webp.GetMetadata(data, "ICCP"); err != nil {
		fmt.Printf("Metadata: err = %v\n", err)
	} else {
		fmt.Printf("Metadata: %s\n", string(metadata))
	}

	// Decode webp
	m, err := webp.Decode(bytes.NewReader(data))
	if err != nil {
		log.Println(err)
	}

	// Encode lossless webp
	if err = webp.Encode(&buf, m, &webp.Options{Lossless: true}); err != nil {
		log.Println(err)
	}
	if err = ioutil.WriteFile("output.webp", buf.Bytes(), 0666); err != nil {
		log.Println(err)
	}
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

BUGS
====

Report bugs to <chaishushan@gmail.com>.

Thanks!
