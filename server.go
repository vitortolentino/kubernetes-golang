package main

import "net/http"

func main() {
	http.HandleFunc("/", HelloHandler)
	http.ListenAndServe(":8000", nil)
}

func HelloHandler(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("<h1>Hello World!!!</h1>"))
}