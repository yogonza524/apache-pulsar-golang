package model

import (
	"context"
	"encoding/json"
	"runtime"

	"github.com/apache/pulsar/pulsar-client-go/pulsar"
	log "github.com/sirupsen/logrus"
)

const TOPIC string = "tweets"

type PulsarConnector interface {
	Connect()
	Produce(message string) int
	Consume() int
}

type Pulsar struct {
	Status int
	Client pulsar.Client
	TwitterProducer pulsar.Producer
}

func (p *Pulsar) Connect() {
	client, err := pulsar.NewClient(pulsar.ClientOptions{
        URL: "pulsar://localhost:6650",
        OperationTimeoutSeconds: 5,
        MessageListenerThreads: runtime.NumCPU(),
	})
	if err != nil {
		panic(err)
	}
	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic : TOPIC,
	})
	if err != nil {
		panic(err)
	}
	p.Status = 1
	p.TwitterProducer = producer
	p.Client = client
	out,_ := json.Marshal(p)
	log.WithField("connector", string(out)).Info("Connection")
}

func (p *Pulsar) Produce(message string) int {
	if p.Status == 0 {
		panic("Error: Connection is closed")
	}
	msg := pulsar.ProducerMessage{
		Payload: []byte(message),
	}
	if err := p.TwitterProducer.Send(context.Background(), msg); err != nil {
		log.Fatalf("Producer could not send tweet: %v", err)
		return 1
	}
	return 0
}

func (p *Pulsar) Consume() int {
	if p.Status == 0 {
		panic("Error: Connection is closed")
	}

	msgChannel := make(chan pulsar.ConsumerMessage)

	consumerOpts := pulsar.ConsumerOptions{
		Topic:            TOPIC,
		SubscriptionName: "twitter-go-consumer",
		Type:             pulsar.Exclusive,
		MessageChannel:   msgChannel,
	}

	consumer, err := p.Client.Subscribe(consumerOpts)

	if err != nil {
		log.Fatalf("Could not establish subscription: %v", err)
	}

	defer consumer.Close()

	for cm := range msgChannel {
		msg := cm.Message

		log.WithField("ID", msg.ID()).Info("Message ID consumed")
		log.WithField("Value", string(msg.Payload())).Info("Message Value consumed")

		consumer.Ack(msg)
	}
	return 0
}

