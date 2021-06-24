package testchannel

import (
	"log"
	"testing"
	"time"
)

func TestTimeout(t *testing.T) {
	log.Printf("begin")
	<-time.After(3 * time.Second)
	log.Printf("end")
}
