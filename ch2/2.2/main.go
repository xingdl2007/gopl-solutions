// Copyright © xingdl2007@gmail.com
// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 43.
//!+

// uc converts its numeric argument to Celsius and Fahrenheit, to Feet and Meter.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"gopl.io/ch2/tempconv"
)

type Feet float64
type Meter float64

func (f Feet) String() string  { return fmt.Sprintf("%g(ft)", f) }
func (m Meter) String() string { return fmt.Sprintf("%g(m)", m) }

func FToM(f Feet) Meter { return Meter(f * 0.3048) }
func MToF(m Meter) Feet { return Feet(m * 3.28083) }

func uc(t float64) {
	// temperature
	fmt.Println("Temperature Conversion:")
	f := tempconv.Fahrenheit(t)
	c := tempconv.Celsius(t)
	fmt.Printf("%s = %s, %s = %s\n\n",
		f, tempconv.FToC(f), c, tempconv.CToF(c))

	// length
	fmt.Println("Length Conversion:")
	lf := Feet(t)
	lm := Meter(t)
	fmt.Printf("%s = %s, %s = %s\n\n",
		lf, FToM(lf), lm, MToF(lm))
}

func main() {
	if len(os.Args[1:]) > 0 {
		for _, arg := range os.Args[1:] {
			t, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "uc: %v\n", err)
				os.Exit(1)
			}
			uc(t)
		}
	} else {
		for {
			input := bufio.NewReader(os.Stdin)
			fmt.Fprintf(os.Stdout, "=> ")
			s, err := input.ReadString('\n')
			if err != nil {
				fmt.Fprintf(os.Stderr, "uc: %v\n", err)
				os.Exit(1)
			}
			// strip `\n`
			t, err := strconv.ParseFloat(s[:len(s)-1], 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "uc: %v\n", err)
				os.Exit(1)
			}
			uc(t)
		}
	}
}

//!-
