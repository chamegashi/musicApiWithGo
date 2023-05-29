package main

import (
	"fmt"
	"log"
	"net/http"
)

func testHandler(w http.ResponseWriter, r *http.Request) {
	hello := []byte("hello")
	_, err := w.Write(hello)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/test", testHandler)
	fmt.Println("server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
