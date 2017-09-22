package main

import (
	"os"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
	"github.com/astaxie/beego/logs"
	"igen.com/bee/kafka/bg-device-bind/config"
	"strconv"
	"strings"
)

var (
	myConfig *config.Config
	env string
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
	//fmt.Printf("\n getKeys time is: " + time.Now().Format("2006-01-02 15:04:05"))
	var importKeys []string
	conn.Send("MULTI")
	for _, key := range keys {

		conn.Send("HMGET", key, "_et")
		//var value int64
		//updateTime, _ := redis.Values(conn.Do("HMGET", key, "_et"))
		//fmt.Printf("\n updateTime time is: " + time.Now().Format("2006-01-02 15:04:05"))
		//redis.Scan(updateTime, &value)
		//if (value > 0) {
		//	valueStr := strconv.FormatInt(value, 10)
		//	if (len(valueStr) == 10) {
		//		if (time.Now().Unix() < value) {
		//			importKeys = append(importKeys, key)
		//		}
		//	}
		//	if (len(valueStr) == 13) {
		//		if (time.Now().UnixNano() < value) {
		//			importKeys = append(importKeys, key)
		//		}
		//	}
		//}
	}
	results, err := redis.Values(conn.Do("EXEC"))

	for index,result := range results{
		uptime := result.([]interface{})
		if(uptime!=nil && len(uptime)>0){
			var utime string
			redis.Scan(uptime,&utime)
			//utime := uptime.({}string)
			updateTime, _ := strconv.ParseInt(utime,10, 64)
			//fmt.Printf("\n updateTime time is: " + time.Now().Format("2006-01-02 15:04:05"))
			if (updateTime > 0) {
				valueStr := strconv.FormatInt(updateTime, 10)
				if (len(valueStr) == 10) {
					if (time.Now().Unix() < updateTime) {
						importKeys = append(importKeys, keys[index])
						println(updateTime)
					}
				}
				if (len(valueStr) == 13) {
					if (time.Now().UnixNano() < updateTime) {
						importKeys = append(importKeys, keys[index])
						println(updateTime)
					}
				}
			}
		}
	}
	return importKeys
}

func B2S(bs []uint8) string {
	b := make([]byte, len(bs))
	for i, v := range bs {
		b[i] = byte(v)
	}
	return string(b)
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
	logs.SetLogger(logs.AdapterFile, `{"filename":"` + logPath + `","level":7,"maxlines":1000000,"rotate ":true,"daily":true,"maxdays":5}`)


	sourceRedisUrl := myConfig.Read(env, "reids.addr")
	//source, err := redis.DialURL("redis://root:1234@192.168.1.44:6380/0")
	source, err := redis.DialURL(sourceRedisUrl, redis.DialReadTimeout(300 * time.Second), redis.DialWriteTimeout(300 * time.Second))
	handle(err)
	remoteUrl := myConfig.Read(env, "remote.reids.addr")
	//destination, err := redis.DialURL("redis://root:1234@192.168.1.44:6381/1")
	destination, err := redis.DialURL(remoteUrl, redis.DialReadTimeout(300 * time.Second), redis.DialWriteTimeout(300 * time.Second))
	handle(err)

	//importKeys := []string{"iof1*"}
	importKey := myConfig.Read(env, "reids.key")
	importKeys := strings.Split(importKey, ",")
	logs.Info("start time is: " + time.Now().Format("2006-01-02 15:04:05"))
	count := 0
	fmt.Printf("\nstart time is: " + time.Now().Format("2006-01-02 15:04:05")+"\n")
	for _, key := range importKeys {
		keys := getKeys(source, key)
		// Channel where batches of keys will pass.
		queue := make(chan map[string]string, 100)
		// Scan and send to queue.
		go get(source, keys, queue)
		// Restore keys as they come into queue.
		put(destination, queue)
		count += len(keys)
	}
	fmt.Printf("\nend time is: " + time.Now().Format("2006-01-02 15:04:05"))
	logs.Info("end time is: " + time.Now().Format("2006-01-02 15:04:05") + "    count is %d !!!", count)
	fmt.Printf("\nsync done!!!!!!!!!!!!count is %d !!!\n", count)
	source.Close()
	destination.Close()

	//t = time.NewTimer(time.Second * 1)
	//go timeToWrite()
	//select {
	//}

}

//定时写
func timeToWrite() {
	for {
		select {
		case <-t.C:
			sourceRedisUrl := myConfig.Read(env, "reids.addr")
		//source, err := redis.DialURL("redis://root:1234@192.168.1.44:6380/0")
			source, err := redis.DialURL(sourceRedisUrl, redis.DialReadTimeout(300 * time.Second), redis.DialWriteTimeout(300 * time.Second))
			handle(err)
			remoteUrl := myConfig.Read(env, "remote.reids.addr")
		//destination, err := redis.DialURL("redis://root:1234@192.168.1.44:6381/1")
			destination, err := redis.DialURL(remoteUrl, redis.DialReadTimeout(300 * time.Second), redis.DialWriteTimeout(300 * time.Second))
			handle(err)

		//importKeys := []string{"iof1*"}
			importKey := myConfig.Read(env, "reids.key")
			importKeys := strings.Split(importKey, ",")
			logs.Info("start time is: " + time.Now().Format("2006-01-02 15:04:05"))
			count := 0
		fmt.Printf("\nstart time is: " + time.Now().Format("2006-01-02 15:04:05")+"\n")
			for _, key := range importKeys {
				keys := getKeys(source, key)
				// Channel where batches of keys will pass.
				//queue := make(chan map[string]string, 100)
				//// Scan and send to queue.
				//go get(source, keys, queue)
				//// Restore keys as they come into queue.
				//put(destination, queue)
				count += len(keys)
			}
			fmt.Printf("\nend time is: " + time.Now().Format("2006-01-02 15:04:05"))
			logs.Info("end time is: " + time.Now().Format("2006-01-02 15:04:05") + "    count is %d !!!", count)
			fmt.Printf("\nsync done!!!!!!!!!!!!count is %d !!!\n", count)
			source.Close()
			destination.Close()
			//t.Reset(300 * time.Second)
		}
	}
}