package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", catHandler)
	http.ListenAndServe(":8080", nil)
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
