package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestHttp(t *testing.T) {

}

func TestHttpGet(t *testing.T) {
	resp, err := http.Get("https://163.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		fmt.Println(string(body))
	}
}
