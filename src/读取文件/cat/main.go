package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

func cat(r *bufio.Reader) int {
	var i int
	for {
		i += 1
		buf, err := r.ReadString('\n')
		if err == io.EOF {
			i = 0
			break
		}
		fmt.Fprintf(os.Stdout, strconv.Itoa(i)+" "+string(buf))
	}
	return 1
}

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		cat(bufio.NewReader(os.Stdin))
	}
	for i := 0; i < flag.NArg(); i++ {
		f, err := os.OpenFile(flag.Arg(i), os.O_RDONLY, 0)
		if err != nil {
			fmt.Printf("打开文件失败")
		}
		cat(bufio.NewReader(f))
	}
}
