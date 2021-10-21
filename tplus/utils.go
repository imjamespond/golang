package tplus

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	netUrl "net/url"
	"strconv"
	"strings"
	"time"
)

func GetGMTDateStr() string {
	return time.Now().UTC().Format(http.TimeFormat)
}

type SetReq func(req *http.Request)

func HttpPost(url string, data string, setReq SetReq) string {
	r := request(url, http.MethodPost, strings.NewReader(data))
	r.Header.Add("Content-Type", "application/json;charset=utf-8")
	r.Header.Add("Content-Length", strconv.Itoa(len(data)))
	if setReq != nil {
		setReq(r)
	}
	return doRequest(r)
}

func HttpPostForm(url string, data *netUrl.Values, setReq SetReq) string {
	r := request(url, http.MethodPost, strings.NewReader(data.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	if setReq != nil {
		setReq(r)
	}
	return doRequest(r)
}

func HttpGet(url string, setReq SetReq) string {
	r := request(url, http.MethodGet, nil)
	if setReq != nil {
		setReq(r)
	}
	return doRequest(r)
}

func request(url string, method string, body io.Reader) *http.Request {
	r, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Fatal(err)
	}

	return r
}

func doRequest(r *http.Request) string {
	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		log.Fatal(err)
	}

	// log.Println(res.Status)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	// log.Println(string(body))
	return string(body)
}
