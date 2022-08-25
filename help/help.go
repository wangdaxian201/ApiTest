package help

import (
	"flag"
	"fmt"
	"time"
)

func Init() {
	// 获取命令行参数
	var name string
	var age int
	var married bool
	var delay time.Duration
	var concurrent int
	flag.IntVar(&concurrent, "s", 2, "并发数")

	//解析命令行参数
	flag.Parse()
	fmt.Printf("并发数:%v\n", concurrent)
	fmt.Println(name, age, married, delay)
	//返回命令行参数后的其他参数
	fmt.Println(flag.Args())
	//返回命令行参数后的其他参数个数
	fmt.Println(flag.NArg())
	//返回使用的命令行参数个数
	fmt.Println(flag.NFlag())
}
