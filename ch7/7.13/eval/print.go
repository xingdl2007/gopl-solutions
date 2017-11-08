// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package eval

import (
	"bytes"
	"fmt"
)

// Format formats an expression as a string.
// It does not attempt to remove unnecessary parens.
func Format(e Expr) string {
	var buf bytes.Buffer
	write(&buf, e)
	return buf.String()
}

func write(buf *bytes.Buffer, e Expr) {
	switch e := e.(type) {
	case literal:
		fmt.Fprintf(buf, "%g", e)

	case Var:
		fmt.Fprintf(buf, "%s", e)

	case unary:
		fmt.Fprintf(buf, "(%c", e.op)
		write(buf, e.x)
		buf.WriteByte(')')

	case binary:
		buf.WriteByte('(')
		write(buf, e.x)
		fmt.Fprintf(buf, " %c ", e.op)
		write(buf, e.y)
		buf.WriteByte(')')

	case call:
		fmt.Fprintf(buf, "%s(", e.fn)
		for i, arg := range e.args {
			if i > 0 {
				buf.WriteString(", ")
			}
			write(buf, arg)
		}
		buf.WriteByte(')')

	default:
		panic(fmt.Sprintf("unknown Expr: %T", e))
	}
}

func (v Var) String() string {
	return string(v)
}

func (l literal) String() string {
	return fmt.Sprintf("%g", l)
}

func (u unary) String() string {
	return fmt.Sprintf("(%c%v)", u.op, u.x)
}

func (b binary) String() string {
	return fmt.Sprintf("(%v %c %v)", b.x, b.op, b.y)
}

func (c call) String() string {
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "%s(", c.fn)
	for i, arg := range c.args {
		if i > 0 {
			buf.WriteString(", ")
		}
		write(buf, arg)
	}
	buf.WriteByte(')')
	return buf.String()
}

func (s sum) String() string {
	buf := new(bytes.Buffer)
	fmt.Fprint(buf, "{")
	for i, arg := range s.args {
		if i > 0 {
			buf.WriteString(", ")
		}
		write(buf, arg)
	}
	buf.WriteByte('}')
	return buf.String()
}
