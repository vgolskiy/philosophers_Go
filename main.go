package main

import (
	"fmt"
	"os"
)

func main() {
	var status Status
	if len(os.Args) < 5 || len(os.Args) > 6 {
		fmt.Fprint(os.Stderr, "Wrong args quantity")
		os.Exit(1)
	}
	initStatus(os.Args, &status)
	parallelize(&status)
	return
}
