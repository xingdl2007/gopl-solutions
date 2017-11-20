// Copyright © 2017 xingdl2007@gmail.com
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
package main

import (
	"log"
	_ "gopl.io/gopl-solutions/ch10/10.2/reader/pluggins/zip"
	_ "gopl.io/gopl-solutions/ch10/10.2/reader/pluggins/tar"

	"gopl.io/gopl-solutions/ch10/10.2/reader"
)

func main() {
	// read readme.zip/tar

	err := reader.ArchiveReader("readme.zip")
	if err != nil {
		log.Fatal(err)
	}

	err = reader.ArchiveReader("readme.tar")
	if err != nil {
		log.Fatal(err)
	}
}
