webp
=====

PkgDoc: [http://godoc.org/github.com/chai2010/webp](http://godoc.org/github.com/chai2010/webp)

Report bugs to <chaishushan@gmail.com>.


#### Sample usage:

    package main

    import (
        "github.com/chai2010/webp"
        "image"
        _ "image/jpeg"
        _ "image/png"
        "log"
        "os"
    )

    func main() {
        i, o := os.Args[1], os.Args[2]
        f, err := os.Open(i)
        if err != nil {
            log.Fatal("Error unable to open a file: ", err)
        }
        defer f.Close()
        img, _, err := image.Decode(f)
        if err != nil {
            log.Fatal("Error unable to decode a file: ", err)
        }
        w, err := os.Create(o)
        if err != nil {
            log.Fatal("Error unable to create a file: ", err)
        }
        defer w.Close()
        err = webp.Encode(w, img, &webp.Options{true, 90})
        if err != nil {
            log.Fatal("Error unable to encode a file: ", err)
        }
    }
    

