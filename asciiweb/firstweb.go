package main

import (
	"fmt"
	"net/http"
)

func webHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello welcome to my site"))
}
func siteHandler(w http.ResponseWriter, r *http.Request)  {
	
}
func main() {
	http.HandleFunc("/hello", webHandler)

	fmt.Println("server is still running")

	http.ListenAndServe(":8080", nil)

}
