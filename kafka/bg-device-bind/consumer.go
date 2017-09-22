package main

import (
	"os"
	"os/signal"
	"database/sql"
	"fmt"
	"strings"

	cluster "github.com/bsm/sarama-cluster"
	_ "github.com/go-sql-driver/mysql"
	"igen.com/bee/kafka/bg-device-bind/config"
	"igen.com/bee/kafka/bg-device-bind/process"
	"strconv"
	"github.com/astaxie/beego/logs"
	"net/http"
	_ "net/http/pprof"
	"github.com/tidwall/gjson"
	"time"
)

var (
	myConfig *config.Config
	env string
	t *time.Timer
	jsonLogMsg string
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
	maxIdleConns, _ := strconv.Atoi(myConfig.Read(env, "maxidleconns"))
	maxOpenConns, _ := strconv.Atoi(myConfig.Read(env, "maxopenconns"))
	connSql := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8"
	config.Db, _ = sql.Open("mysql", connSql)
	config.Db.SetMaxOpenConns(maxOpenConns)
	config.Db.SetMaxIdleConns(maxIdleConns)
	config.Db.Ping()
}

func initKafka() {
	// init (custom) config, enable errors and notifications
	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Group.Return.Notifications = true

	// init consumer
	broker := myConfig.Read(env, "kafka.brokers")
	topic := myConfig.Read(env, "kafka.topic")
	group := myConfig.Read(env, "kafka.group")
	brokers := strings.Split(broker, ",")
	topics := []string{topic}
	consumer, err := cluster.NewConsumer(brokers, group, topics, config)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	// trap SIGINT to trigger a shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	// consume messages, watch errors and notifications
	for {
		select {
		case msg, more := <-consumer.Messages():
			if more {
				jsonMsg := string(msg.Value)
				jsonMsg = getJsonMsg(jsonMsg)
				if (gjson.Valid(jsonMsg)) {
					consumer.MarkOffset(msg, "") // mark message as processed
					processChan := make(chan int)
					//业务处理
					commandType := gjson.Get(jsonMsg, "a").String()
					if ( !strings.EqualFold("100", commandType) && !strings.EqualFold("101", commandType)) {
						if (strings.EqualFold("01", commandType) || strings.EqualFold("1", commandType) || strings.EqualFold("11", commandType)) {
							jsonLogMsg = jsonMsg
						}
						go process.Process(jsonMsg, processChan)
						<-processChan
					}
				} else {
					println("jsonMsg is not valid")
				}
			}
		case err, more := <-consumer.Errors():
			if more {
				logs.Info("Error: %s\n", err.Error())
			}
		case ntf, more := <-consumer.Notifications():
			if more {
				logs.Info("Rebalanced: %+v\n", ntf)
			}
		case <-signals:
			return
		}
	}

}

//定时写日志
func timeToLog() {
	for {
		select {
		case <-t.C:
			if (len(jsonLogMsg) > 0) {
				uploadDate := gjson.Get(jsonLogMsg, "zd").String()
				uploadDate = process.ConvertUploadTime(uploadDate)
				var datetime = time.Now()
				datetime.Format("2006-01-02 15:04:05")
				fmt.Printf("%s\n", datetime)
				logs.Info("receive time is %s,current time is %s \n", uploadDate, datetime)
			}
			t.Reset(time.Second * 60)
		}
	}
}

//递归获取json
func getJsonMsg(stringMsg string) (jsonMsg string) {
	startIndex := strings.Index(stringMsg, "{")
	endIndex := strings.Index(stringMsg, "}")
	if endIndex < startIndex {
		getJsonMsg(stringMsg[startIndex:])
	} else {
		jsonMsg = stringMsg[startIndex : endIndex + 1]
	}
	return
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
	//初始化mysql
	initMysql()

	go func() {
		httpPort := myConfig.Read(env, "http.port")
		http.ListenAndServe(":" + httpPort, nil)
	}()
	println("\nsystem is runing.............................")
	t = time.NewTimer(time.Second * 60)

	go timeToLog()

	//初始化kafka
	initKafka()

}

