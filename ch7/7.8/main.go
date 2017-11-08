// Copyright © 2017 xingdl2007@gmail.com
// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 187.

// Sorting sorts a music playlist into a variety of orders.
package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

//!+main
type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

//!-main

//!+printTracks
func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}

type multier struct {
	t         []*Track
	primary   string
	secondary string
	third     string
}

func (x *multier) Len() int      { return len(x.t) }
func (x *multier) Swap(i, j int) { x.t[i], x.t[j] = x.t[j], x.t[i] }

func (x *multier) Less(i, j int) bool {
	key := x.primary
	for k := 0; k < 3; k++ {
		switch key {
		case "Title":
			if x.t[i].Title != x.t[j].Title {
				return x.t[i].Title < x.t[j].Title
			}
		case "Year":
			if x.t[i].Year != x.t[j].Year {
				return x.t[i].Year < x.t[j].Year
			}
		case "Length":
			if x.t[i].Length != x.t[j].Length {
				return x.t[i].Length < x.t[j].Length
			}
		}
		if k == 0 {
			key = x.secondary
		} else if k == 1 {
			key = x.third
		}
	}
	return false
}

// update primary sorting key
func setPrimary(x *multier, p string) {
	x.primary, x.secondary, x.third = p, x.primary, x.secondary
}

// if x is *multiple type, then update ordering keys
func SetPrimary(x sort.Interface, p string) {
	switch x := x.(type) {
	case *multier:
		setPrimary(x, p)
	}
}

// return a new multier
func NewMultier(t []*Track, p, s, th string) sort.Interface {
	return &multier{
		t:         t,
		primary:   p,
		secondary: s,
		third:     th,
	}
}

func main() {
	fmt.Println("\nMultier:")
	multi := NewMultier(tracks, "Title", "", "")
	sort.Sort(multi)
	printTracks(tracks)

	// set primary key
	fmt.Println()
	SetPrimary(multi, "Year")
	sort.Sort(multi)
	printTracks(tracks)

	fmt.Println()
	SetPrimary(multi, "Length")
	sort.Sort(multi)
	printTracks(tracks)

	fmt.Println()
	SetPrimary(multi, "Title")
	sort.Sort(multi)
	printTracks(tracks)
}
