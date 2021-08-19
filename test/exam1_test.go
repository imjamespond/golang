package test

import (
	"log"
	"testing"

	"github.com/minio/minio-go"
)

func TestExam1(t *testing.T) {
	useSSL := true

	// 初使化 minio client对象。
	minioClient, err := minio.New(Endpoint, AccessKeyID, SecretAccessKey, useSSL)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("%#v\n", minioClient) // minioClient初使化成功
}
