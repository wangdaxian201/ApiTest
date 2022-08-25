package Usage

import (
	"flag"
	"fmt"
	"os"
)

var (
	Help      bool
	Url       string
	Option    string
	Header    string
	UserAgent string
)

func init() {
	flag.BoolVar(&Help, "h", false, "this `help`")
	flag.StringVar(&Url, "u", "", "`Url`")
	flag.StringVar(&Option, "o", "", "`Option` 是你要请求的方式 ")
	flag.StringVar(&Header, "header", "", "`header` 请求头")
	flag.Usage = usage
}

func usage() {
	fmt.Fprintf(os.Stderr, `ApiTest version: ApiTest/0.1.0
Usage: ApiTest [options...] <url>
Options:
`)
	flag.PrintDefaults()
}
