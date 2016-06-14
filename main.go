package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// fast test
func fastHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this request is fast")
}

// read test
func readHandler(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("/etc/hosts")
	if err != nil {
		fmt.Fprintf(w, "this is read file error: "+err.Error())
	} else {
		fmt.Fprintf(w, "this is read file data: "+string(data))
	}
}

// slow test
func slowHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(5 * time.Second)
	fmt.Fprintf(w, "this request is slow")
}

// panic test
func panicHandler(w http.ResponseWriter, r *http.Request) {
	panic("something is too wrong")
	fmt.Fprintf(w, "this request is panic")
}

func main() {
	http.HandleFunc("/fast", fastHandler)
	http.HandleFunc("/read", readHandler)
	http.HandleFunc("/slow", slowHandler)
	http.HandleFunc("/panic", panicHandler)

	http.ListenAndServe(":3000", nil)
}
