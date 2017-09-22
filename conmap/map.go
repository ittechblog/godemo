package main

import "sync"

func main() {
	syncMap := sync.Map{}
	syncMap.Store("key","value")
	syncMap.Delete("key")
	value,ok :=syncMap.Load("key")
	if ok{
		println(value.(string))
	}else{
		println("-------")
	}

}
