package main

import (
	"fmt"
	"time"
	"strconv"
)

func main() {
	//fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	//uploadTime,_:=time.Parse("2006-01-02 15:04:05","1496641576")
	//fmt.Println(uploadTime.Format("2006-01-02 15:04:05"))
	int64, _ := strconv.ParseInt("1496641576", 10, 64)
	tm := time.Unix(int64,0)
	fmt.Println(tm.Format("2006-01-02 15:04:05"))
}
