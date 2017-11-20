// Copyright © 2017 xingdl2007@gmail.com
// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 287.

//!+main

// The jpeg command reads a PNG image from the standard input
// and writes it as a JPEG image to the standard output.
package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"os"
	"image/png"
	"flag"
	"log"
)

var format = flag.String("format", "jpg", "output format(jpg/png)")

func main() {
	flag.Parse()
	for _, p := range flag.Args() {
		file, err := os.Open(p)
		if err != nil {
			log.Printf("%s open failed: %v\n", file, err)
			continue
		}

		img, kind, err := image.Decode(file)
		if err != nil {
			log.Printf("%s file decode err: %v", p, err)
			continue
		}
		fmt.Fprintln(os.Stderr, "Input format =", kind)
		if kind == *format {
			log.Printf("%s is already %s format", p, kind)
			continue
		}

		outfilename := p + "." + *format
		outfile, err := os.Create(outfilename)
		if err != nil {
			log.Printf("create %s failed: %v", outfilename, err)
			continue
		}
		if *format == "jpg" {
			if err := toJPEG(img, outfile); err != nil {
				fmt.Fprintf(os.Stderr, "jpeg: %v\n", err)
				os.Exit(1)
			}
		} else if *format == "png" {
			if err := toPNG(img, outfile); err != nil {
				fmt.Fprintf(os.Stderr, "jpeg: %v\n", err)
				os.Exit(1)
			}
		}
		file.Close()
		outfile.Close()
	}
}

func toJPEG(img image.Image, out io.Writer) error {
	return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
}

func toPNG(img image.Image, out io.Writer) error {
	return png.Encode(out, img)
}

//!-main

/*
//!+with
$ go build gopl.io/ch3/mandelbrot
$ go build gopl.io/ch10/jpeg
$ ./mandelbrot | ./jpeg >mandelbrot.jpg
Input format = png
//!-with

//!+without
$ go build gopl.io/ch10/jpeg
$ ./mandelbrot | ./jpeg >mandelbrot.jpg
jpeg: image: unknown format
//!-without
*/
