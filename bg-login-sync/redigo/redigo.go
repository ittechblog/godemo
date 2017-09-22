package main

import (
	"time"
	"github.com/astaxie/beego/logs"
	"fmt"
	"os"
	"igen.com/bee/kafka/bg-device-bind/config"
	"github.com/garyburd/redigo/redis"
	"strconv"
)

var (
	myConfig *config.Config
	env      string
	t *time.Timer
)

func initConfig(filePath string) {
	myConfig = new(config.Config)
	myConfig.InitConfig(filePath)
	fmt.Printf("%v", myConfig.Mymap)
}

// Report all errors to stdout.
func handle(err error) {
	if err != nil && err != redis.ErrNil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func getKeys(conn redis.Conn, importKey string) []string {
	keys, err := redis.Strings(conn.Do("KEYS", importKey))
	handle(err)
	var importKeys []string
	for _, key := range keys {
		var value int64
		updateTime, _ := redis.Values(conn.Do("HMGET", key, "_et"))
		redis.Scan(updateTime, &value)
		if (value > 0) {
			valueStr := strconv.FormatInt(value, 10)
			if (len(valueStr) == 10) {
				if (time.Now().Unix() < value) {
					importKeys = append(importKeys, key)
				}
			}
			if (len(valueStr) == 13) {
				if (time.Now().UnixNano() < value) {
					importKeys = append(importKeys, key)
				}
			}
		}
	}
	return importKeys
}

// Scan and queue source keys.
func get(conn redis.Conn, keys []string, queue chan <- map[string]string) {
	var (
		cursor int64
		//keys []string
	)

	for {
		// Scan a batch of keys.
		//values, err := redis.Values(conn.Do("SCAN", cursor))
		//handle(err)
		//values, err = redis.Scan(values, &cursor, &keys)
		//handle(err)

		// Get pipelined dumps.
		for _, key := range keys {
			conn.Send("DUMP", key)
		}
		dumps, err := redis.Strings(conn.Do(""))
		handle(err)

		// Build batch map.
		batch := make(map[string]string)
		for i, _ := range keys {
			batch[keys[i]] = dumps[i]
		}
		cursor = 0

		// Last iteration of scan.
		if cursor == 0 {
			// queue last batch.
			select {
			case queue <- batch:
			}
			close(queue)
			break
		}

		fmt.Printf(">")
		// queue current batch.
		queue <- batch
	}
}

// Restore a batch of keys on destination.
func put(conn redis.Conn, queue <-chan map[string]string) {
	for batch := range queue {
		for key, value := range batch {
			conn.Send("RESTORE", key, "0", value)
		}
		_, err := conn.Do("")
		handle(err)
		fmt.Printf(".")
	}
}

func main() {
	//from := flag.String("from", "", "redis://192.168.1.44:6381/0")
	//to := flag.String("to", "", "redis://192.168.1.44:6381/1")
	//flag.Parse()

	if len(os.Args) == 2 {
		filePath := os.Args[1]
		//初始化配置
		initConfig(filePath)
		env = "default"
	} else if len(os.Args) >= 3 {
		env = os.Args[2]
	} else {
		fmt.Println("需要指定配置文件路径！！！！！！！")
		os.Exit(0)
	}
	//日志初始化
	logPath := myConfig.Read(env, "log.path")
	logs.SetLogger(logs.AdapterFile, `{"filename":"`+logPath+`","level":7,"maxlines":1000000,"rotate ":true,"daily":true,"maxdays":5}`)

	t = time.NewTimer(time.Second * 1)
	go timeToWrite()
	select{}

}


//定时写
func timeToWrite() {
	for {
		select {
		case <-t.C:
			sourceRedisUrl := myConfig.Read(env, "reids.addr")
		//source, err := redis.DialURL("redis://root:1234@192.168.1.44:6380/0")
			source, err := redis.DialURL(sourceRedisUrl,redis.DialReadTimeout(600*time.Second),redis.DialWriteTimeout(600*time.Second))
			handle(err)

			remoteUrl:= myConfig.Read(env, "remote.reids.addr")
		//destination, err := redis.DialURL("redis://root:1234@192.168.1.44:6381/1")
			destination, err := redis.DialURL(remoteUrl,redis.DialReadTimeout(600*time.Second),redis.DialWriteTimeout(600*time.Second))

			handle(err)

			importKeys := []string{"iof1*"}
		//importKeys := []string{"dv:*", "dl:*"}
		//var allKeys []string = make([]string,10000)
			logs.Info("start time is: "+time.Now().Format("2006-01-02 15:04:05"))
			count:=0
			for _, key := range importKeys {
				keys := getKeys(source, key)
				count += len(keys)
			}
			logs.Info("end time is: "+time.Now().Format("2006-01-02 15:04:05")+"    count is %d !!!",count)
			fmt.Printf("sync done!!!!!!!!!!!!\n")
			source.Close()
			destination.Close()
		//t.Reset(time.Second * 300)
		}
	}
}