package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// 增加了比较功能。map + Process函数
func Process(name string, f func()) (duration time.Duration) {
	start := time.Now()
	f()
	dur := time.Since(start)
	fmt.Printf("%s : %v", name, dur)
	return dur
}
func lala() {
	var bestime time.Duration
	var bestname string
	m := make(map[string]time.Duration, 3)
	m["echov1"] = Process("echov1", Echov1)
	m["echov2"] = Process("echov2", Echov2)
	m["echov3"] = Process("echov3", Echov3)
    // 给map的value进行排序要会
	for name, time := range m {
		if bestime > time || bestname == "" {
			bestime = time
			bestname = name
		}
	}
	fmt.Println()
	fmt.Printf("besttime is %v,bestname is %s\n", bestime, bestname)

}
func Echov1() {
	sep := " "
	result := ""
	for _, args := range os.Args[1:] {
		args += sep
		result += args
	}
	fmt.Println(result)
}

func Echov2() {
	sep, result := "", ""
	for _, args := range os.Args[1:] {
		result += sep + args
		sep = " "
	}
	fmt.Println(result)
}

func Echov3() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
