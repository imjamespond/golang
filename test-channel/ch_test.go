package main

import (
	"fmt"
	"testing"
	"time"
)

type Job struct {
	name string
}

func producer(ch chan *Job, job *Job) {
	ch <- job
}

func TestCh(t *testing.T) {
	// func main() {

	jobCh := make(chan *Job)

	for i := 0; i < 5; i++ {
		go func(i int) {
			for {
				job := <-jobCh
				fmt.Printf("%d, %s\n", i, job.name)
			}
		}(i)
	}

	for {
		var name string
		fmt.Print("Enter your name: \n")
		if _, err := fmt.Scanf("%s", &name); err != nil {
			fmt.Printf("Error reading: %v\n", err)
		}
		fmt.Println("Hello", name)
		producer(jobCh, &Job{name: name})
	}
}

func Test2(t *testing.T) {
	output1 := make(chan string)
	output2 := make(chan string)
	go server1(output1)
	go server2(output2)
	select {
	case s1 := <-output1:
		fmt.Println(s1)
	case s2 := <-output2:
		fmt.Println(s2)
	}

}

func server1(ch chan string) {
	time.Sleep(6 * time.Second)
	ch <- "from server1"
}
func server2(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "from server2"
}
