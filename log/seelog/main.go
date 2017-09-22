package main

import (
	"github.com/cihub/seelog"
	"github.com/tidwall/gjson"
)

type MyLog struct {
	Date  string `json:"date"`
	Msg string `json:"msg"`
	Level   string    `json:"level"`
}

var json = `{
	"date": "2017-06-29 10:15:25",
	"msg": "Bark",
	"level": "Info"
}`

func main() {
	mainlog, _ := seelog.LoggerFromConfigAsFile("conf/seelog-main.xml")
	defer mainlog.Flush()
	mainlog.Info("需要输入的日志")

	mainlog2, _ := seelog.LoggerFromConfigAsFile("conf/seelog-main2.xml")
	defer mainlog2.Flush()
	var mylog MyLog
	gjson.Unmarshal([]byte(json),&mylog)
	gjson.Parse(json)
	mainlog2.Info(json)
}
