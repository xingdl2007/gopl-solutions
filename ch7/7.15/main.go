// Copyright © 2017 xingdl2007@gmail.com
// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package main

import (
	"fmt"
	"gopl.io/gopl-solutions/ch7/7.14/eval"
	"bufio"
	"os"
	"strings"
	"strconv"
)

//!+parseAndCheck
func parseAndCheck(s string) (eval.Expr, map[eval.Var]bool, error) {
	if s == "" {
		return nil, nil, fmt.Errorf("empty expression")
	}
	expr, err := eval.Parse(s)
	if err != nil {
		return nil, nil, err
	}
	vars := make(map[eval.Var]bool)
	if err := expr.Check(vars); err != nil {
		return nil, nil, err
	}
	return expr, vars, nil
}

func main() {
	var env = make(eval.Env)
	input := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stdout, "=> ")
		s, _ := input.ReadString('\n')
		if s = strings.TrimSpace(s); s == "" {
			continue
		}

		expr, vars, err := parseAndCheck(s)
		if err != nil {
			fmt.Printf("parse %s failed: %v\n", s, err)
			continue
		}

		for k, _ := range vars {
			if _, ok := env[k]; !ok {
				fmt.Fprintf(os.Stdout, "=> %v is ? ", k)
				var n float64
				for {
					d, _ := input.ReadString('\n')
					d = strings.TrimSpace(d)

					n, err = strconv.ParseFloat(d, 64)
					if err != nil {
						fmt.Printf("=> %g is invalid\n", n)
						continue
					}
					break
				}
				env[k] = n
			}
		}
		fmt.Printf("%v = %v\n", expr, expr.Eval(env))
	}
}
