package main

import (
	"time"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var (
	t *time.Timer
)

func main() {
	t = time.NewTimer(time.Second * 1)
	go timeToWrite()
	select{}
}

func timeToWrite() {
	for {
		select {
		case <-t.C:
			sourceRedisUrl := "redis://@118.31.69.127:6379/0"
			//source, err := redis.DialURL("redis://root:1234@192.168.1.44:6380/0")
			source, err := redis.DialURL(sourceRedisUrl,redis.DialReadTimeout(600*time.Second),redis.DialWriteTimeout(600*time.Second))
			handle(err)

			source.Do("FLUSHDB")

			source.Close()
			t.Reset(time.Second * 5)
		}
	}
}

func handle(err error) {
	if err != nil && err != redis.ErrNil {
		fmt.Println(err)
	}
}