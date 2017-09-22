package main

import (
	"fmt"
	"igen.com/bee/thrifttest/thriftgen"
	"github.com/apache/thrift/lib/go/thrift"
	"time"
)

type IDataThriftServiceImpl struct {
}

func (this *IDataThriftServiceImpl) GetInvertorBeanByDate(gatewaySn string, invSn string, startDate string, endDate string, timeZone int32) (r []*thriftgen.InvertorThriftBean, err error){
	fmt.Printf("接收请求...%s\n",gatewaySn)
	invertBean := thriftgen.NewInvertorThriftBean()
	invertBean.EToday = "1"
	r = append(r,invertBean)
	return r,err
}

func (this *IDataThriftServiceImpl) Echo() (r string, err error){
	fmt.Println("接收请求...\n")
	return time.Now().Format("2006-01-02 15:04:05"),err
}

func main() {
	var listen string = ":9090"
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	serverTransport, err := thrift.NewTServerSocket(listen)
	if err != nil {
		fmt.Println("error, thrift init!")
		return
	}
	handler := &IDataThriftServiceImpl{}
	processor := thriftgen.NewIDataThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	server.Serve()
}
