package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)	
	fmt.Println("Server 8080")
}
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("salom \n %s", r.URL.Path[1:])
}
