// Copyright © 2017 xingdl2007@gmail.com
// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 112.
//!+

// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"fmt"
	"gopl.io/ch4/github"
	"log"
	"os"
	"time"
)

//!+

const aMonth = 31 * 24
const aYear = 31 * 365 * 24

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours())
}

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)

	var inMonth, inYear, outYear []*github.Issue
	for _, item := range result.Items {
		// less a month/a year, or more than a year old
		if daysAgo(item.CreatedAt) < aYear {
			inMonth = append(inMonth, item)
		} else if daysAgo(item.CreatedAt) < aMonth {
			inYear = append(inYear, item)
		} else {
			outYear = append(outYear, item)
		}
	}

	fmt.Printf("%d issues less than a month:\n", len(inMonth))
	for _, item := range inMonth {
		fmt.Printf("#%-5d %9.9s %v %.55s\n",
			item.Number, item.User.Login, item.CreatedAt, item.Title)
	}

	fmt.Printf("\n%d issues less than a year:\n", len(inYear))
	for _, item := range inYear {
		fmt.Printf("#%-5d %9.9s %v %.55s\n",
			item.Number, item.User.Login, item.CreatedAt, item.Title)
	}

	fmt.Printf("\n%d issues more than a month:\n", len(outYear))
	for _, item := range outYear {
		fmt.Printf("#%-5d %9.9s %v %.55s\n",
			item.Number, item.User.Login, item.CreatedAt, item.Title)
	}
}

//!-

/*
//!+textoutput
$ go build gopl.io/ch4/issues
$ ./issues repo:golang/go is:open json decoder
13 issues:
#5680    eaigner encoding/json: set key converter on en/decoder
#6050  gopherbot encoding/json: provide tokenizer
#8658  gopherbot encoding/json: use bufio
#8462  kortschak encoding/json: UnmarshalText confuses json.Unmarshal
#5901        rsc encoding/json: allow override type marshaling
#9812  klauspost encoding/json: string tag not symmetric
#7872  extempora encoding/json: Encoder internally buffers full output
#9650    cespare encoding/json: Decoding gives errPhase when unmarshalin
#6716  gopherbot encoding/json: include field name in unmarshal error me
#6901  lukescott encoding/json, encoding/xml: option to treat unknown fi
#6384    joeshaw encoding/json: encode precise floating point integers u
#6647    btracey x/tools/cmd/godoc: display type kind of each named type
#4237  gjemiller encoding/base64: URLEncoding padding is optional
//!-textoutput
*/
