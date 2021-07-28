package service

import (
	"log"
	"test-etcd/test/strsvc/model"
)

type Foobar struct {
	Service
}

func (fb Foobar) Do(req model.Request) (string, error) {
	log.Println(req)
	return "test foobar", nil
}
