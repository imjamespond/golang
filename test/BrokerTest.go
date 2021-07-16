package test

import (
	"fmt"
	"log"

	"github.com/Shopify/sarama"
)

func BrokerTest() {
	broker := sarama.NewBroker(Addr)
	err := broker.Open(nil)
	if err != nil {
		panic(err)
	}

	topicDetails := map[string]*sarama.TopicDetail{
		Topic: {NumPartitions: 2},
	}
	createTopicReq := &sarama.CreateTopicsRequest{TopicDetails: topicDetails}
	createTopicResp, err := broker.CreateTopics(createTopicReq)
	if err != nil {
		panic(err)
	}
	log.Println(createTopicResp)

	fetchReq := sarama.FetchRequest{}
	fetchResp, err := broker.Fetch(&fetchReq)
	if err != nil {
		panic(err)
	}
	log.Println(fetchResp.GetBlock(Topic, 0))

	request := sarama.MetadataRequest{Topics: []string{Topic}}
	response, err := broker.GetMetadata(&request)
	if err != nil {
		_ = broker.Close()
		panic(err)
	}

	fmt.Println("There are", len(response.Topics), "topics active in the cluster.")

	if err = broker.Close(); err != nil {
		panic(err)
	}
}
