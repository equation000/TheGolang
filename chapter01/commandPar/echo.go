package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// 从命令行输入


func process() {
	start := time.Now()
	echov1()
	fmt.Println("echov1 use time", time.Since(start))
	start = time.Now()
	echov2()
	fmt.Println("echov2 use time", time.Since(start))
	start = time.Now()
	echov3()
	fmt.Println("echov3 use time", time.Since(start))
}

func echov1() {
	sep := " "
	result := ""
	for _, args := range os.Args[1:] {
		args += sep
		result += args
	}
	fmt.Println(result)
}

//上面的这个echo实现不好  两点原因
//字符串不会变的只能拼接，拼接就会生成新的变量，导致大量的内存分配和赋值不好
//上面的结果会多一个空格，不够优雅

func echov2() {
	sep, result := "", ""
	for _, args := range os.Args[1:] {
		result += sep + args
		sep = " "
	}
	fmt.Println(result)
}

// 使用strings.join来实现echov3
func echov3() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
