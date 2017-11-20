// Copyright © 2017 xingdl2007@gmail.com
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
package reader

import (
	"strings"
	"fmt"
)

// register for zip/tar file
func Register(n string, f func(s string) error) {
	format := format{name: n, readFunc: f}
	formats = append(formats, format)
}

type format struct {
	name     string
	readFunc func(s string) error
}

var formats []format

func ArchiveReader(s string) error {
	period := strings.LastIndex(s, ".")
	if period == -1 {
		return fmt.Errorf("unknowns archive format")
	}
	suffix := s[period+1:]
	for _, format := range formats {
		if suffix == format.name {
			return format.readFunc(s)
		}
	}
	return fmt.Errorf("unsupported archive format")
}
