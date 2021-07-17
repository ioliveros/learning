package main

import (
	"net/http"
	"fmt"	
	"io/ioutil"
)

func main() {

	resp, err := http.Get("http://example.com/")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	defer resp.Body().Close()

	body, err := io.ReadAll(resp.Body)
	

}
