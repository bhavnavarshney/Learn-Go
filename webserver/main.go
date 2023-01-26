package main

import (
	"fmt"
	"html"
	"net/http"
	"strconv"
	"sync"
)

var mutex = &sync.Mutex{}
var counter int

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q!", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hi")
	})

	http.HandleFunc("/increment", incrementCounter)

	http.ListenAndServe(":8081", nil)
}

func incrementCounter(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	counter++
	fmt.Fprint(w, "The count is: ", strconv.Itoa(counter))
	mutex.Unlock()
}
