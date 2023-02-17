package main

import "net/http"

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte("Hello, World!"))
	})

	mux := http.NewServeMux()

	mux.HandleFunc("/test", func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte("Hello, World!"))
	})

	http.ListenAndServe(":80", mux)
}
