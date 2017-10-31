// Copyright © 2017 xingdl2007@gmail.com
// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 58.
//!+

// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

type CornerType int

const (
	middle CornerType = 0 // not the peak or valley of surface
	peak   CornerType = 1 // the peak of surface corner
	valley CornerType = 2 // the valley of surface corner
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, ct, err := corner(i+1, j)
			if err != nil {
				continue
			}
			bx, by, ct1, err := corner(i, j)
			if err != nil {
				continue
			}
			cx, cy, ct3, err := corner(i, j+1)
			if err != nil {
				continue
			}
			dx, dy, ct2, err := corner(i+1, j+1)
			if err != nil {
				continue
			}

			var color string
			if ct == peak || ct1 == peak || ct2 == peak || ct3 == peak {
				color = "#f00"
			} else if ct == valley || ct1 == valley || ct2 == valley || ct3 == valley {
				color = "#00f"
			} else {
				// same as default
				color = "grey"
			}

			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' style='stroke: %s'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, color)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64, CornerType, error) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z, ct := f(x, y)

	// check
	if math.IsInf(z, 0) || math.IsNaN(z) {
		return 0, 0, 0, fmt.Errorf("invalid value")
	}

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, ct, nil
}

func f(x, y float64) (float64, CornerType) {
	d := math.Hypot(x, y) // distance from (0,0)
	ct := middle

	// f(x) = sin(x)/x, f'(x) = (x*cos(x)-sin(x))/x^2
	// f'(x) = 0 ==> x = tan(x), peak or vally
	// if f''(x) > 0, vally
	// if f''(x) < 0, peak
	// f''(x) = {2(sin(x)-x*cos(x)) - x*x*sin(x)}/x*x*x
	if math.Abs(d-math.Tan(d)) < 3 {
		ct = peak

		if 2*(math.Sin(d)-d*math.Cos(d))-d*d*math.Sin(d) > 0 {
			ct = valley
		}
	}

	return math.Sin(d) / d, ct
}

//!-
