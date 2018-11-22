package main

import (
	"fmt"
	"io"
	"net/http"

	"goproject/pkg/hello"

	"github.com/gorilla/mux"
)

// Respond to call on endpoint
func respond(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello from your Go sample application running in Microclimate!")
	fmt.Printf("main.go: Go app endpoint response\n")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", respond)
	http.HandleFunc("/", respond)
	http.HandleFunc("/hello", hello.SayHello)
	fmt.Printf("main.go: Go app listening on port 8000\n")
	http.ListenAndServe(":8000", nil)
}
