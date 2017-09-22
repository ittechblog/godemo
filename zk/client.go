package main

/**
客户端doc地址：github.com/samuel/go-zookeeper/zk
**/
import (
	"fmt"
	zk "github.com/samuel/go-zookeeper/zk"
	"time"
)

/**
 * 获取一个zk连接
 * @return {[type]}
 */
func getConnect(zkList []string) (conn *zk.Conn) {
	conn, _, err := zk.Connect(zkList, 10*time.Second)
	if err != nil {
		fmt.Println(err)
	}
	return
}

/**
 * 测试连接
 * @return
 */
func test1() {
	zkList := []string{"192.168.1.40:2181","192.168.1.41:2181","192.168.1.42:2181"}
	conn := getConnect(zkList)

	defer conn.Close()
	conn.Create("/go_servers", nil, 0, zk.WorldACL(zk.PermAll))
	conn.Delete("/go_servers",-1)

	time.Sleep(20 * time.Second)
}

/**
 * 测试临时节点
 * @return {[type]}
 */
func test2() {
	zkList := []string{"192.168.1.40:2181","192.168.1.41:2181","192.168.1.42:2181"}
	conn := getConnect(zkList)

	defer conn.Close()
	conn.Create("/testadaadsasdsaw", nil, zk.FlagEphemeral, zk.WorldACL(zk.PermAll))

	time.Sleep(20 * time.Second)
}

/**
 * 获取所有节点
 */
func test3() {
	zkList := []string{"192.168.1.40:2181","192.168.1.41:2181","192.168.1.42:2181"}
	conn := getConnect(zkList)

	defer conn.Close()

	children, _, err := conn.Children("/login-thrift-service")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v \n", children)

	//parse, _, err := conn.Children("/parse-thrift-service")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Printf("%v \n", parse)
}

func main() {
	test3()
}