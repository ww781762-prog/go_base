package main

import (
	"flag"
	"fmt"
	"os"
)

func ListDir(dir string, deep int) {
	ds, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("获取目录失败:", err)
	}
	if deep == 1 {
		fmt.Printf("!----%s\n", dir)
	}
	for _, d := range ds {
		if d.IsDir() {
			for i := 0; i < deep; i++ {
				fmt.Printf("     |")
			}
			fmt.Println("------", d.Name())
			ListDir(dir+"/"+d.Name(), deep+1)
		} else {
			for i := 0; i < deep; i++ {
				fmt.Printf("     |")
			}
			fmt.Println("文件:", d.Name())
		}
	}
}

func main() {
	flag.Parse()
	var dir string
	if flag.NArg() == 0 {
		dir = "."
	} else {
		dir = flag.Arg(0)
	}
	ListDir(dir, 1)
}
