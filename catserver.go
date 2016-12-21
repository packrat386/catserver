package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = ":8080"
	}

	http.HandleFunc("/", catHandler)
	http.ListenAndServe(addr, nil)
}

func catHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got Req")
	w.WriteHeader(200)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
