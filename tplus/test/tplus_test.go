package main

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"sd-2110/tplus"
)

func TestTplus(t *testing.T) {
	fmt.Println(tplus.GetGMTDateStr())
}

func TestTplusGet(t *testing.T) {
	tplus.HttpGet("http://localhost/submit_form.php", func(req *http.Request) {
		q := req.URL.Query()
		q.Add("api_key", "key_from_environment_or_flag")
		q.Add("another_thing", "foo & bar")
		req.URL.RawQuery = q.Encode()
	})
	str := `{"foo":"bar"}`
	tplus.CallGet("http://localhost", &str)
}

func TestTplusPost(t *testing.T) {
	data := url.Values{}
	data.Set("foo", "bar")
	data.Set("bar", "baz")

	tplus.HttpPostForm("http://localhost/submit_form.php", &data, func(r *http.Request) {

	})
}
