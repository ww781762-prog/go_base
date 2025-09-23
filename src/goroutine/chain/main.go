package main

func main() {
	var a chan int = make(chan int)
	<-a

}
