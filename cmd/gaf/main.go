package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Xx0w0wxX/gaf"
)

func main() {
	err := gaf.Run(os.Args[1:], os.Stdout, os.Stderr)
	if err != nil && err != flag.ErrHelp {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}