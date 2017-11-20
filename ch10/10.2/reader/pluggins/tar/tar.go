// Copyright © 2017 xingdl2007@gmail.com
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
package tar

import (
	"archive/tar"
	"fmt"
	"io"
	"os"

	"gopl.io/gopl-solutions/ch10/10.2/reader"
)

func Reader(s string) error {
	// Open the tar archive for reading.
	file, err := os.Open(s)
	if err != nil {
		return err
	}

	tr := tar.NewReader(file)
	// Iterate through the files in the archive.
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			// end of tar archive
			break
		}
		if err != nil {
			return err
		}
		fmt.Printf("%s\n", hdr.Name)
	}
	file.Close()
	return nil
}

func init() {
	reader.Register("tar", Reader)
}
