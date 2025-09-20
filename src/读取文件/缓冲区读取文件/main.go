package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main1() {

	ifile, err := os.Open("fileLog.go")
	if err != nil {
		fmt.Println("open file %s err", err)
		return
	}
	defer ifile.Close()
	iReder := bufio.NewReader(ifile)
	for {
		iString, err := iReder.ReadString('\n')
		if err == io.EOF {
			return
		}
		fmt.Println(iString)
	}

}

func main() {

	ifile, err := os.ReadFile("/Users/wallace/GolandProjects/go_base/src/读取文件/缓冲区读取文件/fileLog.go")
	if err != nil {
		fmt.Println("读取文件失败", err)
	}
	fmt.Println(string((ifile)))
}
