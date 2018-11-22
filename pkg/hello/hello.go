package hello

import (
	"io"
	"net/http"
)

// Function names starting with an upper case letter are exported out of the package
func SayHello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if isValidName(name) {
		io.WriteString(w, "Hello, "+name+"!")
	} else {
		io.WriteString(w, "Hello, World!")
	}
}

// Function names starting with a lower case letter are not exported, and only have a scope to the package it's in
func isValidName(name string) bool {
	return name != ""
}
