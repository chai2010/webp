webp
=====

PkgDoc: [http://godoc.org/github.com/chai2010/webp](http://godoc.org/github.com/chai2010/webp)

Report bugs to <chaishushan@gmail.com>.


### Sample code

    package main
 
    import (
	      "image"
        _ "image/jpeg"
	      _ "image/png"
	      "log"
	      "os"
 
	      "github.com/chai2010/webp"
    )
 
    func main() {

	      i, o := os.Args[1], os.Args[2]

	      f, err := os.Open(i)
	      defer f.Close()
	      if err != nil {
		        log.Fatal("Error unable to open a file: ", err)
	      }

	      img, _, err := image.Decode(f)
	      if err != nil {
		        log.Fatal("Error unable to decode a file: ", err)
	      }

	      w, err := os.Create(o)
	      defer w.Close()
	      if err != nil {
		        log.Fatal("Error unable to create a file: ", err)
	      }

	      webp.Encode(w, img, &webp.Options{true, 90})
    }
