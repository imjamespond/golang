package main

import (
	cc_utils "codechiev/utils"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"net/http"
	"os"
	"strconv"
	"test-go1/utils"
	"time"
)

var (
	num int32
)

func init() {
	rand.Seed(time.Now().UnixNano())
	num = rand.Int31n(1000)
}

func main() {
	for idx, args := range os.Args {
		fmt.Println("arg"+strconv.Itoa(idx)+":", args)
	}

	h1 := func(w http.ResponseWriter, _ *http.Request) {
		conn, err := net.Dial("tcp", "bing.com:80")
		cc_utils.FatalIf(err)
		ip := utils.Ip(conn)
		io.WriteString(w, fmt.Sprintf("%d Hello from #%s\n", num, ip.String()))
	}

	http.HandleFunc("/", h1)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
