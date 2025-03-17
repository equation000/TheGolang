package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

// 实现curl 获取url源代码功能
func lala() {
	getUrl()
}
func getUrl() {
	url := "www.baidu.com"
	if !strings.HasPrefix(url, "https://") {
		url = "https://" + url
	}
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("get url failed, err is ", err)
		return
	}
	defer resp.Body.Close()
	info, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read resp.body failed ,err is ", err)
		return
	}
	fmt.Println(string(info), resp.Status)
}
//补充知识 return 是用来结束main函数的执行的。return上面的语句都可以执行。到他这就退出。
//os.Exit 和 return 在 defer 和 goroutine 上的区别！ io.EOF是什么
// os.Exit 和 return 都是结束程序，但也有着区别
// return 
// 适用于正常退出程序，且需要执行清理工作（如关闭文件、释放资源）。并且返回退出吗0
// return 等待defer ,不等待groutune
// os.Exit
// 适用于需要立即终止程序的场景，例如错误处理中直接退出并返回特定状态码。返回退出吗1
// 什么都不等待直接结束。
// 见例子 

