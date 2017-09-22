package main

import (
	"strings"
	"github.com/bsm/sarama-cluster"
	"os"
	"os/signal"
	"fmt"
)

func main() {
	initKafka()
}

func initKafka() {
	// init (custom) config, enable errors and notifications
	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Group.Return.Notifications = true

	// init consumer
	broker := "192.168.1.60:9092,192.168.1.61:9092,192.168.1.62:9092"
	topic := "deviceAlarm1"
	group := "kafka.group111"
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
				consumer.MarkOffset(msg, "")
				jsonMsg := string(msg.Value)
				fmt.Printf(jsonMsg+"\n")
			}
		case err, more := <-consumer.Errors():
			if more {
				fmt.Printf("Error: %s\n", err.Error())
			}
		case ntf, more := <-consumer.Notifications():
			if more {
				fmt.Printf("Rebalanced: %+v\n", ntf)
			}
		case <-signals:
			return
		}
	}

}