package service

import (
	"log"
	"test-etcd/test/strsvc/model"
)

type Service interface {
	Do(model.Request) (string, error)
}

type Test func(model.Request) (string, error)

func CallTest(req model.Request) (string, error) {
	log.Println(req)
	return "test foobar", nil
}
