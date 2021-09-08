package main

import (
	cc_utils "codechiev/utils"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
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
		before := time.Now()
		resp, err := http.Get("http://testgo-svc.default.svc.cluster.local:8080/ip")
		cc_utils.FatalIf(err)
		body, err := io.ReadAll(resp.Body)
		cc_utils.FatalIf(err)
		io.WriteString(w, fmt.Sprintf("%d, cost %d millis, Hello from %s\n", num, time.Since(before).Milliseconds(), string(body)))
	}

	http.HandleFunc("/", h1)

	log.Fatal(http.ListenAndServe(":8082", nil))
}
