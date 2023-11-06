package csp

import "fmt"

func debugPrintln(args ...interface{}) {
	if debug {
		fmt.Println(args...)
	}
}
