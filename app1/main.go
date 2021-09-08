package main

import (
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
		io.WriteString(w, fmt.Sprintf("Hello from #%d\n", num))
	}

	http.HandleFunc("/", h1)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
