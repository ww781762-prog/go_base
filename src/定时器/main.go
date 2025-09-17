package main

import (
	"fmt"
	"time"
)

func processTask(t time.Time) {
	fmt.Printf("hello worrd,%v\n", t)
}

func main() {
	Tic := time.Tick(10 * time.Second)
	for i := range Tic {
		processTask(i)
	}
}
