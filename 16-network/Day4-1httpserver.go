package main

import (
	"fmt"
	"net/http"
)

func helloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	fmt.Fprintf(w, "Hello %s", r.URL.Path[1:]) //print the url path after /
}
func main() {
	http.HandleFunc("/", helloServer)
	err := http.ListenAndServe("localhost:8090", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Server listening on 8090")
}
