package csp

import (
	"flag"
)

var (
	cspOn      bool
	debug      bool
	outputFile string
)

// For compatibility reason, all options provided by this library are only
// in the long option form with the 'csp' prefix
func init() {
	flag.BoolVar(&cspOn, "csp", false, "Turn on CSP scsripts output mode.")
	flag.BoolVar(&debug, "csp-debug", false, "Output CSP scripts output debug info to terminal.")
	flag.StringVar(&outputFile, "csp-outputfile", "main.csp", "Specify output file name, defualt name: main.csp.")

	flag.Parse()
}

// func helpInfo() {

// 	fmt.Fprintf(os.Stderr, `Golang to PAT-style CSP scripts library

// Options:
// `)
// 	flag.PrintDefaults()
// }
