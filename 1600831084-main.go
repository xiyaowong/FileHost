// Package main provides ...
package main

import (
	"fmt"
	"os"
)

func main() {
	InitConfig()
	var filePath string
	if len(os.Args) == 1 {
		fmt.Printf("输入文件路径: ")
		fmt.Scanf("%s\n", &filePath)
	} else {
		filePath = os.Args[1]
	}
	if !PathExists(filePath) {
		fmt.Println("文件不存在")
		os.Exit(1)
	}
	url, err := Upload(filePath)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(url)
	}
	fmt.Scanln("回车键退出...")
}
