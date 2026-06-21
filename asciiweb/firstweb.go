//  package main

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
package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Task 2: Reject any HTTP method that is not GET
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Task 1: Extract query parameters using r.URL.Query().Get("name")
	name := r.URL.Query().Get("name")

	// Goal: If the parameter is missing, default to "Hello, Guest!"
	if name == "" {
		name = "Guest"
	}

	// Goal: Respond back to the client
	fmt.Fprintf(w, "Hello, %s!", name)
}

func main() {
	// Register the handler for the /hello endpoint
	http.HandleFunc("/hello", helloHandler)

	// Start the server
	fmt.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
