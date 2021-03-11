package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://httpbin.org/get")
	if err != nil {
		log.Fatalln("Unable to read response")
	}
	defer res.Body.Close()
	content , err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln("Unable to read the body")
	}

	fmt.Print(string(content))
}
