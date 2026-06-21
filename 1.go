 package main

// import (
// 	"fmt"
// 	"net/http"
// )

// // func webHandler(w http.ResponseWriter, r *http.Request) {

// // 	fmt.Fprintln(w, "Hello World")

// // }
// func pathHandler(w http.ResponseWriter, r *http.Request) {

// 	name := r.URL.Query().Get("name")

// 	if name == "" {
// 		name = "Guest"
// 	}
// 	if r.Method != http.MethodGet {
// 		http.Error(w, "method not found", http.StatusMethodNotAllowed)
// 		return
// 	}
// }

// func main() {
// 	//http.HandleFunc("/hello", webHandler)
// 	http.HandleFunc("/hello", pathHandler)

// 	fmt.Println("server is running")

// 	http.ListenAndServe(":8090", nil)

// }
