package tplus

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"net/http"
)

var (
	appKey      = "a4e5cec6-7070-4858-a802-57e13adbd699"
	appSecret   = "bjc4b5"
	accessToken = ""
)

func CallGet(url string, jsonBody *string) string {
	authCode := getAuth(url, appKey, appSecret)
	return HttpGet(url, func(r *http.Request) {
		r.Header.Add("Authorization", authCode)
		if jsonBody != nil {
			q := r.URL.Query()
			q.Add("_args", *jsonBody)
			r.URL.RawQuery = q.Encode()
		}
	})
}

func CallPost(url string, jsonBody *string) string {
	authCode := getAuth(url, appKey, appSecret)
	return HttpPostForm(url, nil, func(r *http.Request) {
		r.Header.Add("Authorization", authCode)
		if jsonBody != nil {
			q := r.URL.Query()
			q.Add("_args", *jsonBody)
			r.URL.RawQuery = q.Encode()
		}
	})
}

func getAuth(url string, appKey string, appSecret string) string {
	authstr := ""
	utcDate := GetGMTDateStr()
	authParamInfo := "{\"uri\":\"" + url + "\",\"access_token\":\"" + accessToken + "\",\"date\":\"" + utcDate + "\"}"
	encodedHashValue := base64.StdEncoding.EncodeToString([]byte(HMACSHA1(appSecret, authParamInfo)))
	encodedHashValue = "{\"appKey\":\"" + appKey + "\",\"authInfo\":\"hmac-sha1 " + encodedHashValue + "\",\"paramInfo\":" + authParamInfo + "}"
	authstr = base64.StdEncoding.EncodeToString([]byte(encodedHashValue))
	return authstr
}

/*
//  keyStr 密钥
//  value  消息内容
*/
func HMACSHA1(keyStr, value string) string {

	key := []byte(keyStr)
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte(value))
	//进行base64编码
	res := base64.StdEncoding.EncodeToString(mac.Sum(nil))

	return res
}
