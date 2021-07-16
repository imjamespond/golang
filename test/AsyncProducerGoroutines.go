package test

import (
	"log"
	"os"
	"os/signal"
	"sync"

	"github.com/Shopify/sarama"
	mocks "github.com/Shopify/sarama/mocks"
)

func AsyncProducerGoroutines() {
	config := mocks.NewTestConfig()
	config.Producer.Return.Successes = true
	config.Net.MaxOpenRequests = 10
	producer, err := sarama.NewAsyncProducer(Addrs, config)
	if err != nil {
		panic(err)
	}
	// Trap SIGINT to trigger a graceful shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	var (
		wg                                  sync.WaitGroup
		enqueued, successes, producerErrors int
	)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for range producer.Successes() {
			successes++
		}
		log.Println("exit success goroutine")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for err := range producer.Errors() {
			log.Println(err)
			producerErrors++
		}
		log.Println("exit err goroutine")
	}()

ProducerLoop:
	for i := 0; i < 100; i++ {
		message := &sarama.ProducerMessage{Topic: Topic, Value: sarama.StringEncoder("testing 123")}
		select {
		case producer.Input() <- message:
			enqueued++

		case <-signals:
			producer.AsyncClose() // Trigger a shutdown of the producer.
			break ProducerLoop
		}
	}

	producer.AsyncClose()

	wg.Wait()

	log.Printf("Successfully produced: %d; errors: %d\n", successes, producerErrors)
}
