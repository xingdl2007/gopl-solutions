// Copyright © 2017 xingdl2007@gmail.com
// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 187.

// Sorting sorts a music playlist into a variety of orders.
package main

import (
	"sort"
	"time"
	"html/template"
	"io"
	"net/http"
	"log"
	"fmt"
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
// ref: https://stackoverflow.com/questions/25824095/order-by-clicking-table-header
var trackTable = template.Must(template.New("Track").Parse(`
<h1> Tracks </h1>
<table>
<tr style='text-align: left'>
	<th onclick="submitform('Title')">Title
        <form action="" name="Title" method="post">
            <input type="hidden" name="orderby" value="Title"/>
        </form>
    </th>
	<th>Artist
        <form action="" name="Artist" method="post">
            <input type="hidden" name="orderby" value="Artist"/>
        </form>
    </th>
	<th>Album
        <form action="" name="Album" method="post">
            <input type="hidden" name="orderby" value="Album"/>
        </form>
    </th>
	<th onclick="submitform('Year')">Year
        <form action="" name="Year" method="post">
            <input type="hidden" name="orderby" value="Year"/>
        </form>
    </th>
	<th onclick="submitform('Length')">Length
        <form action="" name="Length" method="post">
            <input type="hidden" name="orderby" value="Length"/>
        </form>
    </th>
</tr>
{{range .T}}
<tr>
	<td>{{.Title}}</td>
	<td>{{.Artist}}</td>
	<td>{{.Album}}</td>
	<td>{{.Year}}</td>
	<td>{{.Length}}</td>
</tr>
{{end}}
</table>

<script>
function submitform(formname) {
    document[formname].submit();
}
</script>
`))

//!+printTracks
func printTracks(w io.Writer, x sort.Interface) {
	switch x := x.(type) {
	case *multier:
		trackTable.Execute(w, x)
	}
}

type multier struct {
	T         []*Track // exported
	primary   string
	secondary string
	third     string
}

func (x *multier) Len() int      { return len(x.T) }
func (x *multier) Swap(i, j int) { x.T[i], x.T[j] = x.T[j], x.T[i] }

func (x *multier) Less(i, j int) bool {
	key := x.primary
	for k := 0; k < 3; k++ {
		switch key {
		case "Title":
			if x.T[i].Title != x.T[j].Title {
				return x.T[i].Title < x.T[j].Title
			}
		case "Year":
			if x.T[i].Year != x.T[j].Year {
				return x.T[i].Year < x.T[j].Year
			}
		case "Length":
			if x.T[i].Length != x.T[j].Length {
				return x.T[i].Length < x.T[j].Length
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
		T:         t,
		primary:   p,
		secondary: s,
		third:     th,
	}
}

func main() {
	// default sort by "Title"
	multi := NewMultier(tracks, "Title", "", "")
	sort.Sort(multi)

	// start a simple server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			fmt.Printf("ParseForm: %v\n", err)
		}
		for k, v := range r.Form {
			if k == "orderby" {
				SetPrimary(multi, v[0])
			}
		}
		sort.Sort(multi)
		printTracks(w, multi)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
