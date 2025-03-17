package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 写程序考虑这几个方面  处理输入(数据源在哪里，拿什么存储)  得到输入处理逻辑就需要变量进行存储 判断错误 根据条件进行结果输出（也可以根据结果来考虑使用什么变量）
//对文件做拷贝、打印、搜索、排序、统计的go程序 待补充

// 统计文件和输入流的数据的重复数
func countRepeat() {
    m := make(map[string]int, 10)
    input := bufio.NewScanner(os.Stdin)
    if input.Err() != nil {
        fmt.Println("get data failed")
    }
    for input.Scan() {
        if strings.TrimSpace(input.Text()) == "exit" {
            break
        }
        m[input.Text()]++
    }
    for line, num := range m {
        if num > 1 {
            fmt.Printf("this %s,appear %d times\n", line, num)
        }
    }
}
