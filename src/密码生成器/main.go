package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var (
	length int
	ty     string
)

const (
	chars string = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	nums  string = "1234567890"
	spec  string = "!@#$%^&*"
)

func argsPass() {
	flag.IntVar(&length, "l", 16, "生产密码的长度")
	flag.StringVar(&ty, "t", "num", `
		num:数字
		char:字符串
		mix: 数字和字符串
		adva: 数字 字符串 特殊字符
	`)
	flag.Parse()
}

func genPassWord() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	sourceStr := ""
	passStr := make([]byte, 0, length)
	if ty == "num" {
		sourceStr = nums
	} else if ty == "char" {
		sourceStr = chars
	} else if ty == "mix" {
		sourceStr = nums + chars
	} else if ty == "adva" {
		sourceStr = nums + chars + spec
	} else {
		sourceStr = nums
	}
	for i := 0; i < length; i++ {
		passStr = append(passStr, sourceStr[r.Intn(len(sourceStr))])
	}
	return string(passStr)
}

func main() {
	argsPass()
	fmt.Println(genPassWord())
}
