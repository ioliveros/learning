package main

import (
	"fmt"
	"net/http"
)

type Hello struct{}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("API %v\n", r.URL.Opaque)
	fmt.Printf("Method %v\n", r.Method)
	fmt.Printf("Body %v\n", r.Body)
	fmt.Printf("Form %v\n", r.Form)

	fmt.Fprint(w, "Hello!")
}

func main() {
	var h Hello
	fmt.Println("Server starting ..")
	http.ListenAndServe("localhost:4000", h)
}
