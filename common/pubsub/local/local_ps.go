package local

import (
	"context"
	"log"
	"sync"
	"to_do_list/common"
	"to_do_list/common/pubsub"
)

type PubSub interface {
	Publish(ctx context.Context, channel string, data *pubsub.Message) error
	Subscribe(ctx context.Context, topic string) (ch <-chan *pubsub.Message, unsubscribe func())
}

type localPubSub struct {
	name         string
	messageQueue chan *pubsub.Message
	mapChannel   map[string][]chan *pubsub.Message
	locker       *sync.RWMutex
}

func NewLocalPubSub(name string) PubSub {
	pb := &localPubSub{
		name:         name,
		messageQueue: make(chan *pubsub.Message, 1000),
		mapChannel:   make(map[string][]chan *pubsub.Message),
		locker:       new(sync.RWMutex),
	}

	pb.run()

	return pb
}

func (ps *localPubSub) Publish(ctx context.Context, channel string, data *pubsub.Message) error {
	data.SetChannel(channel)

	go func() {
		defer common.Recover()
		ps.messageQueue <- data
		log.Println("New message published: ", data.String())
	}()
	return nil
}

func (ps *localPubSub) Subscribe(ctx context.Context, topic string) (ch <-chan *pubsub.Message, unsubscribe func()) {
	c := make(chan *pubsub.Message)
	ps.locker.Lock()

	if val, ok := ps.mapChannel[topic]; ok {
		val = append(ps.mapChannel[topic], c)
		ps.mapChannel[topic] = val
	} else {
		ps.mapChannel[topic] = []chan *pubsub.Message{c}
	}

	ps.locker.Unlock()
	return c, func() {
		log.Println("Unsubscribe")

		if chans, ok := ps.mapChannel[topic]; ok {
			for i := range chans {
				if chans[i] == c {
					chans = append(chans[:i], chans[i+1:]...)
					ps.locker.Lock()
					ps.mapChannel[topic] = chans
					ps.locker.Unlock()

					close(c)
					break

				}
			}
		}
	}
}

func (ps *localPubSub) run() error {
	go func() {
		defer common.Recover()
		for {
			mess := <-ps.messageQueue
			log.Println("Message dequeue:", mess.String())

			if subs, ok := ps.mapChannel[mess.Channel()]; ok {
				for i := range subs {
					go func(c chan *pubsub.Message) {
						defer common.Recover()
						c <- mess
						//f(mess)
					}(subs[i])
				}
			}
			//else {
			//	ps.messageQueue <- mess
			//}
		}
	}()

	return nil
}

func (ps *localPubSub) GetPrefix() string {
	return ps.name
}

func (ps *localPubSub) Get() interface{} {
	return ps
}

func (ps *localPubSub) Name() string {
	return ps.name
}

func (ps *localPubSub) InitFlags() {
}

func (ps *localPubSub) Configure() error {
	return nil
}

func (ps *localPubSub) Run() error {
	return nil
}

func (ps *localPubSub) Stop() <-chan bool {
	c := make(chan bool)
	go func() { c <- true }()
	return c
}
