package utils

import (
	"context"
	"github.com/Shopify/sarama"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type kafka struct {
	brokers           []string
	topics            []string
	startOffset       int64
	version           string
	ready             chan bool
	group             string
	channelBufferSize int
	fn                func(*sarama.ConsumerMessage)
}

func NewKafka(brokers, topics []string, group string, fn func(msg *sarama.ConsumerMessage)) *kafka {
	return &kafka{
		brokers:           brokers,
		topics:            topics,
		group:             group,
		channelBufferSize: 2,
		ready:             make(chan bool),
		version:           "1.1.1",
		fn:                fn,
	}
}

func (p *kafka) Init() func() {
	log.Infoln("kafka init...")
	version, err := sarama.ParseKafkaVersion(p.version)
	if err != nil {
		log.Fatalf("Error parsing Kafka version: %v", err)
	}
	config := sarama.NewConfig()
	config.Version = version
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange // 分区分配策略
	config.Consumer.Offsets.Initial = -2                                   // 未找到组消费位移的时候从哪边开始消费
	config.ChannelBufferSize = p.channelBufferSize                         // channel长度

	ctx, cancel := context.WithCancel(context.Background())
	client, err := sarama.NewConsumerGroup(p.brokers, p.group, config)
	if err != nil {
		log.Fatalf("Error creating consumer group client: %v", err)
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
		}()
		for {
			if err := client.Consume(ctx, p.topics, p); err != nil {
				log.Fatalf("Error from consumer: %v", err)
			}
			if ctx.Err() != nil {
				log.Println(ctx.Err())
				return
			}
			p.ready = make(chan bool)
		}
	}()
	<-p.ready
	// 保证在系统退出时，通道里面的消息被消费
	return func() {
		log.Info("kafka 关闭")
		cancel()
		wg.Wait()
		if err = client.Close(); err != nil {
			log.Errorf("Error closing client: %v", err)
		}
	}
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (p *kafka) Setup(sarama.ConsumerGroupSession) error {
	// Mark the consumer as ready
	close(p.ready)
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (p *kafka) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// 消费方法
func (p *kafka) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// 具体消费消息
	for message := range claim.Messages() {
		p.fn(message)
		// time.Sleep(time.Second)
		// 更新位移
		session.MarkMessage(message, "")
	}
	return nil
}

func ConsumeGroup(adders, topic []string, group string, fn func(msg *sarama.ConsumerMessage)) {
	k := NewKafka(adders, topic, group, fn)
	f := k.Init()
	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-sigterm:
		log.Warnln("terminating: via signal")
	}
	f()
}
