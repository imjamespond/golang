package main

import (
	"testing"
)

func TestCurl(t *testing.T) {
	command := `curl 'http://mail.keymobile.com.cn/' \
	-H 'Connection: keep-alive' \
	-H 'Cache-Control: max-age=0' \
	-H 'Upgrade-Insecure-Requests: 1' \
	-H 'Origin: http://mail.keymobile.com.cn' \
	-H 'Content-Type: multipart/form-data; boundary=----WebKitFormBoundaryeB3SGCWBtfb04Z0e' \
	-H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.69 Safari/537.36' \
	-H 'Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9' \
	-H 'Referer: http://mail.keymobile.com.cn/Session/2262-JofRNHEj229osjQt8Bpq/Mailboxes.wssp' \
	-H 'Accept-Language: en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7' \
	--data-raw $'------WebKitFormBoundaryeB3SGCWBtfb04Z0e\r\nContent-Disposition: form-data; name="Username"\r\n\r\nfoobar\r\n------WebKitFormBoundaryeB3SGCWBtfb04Z0e\r\nContent-Disposition: form-data; name="Password"\r\n\r\njames1212\r\n------WebKitFormBoundaryeB3SGCWBtfb04Z0e\r\nContent-Disposition: form-data; name="login"\r\n\r\nResume\r\n------WebKitFormBoundaryeB3SGCWBtfb04Z0e\r\nContent-Disposition: form-data; name="restoreSessionPage"\r\n\r\nMailboxes.wssp\r\n------WebKitFormBoundaryeB3SGCWBtfb04Z0e\r\nContent-Disposition: form-data; name="restoreCharset"\r\n\r\nutf-8\r\n------WebKitFormBoundaryeB3SGCWBtfb04Z0e--\r\n' \
	--compressed \
	--insecure`
	bash := "D:/app/msys64/usr/bin/bash.exe"
	curl(bash, command)
}
