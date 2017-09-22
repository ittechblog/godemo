package main

import (
	"github.com/Shopify/sarama"
	"time"
	"strings"
	"github.com/golang/glog"
	"strconv"
	"fmt"
	"os"
	"math/rand"
)

func main() {
	for i:=0;i<10 ;i++  {
		syncProducer()
	}

}

func syncProducer() {
	config := sarama.NewConfig()
	//  config.Producer.RequiredAcks = sarama.WaitForAll
	//  config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	config.Producer.Timeout = 5 * time.Second
	p, err := sarama.NewSyncProducer(strings.Split("192.168.1.60:9092,192.168.1.61:9092,192.168.1.62:9092", ","), config)
	defer p.Close()
	if err != nil {
		glog.Errorln(err)
		return
	}

	v := "sync: " + strconv.Itoa(rand.New(rand.NewSource(time.Now().UnixNano())).Intn(10000))
	fmt.Fprintln(os.Stdout, v)
	msg := &sarama.ProducerMessage{
		Topic: "deviceAlarm1",
		Value: sarama.ByteEncoder(v),
	}
	if _, _, err := p.SendMessage(msg); err != nil {
		glog.Errorln(err)
		return
	}
}