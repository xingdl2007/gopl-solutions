// Copyright © 2017 xingdl2007@gmail.com
// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 61.
//!+

// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"math/cmplx"

	"time"
	"fmt"
)

// N represent the number of goroutines
const N = 128

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	start := time.Now()
	// partition jobs
	size := height / N
	var done = make(chan struct{})
	for i := 0; i < N; i++ {
		go func(i int) {
			for py := size * i; py < size*(i+1); py++ {
				y := float64(py)/height*(ymax-ymin) + ymin
				for px := 0; px < width; px++ {
					x := float64(px)/width*(xmax-xmin) + xmin
					z := complex(x, y)

					// Image point (px, py) represents complex value z.
					img.Set(px, py, mandelbrot(z))
				}
			}
			done <- struct{}{}
		}(i)
	}

	// drain done channel
	for i := 0; i < N; i++ {
		<-done
	}

	// skip output
	//png.Encode(os.Stdout, img) // NOTE: ignoring errors
	fmt.Printf("elapsed %fms\n", time.Since(start).Seconds()*1000)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
