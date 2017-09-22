package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
	_ "igen.com/bee/routers"
	"igen.com/bee/util"
	"github.com/astaxie/beego/logs"
)

// Model Struct
type Forward struct {
	Id           int
	DataloggerSn string `orm:"size(100)"`
	Ip           string `orm:"size(100)"`
	Port         string `orm:"size(100)"`
}

func init() {
	initMysql()
	initRedisPool()
	logs.SetLogger(logs.AdapterFile,`{"filename":"project.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10}`)
}

func initMysql() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	// set default database
	dbhost := beego.AppConfig.String("dbhost")
	dbport := beego.AppConfig.String("dbport")
	dbuser := beego.AppConfig.String("dbuser")
	dbpassword := beego.AppConfig.String("dbpassword")
	db := beego.AppConfig.String("db")
	conn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + db + "?charset=utf8"
	orm.RegisterDataBase("default", "mysql", conn, 30)
	orm.SetMaxIdleConns("default", 5)
	orm.SetMaxOpenConns("default", 10)
	// register model
	orm.RegisterModel(new(Forward))

	// create table
	//orm.RunSyncdb("default", false, true)
}

func initRedisPool() {
	// 从配置文件获取redis的ip以及db
	REDIS_HOST := beego.AppConfig.String("reids.addr")
	REDIS_PASS := beego.AppConfig.String("reids.password")
	REDIS_DB, _ := beego.AppConfig.Int("reids.db")

	// 建立连接池
	util.RedisClient = &redis.Pool{
		// 从配置文件获取maxidle以及maxactive，取不到则用后面的默认值
		MaxIdle:     beego.AppConfig.DefaultInt("redis.maxidle", 1),
		MaxActive:   beego.AppConfig.DefaultInt("redis.maxactive", 10),
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", REDIS_HOST)
			if err != nil {
				return nil, err
			}
			// 选择db
			if _, err := c.Do("AUTH", REDIS_PASS); err != nil {
				c.Close()
				return nil, err
			}
			c.Do("SELECT", REDIS_DB)
			return c, nil
		},
	}
}

func page_not_found(rw http.ResponseWriter, r *http.Request) {
	t, _ := template.New("404.html").ParseFiles("/404.html")
	data := make(map[string]interface{})
	data["content"] = "page not found"
	t.Execute(rw, data)
}

func main() {
	beego.BConfig.Listen.EnableAdmin = true
	beego.BConfig.Listen.AdminAddr = "localhost"
	beego.BConfig.Listen.AdminPort = 8088
	//beego.ErrorHandler("404",page_not_found)

	rc := util.RedisClient.Get()
	rc.Do("SET", "test", "test")
	reply, err := rc.Do("GET", "test")
	value, _ := redis.String(reply, err)
	fmt.Println(value)
	defer rc.Close()

	logs.Info("this %s cat is %v years old", "yellow", 3)
	beego.Run()

}
