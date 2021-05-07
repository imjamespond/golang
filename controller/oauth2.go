package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

// Protect and proxy oauth2 server api

func OAuth2_GetToken(c *gin.Context) {
	// 取得 auth code
	code := c.Request.URL.Query().Get("code")
	fmt.Println(code)
	/*
		curl --location --request GET 'http://localhost:8080/oauth2/authorize-code?client_id=000000&redirect_uri=http://localhost:8080/test-oauth2/getToken&response_type=code&state=12345&userId=1001' \
		--header 'Authorization: Basic MDAwMDAwOjk5OTk5OQ==' -v
	*/

	// 用 auth code 获取 access token
	req, err := http.NewRequest("GET", "http://localhost:8001/token", nil)
	if err != nil {
		panic(err.Error())
	}

	q := req.URL.Query()
	q.Add("grant_type", "authorization_code")
	q.Add("code", code)
	q.Add("client_id", "000000")
	q.Add("client_secret", "999999")
	q.Add("redirect_uri", "http://localhost:8080/test-oauth2/getToken")
	req.URL.RawQuery = q.Encode()

	queryUrl := req.URL.String()
	fmt.Println(queryUrl)
	resp, err := http.Get(queryUrl)
	if err != nil {
		panic(err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(string(body))

	var result map[string]json.RawMessage
	if err := json.Unmarshal(body, &result); err != nil {
		panic(err.Error())
	}
	var token string
	var exp int32
	json.Unmarshal(result["access_token"], &token)
	json.Unmarshal(result["expires_in"], &exp)
	fmt.Println(token, exp)

	// 测试 access token
	testReq, _ := http.NewRequest("GET", "http://localhost:8001/test", nil)
	testReq.Header.Add("Authorization", "Bearer "+token)
	client := &http.Client{}
	testResp, _ := client.Do(testReq)
	testRs, _ := ioutil.ReadAll(testResp.Body)
	fmt.Println(string(testRs))
}

func OAuth2_Proxy(c *gin.Context) {
	remote, err := url.Parse("http://localhost:8001")
	if err != nil {
		panic(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)
	//Define the director func
	//This is a good place to log, for example
	proxy.Director = func(req *http.Request) {
		req.Header = c.Request.Header
		req.Host = remote.Host
		req.URL.Scheme = remote.Scheme
		req.URL.Host = remote.Host
		// req.URL.Path = c.Param("proxyPath")
	}

	c.Request.URL.Path = strings.Replace(c.Request.URL.Path, "/oauth2", "", 1)

	proxy.ServeHTTP(c.Writer, c.Request)
}

// router.POST("/api/v1/endpoint1", ReverseProxy()
// func ReverseProxy() gin.HandlerFunc {
// 	target := "localhost:3000"
// 	return func(c *gin.Context) {
// 			director := func(req *http.Request) {
// 					r := c.Request
// 					req = r
// 					req.URL.Scheme = "http"
// 					req.URL.Host = target
// 					req.Header["my-header"] = []string{r.Header.Get("my-header")}
// 											// Golang camelcases headers
// 					delete(req.Header, "My-Header")
// 			}
// 			proxy := &httputil.ReverseProxy{Director: director}
// 			proxy.ServeHTTP(c.Writer, c.Request)
// 	}
// }
