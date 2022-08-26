package main

import (
	"apitest/cmd"
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	fmt.Println()
	if cmd.Help {
		flag.Usage()
	}
}
