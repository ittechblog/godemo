package main

import (
	"os"
	"fmt"
	"github.com/astaxie/beego/logs"
	"igen.com/bee/kafka/bg-device-bind/config"
	"strconv"
	"github.com/go-redis/redis"
	"time"
)

var (
	myConfig *config.Config
	env string
)

func initConfig(filePath string) {
	myConfig = new(config.Config)
	myConfig.InitConfig(filePath)
	fmt.Printf("%v", myConfig.Mymap)
}
func NewClient() *redis.Client {
	addr := myConfig.Read(env, "reids.url")
	password := myConfig.Read(env, "reids.password")
	db, _ := strconv.Atoi(myConfig.Read(env, "reids.db"))
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	if err != nil {
		logs.Info("get redis conn fail!!!!!")
	}
	return client
}

func NewRemoteClient() *redis.Client {
	addr := myConfig.Read(env, "remote.reids.url")
	password := myConfig.Read(env, "remote.reids.password")
	db, _ := strconv.Atoi(myConfig.Read(env, "remote.reids.db"))
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	if err != nil {
		logs.Info("get redis conn fail!!!!!")
	}
	return client
}

func readKeys(key string) []string {
	client := NewClient()
	keys, err := client.Keys(key).Result()
	if (err != nil) {
		println(err)
	}
	defer client.Close()
	var importKeys []string
	for _, key := range keys {
		updateDateSlice , _ := client.HMGet(key, "_et").Result()
		for _,updateDate := range updateDateSlice{
			if(updateDate != nil){
				updateTime ,_ := strconv.ParseInt((updateDate.(string)),10, 64)
					valueStr := strconv.FormatInt(updateTime, 10)
					if (len(valueStr) == 10) {
						if (time.Now().Unix() < updateTime) {
							importKeys = append(importKeys, key)
						}
					}
					if (len(valueStr) == 13) {
						if (time.Now().UnixNano() < updateTime) {
							importKeys = append(importKeys, key)
						}
					}
				}
		}
	}
	return importKeys
}

func main() {
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

	importKeys := []string{"dv:*", "dl:*"}
	for _, key := range importKeys {
		keys := readKeys(key)
		println(keys)
		readData(keys)
	}
	//go syncToRemote()
}

func readData(keys []string) {
	client := NewClient()
	for _, key := range keys {
		//result,_ ,_:= client.HScan(key,0,"MATCH",100).Result()
		results, _ := client.HMGet(key, "_et","_st").Result()
		hashFields :=make(map[string]interface{})
		hashFields["_et"] = results[0]
		hashFields["_st"] = results[1]
		//hashFields := map[string]interface{
		//	"_et" : results[0],
		//	"_st" : results[1],
		//}
		client.HMSet(key,hashFields)
		for _, result := range results {
			if (result != nil) {
				fmt.Printf("%s ", result.(string))
			}
		}
		fmt.Printf("\n")
	}
}

func syncToRemote() {
	//client := NewRemoteClient()
	//client.HMSet(key, fields)
}