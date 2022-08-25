package main

import (
	"ApiTest/Usage"
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	fmt.Println(Usage.Option)
	if Usage.Help {
		flag.Usage()
	}
}
