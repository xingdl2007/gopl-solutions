// Copyright © 2017 xingdl2007@gmail.com
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"sort"
)

type Package struct {
	ImportPath string   // import path of package in dir
	Name       string   // package name
	Deps       []string // all (recursively) imported dependencies
}

func main() {
	var err error
	// first validate parameter package
	if len(os.Args) < 2 {
		log.Fatalf("useage: main package")
	}
	key := os.Args[1]
	cmd := exec.Command("go", "list", key)
	if _, err = cmd.Output(); err != nil {
		log.Fatalf("package %s invalid: %v", key, err)
	}

	// list all packages in workspaces, first in text format
	cmd = exec.Command("go", "list", "-json", "...")
	if cmd == nil {
		log.Fatalf("can't run go list")
	}

	var output []byte
	if output, err = cmd.Output(); err != nil {
		log.Fatal(err)
	}

	var stack []byte
	var buf bytes.Buffer
	for _, b := range output {
		switch b {
		case '{':
			stack = append(stack, b)
		case '}':
			stack = stack[0:len(stack)-1]
		}
		// delete all newline and space
		buf.WriteByte(b)
		if b == '}' && len(stack) == 0 {
			// unmarshal json
			var info Package
			if err = json.Unmarshal(buf.Bytes(), &info); err != nil {
				log.Fatal(err)
			}
			if sort.SearchStrings(info.Deps, key) == len(info.Deps) {
				fmt.Println(info.ImportPath)
			}
			buf.Truncate(0)
		}
	}
}
