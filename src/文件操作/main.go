package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("/Users/wallace/GolandProjects/go_base/src/文件操作/1.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("打开文件失败")
	}
	file1, err := os.OpenFile("/Users/wallace/GolandProjects/go_base/src/文件操作/1.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("打开文件失败")
	}

	fmt.Println(file.Fd())
	fmt.Println(file1.Fd())
}
