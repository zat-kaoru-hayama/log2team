package main

import (
	"fmt"
	"os"

	"github.com/zat-kaoru-hayama/log2team/lib"
)

func main() {
	args := os.Args
	if len(args) < 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s URL TITLE\n", os.Args[0])
		return
	}
	if err := log2team.Send(args[1], args[2], os.Stdin); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
