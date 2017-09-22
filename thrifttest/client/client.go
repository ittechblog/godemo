package main

import (
	"net"
	"log"
	"github.com/apache/thrift/lib/go/thrift"
	"fmt"
	"igen.com/bee/thrifttest/thriftgen"
)

const (
	HOST = "localhost"
	PORT = "9090"
)

func main()  {
	var transport thrift.TTransport
	socket, err := thrift.NewTSocket(net.JoinHostPort(HOST, PORT))
	if err != nil {
		log.Fatalln("tSocket error:", err)
	}
	transport = thrift.NewTFramedTransport(socket)
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	client := thriftgen.NewIDataThriftServiceClientFactory(transport, protocolFactory)

	if err := transport.Open(); err != nil {
		log.Fatalln("Error opening:", HOST + ":" + PORT)
	}
	defer transport.Close()

	result,_ := client.Echo()
	fmt.Printf("client.Echo()-------------------%s\n",result)

	r,err := client.GetInvertorBeanByDate("111","1","1","1",32)
	for _,bean := range r {
		fmt.Printf("client.GetInvertorBeanByDate()-------------------%s\n",bean.EToday)
	}
	fmt.Printf("client.GetInvertorBeanByDate()-------------------%v\n",r)
}