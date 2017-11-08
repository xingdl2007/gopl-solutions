// Copyright © 2017 xingdl2007@gmail.com
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
package main

import "fmt"

/*
// String return values of tree in a list, if no value at all, return `[]`
// otherwise return `[1 2 3...]`, values are separated by space
func (root *tree) String() string {
	var buf bytes.Buffer
	buf.WriteByte('[')

	for _, v := range appendValues(nil, root) {
		buf.WriteString(strconv.Itoa(v))
		buf.WriteByte(' ')
	}

	// delete last ' ' in case
	if buf.Len() > 1 {
		buf.Truncate(buf.Len() - 1)
	}

	buf.WriteByte(']')
	return buf.String()
}
*/
func main() {
	fmt.Println("See ch4/treesort/main.go")
}
