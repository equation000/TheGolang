package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)
//以下是获取输入的方式
// 从命令行输入
func fromCommandArgs() {
	os.Args = []string{"cmd", "world", "nihao", "nihao"}
	for i := 0; i < len(os.Args[1:]); i++ {
		fmt.Printf("%s\n", os.Args[i])
	}
}

// 从环境变量输入
func fromEnv() {
	//printenv can print envirment vars
	envVar := os.Getenv("GOPATH")
	fmt.Printf("var is %s\n ", envVar)
}

// 从互联网获取信息作为输入
func fromInternet() {
	resp, err := http.Get("https://www.baidu.com")
	if err != nil {
		fmt.Println("Get info failed from Internet", err)
	}
	defer resp.Body.Close()
	info, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read internet info failed,err is", err)
	}
	fmt.Println("result is", string(info))

}

// 从本地文件获取输入
func fromfile() {
	info, err := os.ReadFile("./xxx")
	if err != nil {
		fmt.Println("get info faliled,err is ", err)
	}
	fmt.Println(string(info))
}
// 从标准输入的到数据
func fromStdin() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("Enter data:")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("get info failed from standord input, err is", err)
			break
		}
		input = strings.TrimSpace(input)
		if input == "exit" {
			fmt.Println("exit prgram")
			break
		}
		fmt.Println(input)
	}
}

func lalaa() {
	fromStdin()
}
