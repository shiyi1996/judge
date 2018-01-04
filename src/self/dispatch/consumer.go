/**
 * Created by shiyi on 2017/10/15.
 * Email: shiyi@fightcoder.com
 */

package dispatch

import (
	"self/commons/g"

	"github.com/nsqio/go-nsq"
	log "github.com/sirupsen/logrus"
)

type Consumer struct {
	NsqConsumer *nsq.Consumer
	Cfg         *nsq.Config
	Topic       string
	Channel     string
}

var consumers []*Consumer
var handlerCount chan int

func StartConsume() {
	cfg := g.Conf()

	consumer := new(Consumer)
	go consumer.newConsumer(cfg.Nsq.JudgeTopic, cfg.Nsq.JudgeChannel)
	consumers = append(consumers, consumer)

	handlerCount = make(chan int, cfg.Nsq.HandlerCount)
}

func StopConsume() {
	for _, c := range consumers {
		c.NsqConsumer.Stop()
	}
}

func (this *Consumer) newConsumer(topic, channel string) {
	this.Cfg = nsq.NewConfig()
	this.Topic = topic
	this.Channel = channel

	var err error
	this.NsqConsumer, err = nsq.NewConsumer(topic, channel, this.Cfg)
	if err != nil {
		log.Fatal(err)
	}

	this.NsqConsumer.AddHandler(&Handler{Topic: topic})

	err = this.NsqConsumer.ConnectToNSQLookupds(g.Conf().Nsq.Lookupds)
	if err != nil {
		log.Fatal(err)
	}
}
