package datasource


import (
	"StreamChannelSwitch/config"
	"errors"
	"fmt"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

type KakfaSouce struct {
	ServerAddr    string
	consumer 	  *kafka.Consumer
	Topics 		  []string
}

type ReportInfo struct {
	Msg  []byte
	Topic string
}

func NewKakfaSouce(config *config.Config) *KakfaSouce {

	client := KakfaSouce{
		ServerAddr:config.KafkaAddr,
	}
	return &client
}


func (k *KakfaSouce)NewConsumer(groupid string, topics []string)  error {

	if len(topics) == 0{
		return errors.New("topic list can not empty")
	}

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": k.ServerAddr,
		"group.id":          groupid,
		"auto.offset.reset": "earliest",
	})

	if err != nil{
		return errors.New("create consumer faile")
	}
	c.SubscribeTopics(topics, nil)
	return nil
}

func (k *KakfaSouce)Fetch() <- chan ReportInfo {

	ch := make(chan ReportInfo)
	k.consumer.SubscribeTopics(k.Topics, nil)

	go func(){
		for {
		msg, err := k.consumer.ReadMessage(-1)
		if err == nil {
			ch <-ReportInfo{
				Msg:   msg.Value,
				Topic: *msg.TopicPartition.Topic,
			}
		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
	}()
	return ch
}
