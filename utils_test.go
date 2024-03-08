package main

import (
	"fmt"
	"testing"
	"time"
)

func TestCurl(t *testing.T) {
	Curl(`curl 'http://192.168.0.111:8189/api/datamodeler/v2/api-docs' \
  -H 'Accept-Language: zh-CN,zh;q=0.9' \
  -H 'Cache-Control: no-cache' \
  -H 'Connection: keep-alive' \
  -H 'Cookie: SESSION=MmUzY2U4MDctNWYyOC00ZWQ1LWEwODUtNzVjZDI2MzQ1ZjFh' \
  -H 'Pragma: no-cache' \
  -H 'Referer: http://192.168.0.111:8189/api/datamodeler/swagger-ui.html' \
  -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36' \
  -H 'accept: application/json;charset=utf-8,*/*' \
  --insecure`)
}

func TestSave(t *testing.T) {
	Save("out.txt", time.Now().String())
}

func TestRead(t *testing.T) {
	data := Read("out.txt")
	if nil != data {
		fmt.Println(*data)
	}
}

func TestGenTpl(t *testing.T) {
	tpl := `
	/**
	* @description ${pathItem.summary}
	*/
	${name}_Req: Fetcher<${resp}, SWRKeyType<{ ${params} }, {}>>`
	Save("get.inf.tpl", tpl)
	tpl = `
	/**
	* @description ${pathItem.summary}
	*/
	${name}_Req ({ args }) {
		return request('${method}', '${swg.basePath}${pathKey}', args)
	},`
	Save("get.fun.tpl", tpl)
	tpl = `
	/**
	* @description ${pathItem.summary}
	*/
	${name}: MutationFetcher<${resp}, string, ${body}>`
	Save("post.inf.tpl", tpl)
	tpl = `
	/**
	* @description ${pathItem.summary}
	*/
	${name} (_k, extra) {
		return request('${method}', '${swg.basePath}${pathKey}', { body: extra.arg })
	},`
	Save("post.fun.tpl", tpl)
}

func TestPrintBytes(t *testing.T) {
	var val []byte

	val = []byte{0, 1, 2, 3}
	PrintBytes(val)

	val = []byte("你好！")
	PrintBytes(val)

	val = []byte{0xc4, 0xe3, 0xba, 0xc3}
	{
		data, err := GbkToUtf8(val)
		if err == nil {
			fmt.Println(string(data))
		}
	}

}
