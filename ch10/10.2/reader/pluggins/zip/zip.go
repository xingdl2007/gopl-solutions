// Copyright © 2017 xingdl2007@gmail.com
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
package zip

import (
	"archive/zip"
	"fmt"
	"gopl.io/gopl-solutions/ch10/10.2/reader"
)

func Reader(s string) error {
	// Open a zip archive for reading.
	r, err := zip.OpenReader(s)
	if err != nil {
		return err
	}
	defer r.Close()

	// Iterate through the files in the archive, printing all file names
	for _, f := range r.File {
		fmt.Printf("%s\n", f.Name)
	}
	return nil
}

func init() {
	reader.Register("zip", Reader)
}
