package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 在函数外，你可以声明一个 map，但不能直接使用 make() 进行初始化
type countInfo struct {
	totalCount int
	sigleInfo  map[string]int //string 是文件名  int是这个文件下面的这个line的数量
}

func main() {
	os.Args = []string{
		"cmd",
		"./xxx",
		"./xxx1",
		"./xxx2",
	}
	m := make(map[string]*countInfo, 10)
	err := openfile(os.Args, m)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}
	printResult(m)
}
func openfile(f []string, m map[string]*countInfo) error {
	for _, file := range f[1:] {
		fi, err := os.Open(file)
		if err != nil {
			return fmt.Errorf("open %s failed,err is %v ", file, err)
			//continue //如果第一个文件不存在，还想继续往下运行忽略错误就continue
		}
		defer fi.Close()
		if err := process(file, fi, m); err != nil {
			return err
		}
	}
	return nil
}

func process(fname string, f *os.File, m map[string]*countInfo) error {
	scan := bufio.NewScanner(f)
	for scan.Scan() {
		line := strings.TrimSpace(scan.Text())
		if m[line] == nil {
			m[line] = &countInfo{
				totalCount: 0,
				sigleInfo:  make(map[string]int, 10),
			}
		}
		m[line].sigleInfo[fname]++
		m[line].totalCount++
	}
	if err := scan.Err(); err != nil {
		return fmt.Errorf("scan failed,err is %v", err)
	}
	return nil
}
func printResult(m map[string]*countInfo) {
	for line, info := range m {
		s := make([]string,0 )
		for sigfile, count := range info.sigleInfo {
			s = append(s, fmt.Sprintf("%d times from %s", count, sigfile))
		}
		fmt.Printf("the %s totalcount is %d,(%s)\n", line, info.totalCount, strings.Join(s,", "))
	}
}
