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
	if err := http.ListenAndServe(addr, nil); err != nil {
		fmt.Fprintf(os.Stderr, "Fatal: %s\n", err)
	}
}

func catHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Request from %s\n", r.RemoteAddr)

	w.WriteHeader(200)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s %s %s\n", r.Method, r.RequestURI, r.Proto)
	r.Header.WriteSubset(os.Stdout, nil)
	fmt.Printf("\n")
	fmt.Println(string(body))
}
