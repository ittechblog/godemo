package main

import (
	"fmt"
	"runtime"
	"time"
)

var c = make(chan int,5)

func f() {
	c <- 'c'
	time.Sleep(time.Second * 2)
	fmt.Println("在goroutine内")
}


func main() {
	println(time.Now().Add(15*time.Minute).Unix())
	println(time.Now().Unix())
	//t := make (chan int)
	for i := 0; i < 5; i++ {
		//go func(){
		//	time.Sleep(time.Second * 1)
		//	fmt.Println("111111111111")
		//}()
		//go sum2(i,18,t)
		//fmt.Println(<-t)
		go f()
		<-c
		//c <- 'c'
		//<-c
	}

	println(runtime.NumGoroutine())
	time.Sleep(time.Second * 3)
}

