package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello World!")

	var m map[string]int
	m = make(map[string]int)
	m["key"] = 42

	fmt.Printf("Map --> %v", m)

	w.Header().Add("Content-Type", "application/json")
	// w.Write(http.StatusOK, )

	fmt.Fprint(w, m)
}

func main() {
	fmt.Println("Running HTTP Server..")

	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err)
	}

}
