package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"github.com/astaxie/beego/orm"
	"igen.com/bee/util"
	"strconv"
	"github.com/go-redis/redis"
)

type WriteController struct {
	beego.Controller
}

type Map interface {
	Value(string) interface{}
}
type StringMap map[string]string

var complete chan int = make(chan int)

func (m StringMap) Value(key string) interface{} {
	return m[key]
}

func (c *WriteController) Get() {
	REDIS_KEY := beego.AppConfig.String("reids.key")
	o := orm.NewOrm()
	var maps []orm.Params
	_, err := o.Raw("SELECT datalogger_sn,ip,port from forward where forward_mode=1 and protocol=1").Values(&maps)

	i := 0
	client := util.NewClient()
	if (err != nil) {
		fmt.Println("Query database error!!!!!")
	} else {
		//pipe := client.Pipeline()
		//rc := util.RedisClient.Get()
		for _, term := range maps {
			//fields["ip"]=strconv.Quote(term["ip"])
			//fields["port"]=strconv.Quote(term["port"])
			//fmt.Println(term["datalogger_sn"],term["ip"],":",term["port"])
			util.UpdateF2opOne(client,REDIS_KEY+term["datalogger_sn"].(string),term)
			i++
			//pipe.HMSet("iof1:"+term["datalogger_sn"].(string),term)

			//rc.Send("HMSET", REDIS_KEY + term["datalogger_sn"].(string), "ip", term["ip"].(string), "port", term["port"].(string))
		}
		//rc.Flush()
		fmt.Println("done!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
		//pipe.Exec()
		//defer rc.Close()
		//go insertToredis(maps,client)
		//<- complete // 直到线程跑完, 取到消息. main在此阻塞住
	}
	c.Data["Website"] = "Write " + strconv.Itoa(i) + " records to redis done!!!!!!"
	c.Data["Email"] = "Write " + strconv.Itoa(i) + " records to redis done!!!!!!"
	c.TplName = "index.tpl"
}

func insertToredis(maps []orm.Params, client *redis.Client) {
	i := 1
	for _, term := range maps {
		//fields["ip"]=strconv.Quote(term["ip"])
		//fields["port"]=strconv.Quote(term["port"])
		fmt.Println(term["datalogger_sn"], term["ip"], ":", term["port"])
		util.UpdateF2opOne(client, "iof1:" + term["datalogger_sn"].(string), term)
		i++
	}
	complete <- 0 // 执行完毕了，发个消息
}