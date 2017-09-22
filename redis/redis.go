package main

import (
	"fmt"
	"strconv"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"igen.com/bee/kafka/bg-device-bind/config"
	"os"
	"github.com/astaxie/beego/logs"
	"github.com/go-redis/redis"
	"log"
)

var (
	myConfig *config.Config
	env      string
)

func initConfig(filePath string) {
	myConfig = new(config.Config)
	myConfig.InitConfig(filePath)
	fmt.Printf("%v", myConfig.Mymap)
}

func initMysql() {
	dbHost := myConfig.Read(env, "dbhost")
	dbPort := myConfig.Read(env, "dbport")
	dbUser := myConfig.Read(env, "dbuser")
	dbPassword := myConfig.Read(env, "dbpassword")
	dbName := myConfig.Read(env, "db")
	maxIdleConns,_ := strconv.Atoi(myConfig.Read(env, "maxidleconns"))
	maxOpenConns,_ := strconv.Atoi(myConfig.Read(env, "maxopenconns"))
	connSql := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8"
	println(connSql)
	config.Db, _ = sql.Open("mysql", connSql)
	config.Db.SetMaxOpenConns(maxOpenConns)
	config.Db.SetMaxIdleConns(maxIdleConns)
	config.Db.Ping()
}

func NewClient() *redis.Client {
	addr := myConfig.Read(env, "reids.addr")
	password := myConfig.Read(env, "reids.password")
	db,_ := strconv.Atoi(myConfig.Read(env, "reids.db"))
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

func UpdateF2opOne(client *redis.Client, key string, fields map[string]interface{}) {
	client.HMSet(key, fields)
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
	logs.SetLogger(logs.AdapterFile, `{"filename":"`+logPath+`","level":7,"maxlines":1000000,"rotate ":true,"daily":true,"maxdays":2}`)
	//初始化mysql
	initMysql()

	println("\nsystem is runing.............................")

	rows, queryErr := config.Db.Query("select datalogger_sn,ip,port  from forward where forward_mode=1 and protocol=1")
	if (queryErr != nil) {
		logs.Error(queryErr)
	}

	i := 0
	for rows.Next() {
		var dataLoggerSn string
		var ip string
		var port int
		err := rows.Scan(&dataLoggerSn,&ip,&port)
		if err != nil {
			log.Fatal(err)
		}
		client := NewClient()
		tempMap := make(map[string]string)
		tempMap["ip"] = ip
		tempMap["port"] = strconv.Itoa(port)

		var tmp map[string]interface{} = map[string]interface{}{"ip":ip, "port":port}

		UpdateF2opOne(client,"iof1:"+dataLoggerSn,tmp)
		i++
	}
	fmt.Println("Write " + strconv.Itoa(i) + " records to redis done!!!!!!")
	os.Exit(0)
}
