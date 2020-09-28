package tests

import (
	"fmt"
	"github.com/Shopify/sarama"
	"im/config"
	"testing"
	"utils"
)

type Kafka struct {
	brokers           []string
	topics            []string
	startOffset       int64
	version           string
	ready             chan bool
	group             string
	channelBufferSize int
}

func TestKafka(t *testing.T) {
	msg := &sarama.ProducerMessage{}
	msg.Topic = "tests"
	msg.Key = sarama.StringEncoder("key")
	msg.Value = sarama.StringEncoder("good")
	msg.Partition = 1
	println(config.KafkaProducer.SendMessage(msg))
}

func TestConsumer(t *testing.T) {
	utils.ConsumeGroup(config.Config.Im.Adders, []string{"tests"}, "hello", func(msg *sarama.ConsumerMessage) {
		fmt.Println("---> ", string(msg.Value))
	})
	//consumerKafka(config.Config.Im.Adders, config.Config.Im.Topic, 0, func(msg *sarama.ConsumerMessage) {
	//	fmt.Println(string(msg.Value))
	//})
}

func consumerKafka(adders []string, topic string, partition int32, f func(msg *sarama.ConsumerMessage)) {
	consumer, err := sarama.NewConsumer(adders, nil)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()
	pc, err := consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)
	if err != nil {
		panic(err)
	}

	for msg := range pc.Messages() {
		f(msg)
	}
	defer pc.AsyncClose()
}
