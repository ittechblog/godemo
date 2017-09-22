package util

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/astaxie/beego"
)

func NewClient() *redis.Client {
	addr := beego.AppConfig.String("reids.addr")
	password := beego.AppConfig.String("reids.password")
	db, error := beego.AppConfig.Int("reids.db")
	if (error != nil) {
		fmt.Println("Get redis db error")
	}
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	if err != nil {
		return nil
	}
	return client
}

type Map interface {
	Value(string) interface{}
}
type StringMap map[string]string

func (m StringMap) Value(key string) interface{} {
	return m[key]
}
func F(m Map) {
	str, ok := m.Value("hello").(string)
	if ok {
		fmt.Println(str)
	}
}

func UpdateF2opOne(client *redis.Client, key string, fields map[string]interface{}) {

	client.HMSet(key, fields)
}

func Client(client *redis.Client) {

	err := client.Set("key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exists")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exists
}
